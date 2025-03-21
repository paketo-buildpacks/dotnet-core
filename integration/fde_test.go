package integration_test

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/paketo-buildpacks/occam"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
	. "github.com/paketo-buildpacks/occam/matchers"
)

func testFDE(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect     = NewWithT(t).Expect
		Eventually = NewWithT(t).Eventually

		pack   occam.Pack
		docker occam.Docker
	)

	it.Before(func() {
		pack = occam.NewPack()
		docker = occam.NewDocker()
	})

	context("when building a .Net core app that is an FDE", func() {
		var (
			image     occam.Image
			container occam.Container

			name    string
			source  string
			sbomDir string
		)

		it.Before(func() {
			var err error
			name, err = occam.RandomName()
			Expect(err).NotTo(HaveOccurred())

			source, err = occam.Source(filepath.Join("testdata", "fde-app"))
			Expect(err).NotTo(HaveOccurred())

			sbomDir, err = os.MkdirTemp("", "sbom")
			Expect(err).NotTo(HaveOccurred())
			Expect(os.Chmod(sbomDir, os.ModePerm)).To(Succeed())
		})

		it.After(func() {
			Expect(docker.Container.Remove.Execute(container.ID)).To(Succeed())
			Expect(docker.Image.Remove.Execute(image.ID)).To(Succeed())
			Expect(docker.Volume.Remove.Execute(occam.CacheVolumeNames(name))).To(Succeed())

			Expect(os.RemoveAll(source)).To(Succeed())
			Expect(os.RemoveAll(sbomDir)).To(Succeed())
		})

		it("creates a working OCI image", func() {
			var err error
			var logs fmt.Stringer
			image, logs, err = pack.WithNoColor().Build.
				WithBuildpacks(dotnetCoreBuildpack).
				WithSBOMOutputDir(sbomDir).
				WithPullPolicy("never").
				Execute(name, source)
			Expect(err).NotTo(HaveOccurred(), logs.String())

			container, err = docker.Container.Run.
				WithEnv(map[string]string{"PORT": "8080"}).
				WithPublish("8080").
				WithPublishAll().
				Execute(image.ID)
			Expect(err).NotTo(HaveOccurred())

			Expect(logs).To(ContainLines(ContainSubstring("Buildpack for ASP.NET Core Runtime")))
			Expect(logs).To(ContainLines(ContainSubstring("Buildpack for ICU")))
			Expect(logs).To(ContainLines(ContainSubstring("Buildpack for .NET Execute")))

			Expect(logs).NotTo(ContainLines(ContainSubstring("Buildpack for Environment Variables")))
			Expect(logs).NotTo(ContainLines(ContainSubstring("Buildpack for Image Labels")))

			Eventually(container).Should(Serve(ContainSubstring("<title>source_app</title>")).OnPort(8080))

			// check that all required SBOM files are present
			Expect(filepath.Join(sbomDir, "sbom", "launch", "paketo-buildpacks_dotnet-core-aspnet-runtime", "dotnet-core-aspnet-runtime", "sbom.cdx.json")).To(BeARegularFile())
			Expect(filepath.Join(sbomDir, "sbom", "launch", "paketo-buildpacks_dotnet-core-aspnet-runtime", "dotnet-core-aspnet-runtime", "sbom.spdx.json")).To(BeARegularFile())
			Expect(filepath.Join(sbomDir, "sbom", "launch", "paketo-buildpacks_dotnet-core-aspnet-runtime", "dotnet-core-aspnet-runtime", "sbom.syft.json")).To(BeARegularFile())

			Expect(filepath.Join(sbomDir, "sbom", "launch", "paketo-buildpacks_icu", "icu", "sbom.cdx.json")).To(BeARegularFile())
			Expect(filepath.Join(sbomDir, "sbom", "launch", "paketo-buildpacks_icu", "icu", "sbom.spdx.json")).To(BeARegularFile())
			Expect(filepath.Join(sbomDir, "sbom", "launch", "paketo-buildpacks_icu", "icu", "sbom.syft.json")).To(BeARegularFile())

			Expect(filepath.Join(sbomDir, "sbom", "launch", "paketo-buildpacks_dotnet-execute", "sbom.cdx.json")).To(BeARegularFile())
			Expect(filepath.Join(sbomDir, "sbom", "launch", "paketo-buildpacks_dotnet-execute", "sbom.spdx.json")).To(BeARegularFile())
			Expect(filepath.Join(sbomDir, "sbom", "launch", "paketo-buildpacks_dotnet-execute", "sbom.syft.json")).To(BeARegularFile())
		})

		context("when using ca certs buildpack", func() {
			var (
				client *http.Client
			)
			it.Before(func() {
				var err error

				// Remove source directory created in the it.Before
				Expect(os.RemoveAll(source)).To(Succeed())

				source, err = occam.Source(filepath.Join("testdata", "ca-cert-apps"))
				Expect(err).NotTo(HaveOccurred())

				caCert, err := os.ReadFile(filepath.Join(source, "client-certs", "ca.pem"))
				Expect(err).ToNot(HaveOccurred())

				caCertPool := x509.NewCertPool()
				caCertPool.AppendCertsFromPEM(caCert)

				cert, err := tls.LoadX509KeyPair(filepath.Join(source, "client-certs", "cert.pem"), filepath.Join(source, "client-certs", "key.pem"))
				Expect(err).ToNot(HaveOccurred())

				client = &http.Client{
					Transport: &http.Transport{
						TLSClientConfig: &tls.Config{
							RootCAs:      caCertPool,
							Certificates: []tls.Certificate{cert},
							MinVersion:   tls.VersionTLS12,
						},
					},
				}
			})

			it("builds a working OCI image and uses a client-side CA cert for requests", func() {
				var err error
				var logs fmt.Stringer
				image, logs, err = pack.WithNoColor().Build.
					WithBuildpacks(dotnetCoreBuildpack).
					WithPullPolicy("never").
					Execute(name, filepath.Join(source, "fde-app"))
				Expect(err).NotTo(HaveOccurred(), logs.String())

				Expect(logs).To(ContainLines(ContainSubstring("Buildpack for CA Certificates")))
				Expect(logs).To(ContainLines(ContainSubstring("Buildpack for ICU")))
				Expect(logs).To(ContainLines(ContainSubstring("Buildpack for ASP.NET Core Runtime")))
				Expect(logs).To(ContainLines(ContainSubstring("Buildpack for .NET Execute")))

				container, err = docker.Container.Run.
					WithPublish("8080").
					WithEnv(map[string]string{
						"SERVICE_BINDING_ROOT": "/bindings",
					}).
					WithVolumes(fmt.Sprintf("%s:/bindings/ca-certificates", filepath.Join(source, "binding"))).
					Execute(image.ID)
				Expect(err).NotTo(HaveOccurred())

				Eventually(func() string {
					cLogs, err := docker.Container.Logs.Execute(container.ID)
					Expect(err).NotTo(HaveOccurred())
					return cLogs.String()
				}).Should(
					ContainSubstring("Added 1 additional CA certificate(s) to system truststore"),
				)

				request, err := http.NewRequest("GET", fmt.Sprintf("https://localhost:%s", container.HostPort("8080")), nil)
				Expect(err).NotTo(HaveOccurred())

				var response *http.Response
				Eventually(func() error {
					var err error
					response, err = client.Do(request)
					return err
				}).Should(BeNil())
				defer response.Body.Close()

				Expect(response.StatusCode).To(Equal(http.StatusOK))

				content, err := io.ReadAll(response.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(content)).To(ContainSubstring("Hello World!"))
			})
		})

		context("when using optional utility buildpacks", func() {
			var procfileContainer occam.Container
			it.Before(func() {
				Expect(os.WriteFile(filepath.Join(source, "Procfile"), []byte("procfile: echo Procfile command"), 0644)).To(Succeed())
			})

			it.After(func() {
				Expect(docker.Container.Remove.Execute(procfileContainer.ID)).To(Succeed())
			})

			it("builds a working OCI image and run the app with the start command from the Procfile and other utility buildpacks", func() {
				var err error
				var logs fmt.Stringer
				image, logs, err = pack.WithNoColor().Build.
					WithBuildpacks(dotnetCoreBuildpack).
					WithPullPolicy("never").
					WithEnv(map[string]string{
						"BPE_SOME_VARIABLE":      "some-value",
						"BP_IMAGE_LABELS":        "some-label=some-value",
						"BP_LIVE_RELOAD_ENABLED": "true",
					}).
					Execute(name, source)
				Expect(err).NotTo(HaveOccurred(), logs.String())

				Expect(logs).To(ContainLines(ContainSubstring("Buildpack for ICU")))
				Expect(logs).To(ContainLines(ContainSubstring("Buildpack for ASP.NET Core Runtime")))
				Expect(logs).To(ContainLines(ContainSubstring("Buildpack for .NET Execute")))
				Expect(logs).To(ContainLines(ContainSubstring("Buildpack for Procfile")))
				Expect(logs).To(ContainLines(ContainSubstring("Buildpack for Environment Variables")))
				Expect(logs).To(ContainLines(ContainSubstring("Buildpack for Image Labels")))
				Expect(logs).To(ContainLines(ContainSubstring("Buildpack for Watchexec")))

				envVar, err := image.BuildpackForKey("paketo-buildpacks/environment-variables")
				Expect(err).NotTo(HaveOccurred())
				Expect(envVar.Layers["environment-variables"].Metadata["variables"]).To(Equal(map[string]interface{}{"SOME_VARIABLE": "some-value"}))
				Expect(image.Labels["some-label"]).To(Equal("some-value"))

				container, err = docker.Container.Run.
					WithEnv(map[string]string{"PORT": "8080"}).
					WithPublish("8080").
					WithPublishAll().
					Execute(image.ID)
				Expect(err).NotTo(HaveOccurred())

				Eventually(container).Should(Serve(ContainSubstring("<title>source_app</title>")).OnPort(8080))

				procfileContainer, err = docker.Container.Run.
					WithEntrypoint("procfile").
					Execute(image.ID)
				Expect(err).NotTo(HaveOccurred())

				Eventually(func() string {
					containerLogs, err := docker.Container.Logs.Execute(procfileContainer.ID)
					Expect(err).NotTo(HaveOccurred())
					return containerLogs.String()
				}).Should(ContainSubstring("Procfile command"))
			})
		})
	})
}
