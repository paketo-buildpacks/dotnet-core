package integration_test

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/paketo-buildpacks/occam"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
	. "github.com/paketo-buildpacks/occam/matchers"
)

func testFDD(t *testing.T, context spec.G, it spec.S) {
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

	context("when building a .Net core app that is an FDD", func() {
		var (
			image     occam.Image
			container occam.Container

			name   string
			source string
		)

		it.Before(func() {
			var err error
			name, err = occam.RandomName()
			Expect(err).NotTo(HaveOccurred())

			source, err = occam.Source(filepath.Join("testdata", "fdd-app"))
			Expect(err).NotTo(HaveOccurred())
		})

		it.After(func() {
			Expect(docker.Container.Remove.Execute(container.ID)).To(Succeed())
			Expect(docker.Image.Remove.Execute(image.ID)).To(Succeed())
			Expect(docker.Volume.Remove.Execute(occam.CacheVolumeNames(name))).To(Succeed())
			Expect(os.RemoveAll(source)).To(Succeed())
		})

		it("creates a working OCI image", func() {
			var err error
			var logs fmt.Stringer
			image, logs, err = pack.WithNoColor().Build.
				WithBuildpacks(dotnetCoreBuildpack).
				WithPullPolicy("never").
				Execute(name, source)
			Expect(err).NotTo(HaveOccurred(), logs.String())

			container, err = docker.Container.Run.
				WithEnv(map[string]string{"PORT": "8080"}).
				WithPublish("8080").
				WithPublishAll().
				Execute(image.ID)
			Expect(err).NotTo(HaveOccurred())

			Eventually(container).Should(BeAvailable())

			Expect(logs).To(ContainLines(ContainSubstring(".NET Core Runtime Buildpack")))
			Expect(logs).To(ContainLines(ContainSubstring("ASP.NET Core Buildpack")))
			Expect(logs).To(ContainLines(ContainSubstring(".NET Core SDK Buildpack")))
			Expect(logs).To(ContainLines(ContainSubstring("ICU Buildpack")))
			Expect(logs).To(ContainLines(ContainSubstring(".NET Execute Buildpack")))

			Expect(logs).NotTo(ContainLines(ContainSubstring("Environment Variables Buildpack")))
			Expect(logs).NotTo(ContainLines(ContainSubstring("Image Labels Buildpack")))

			response, err := http.Get(fmt.Sprintf("http://localhost:%s", container.HostPort("8080")))
			Expect(err).NotTo(HaveOccurred())
			defer response.Body.Close()

			Expect(response.StatusCode).To(Equal(http.StatusOK))

			content, err := ioutil.ReadAll(response.Body)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(content)).To(ContainSubstring("<title>react_app</title>"))
		})

		context("when using ca certs buildpack", func() {
			var (
				client *http.Client
			)
			it.Before(func() {
				var err error
				source, err = occam.Source(filepath.Join("testdata", "ca-cert-apps"))
				Expect(err).NotTo(HaveOccurred())

				caCert, err := ioutil.ReadFile(filepath.Join(source, "client-certs", "ca.pem"))
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
					Execute(name, filepath.Join(source, "fdd-app"))
				Expect(err).NotTo(HaveOccurred(), logs.String())

				Expect(logs).To(ContainLines(ContainSubstring("CA Certificates Buildpack")))
				Expect(logs).To(ContainLines(ContainSubstring(".NET Core Runtime Buildpack")))
				Expect(logs).To(ContainLines(ContainSubstring("ASP.NET Core Buildpack")))
				Expect(logs).To(ContainLines(ContainSubstring(".NET Core SDK Buildpack")))
				Expect(logs).To(ContainLines(ContainSubstring("ICU Buildpack")))
				Expect(logs).To(ContainLines(ContainSubstring(".NET Execute Buildpack")))

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

				content, err := ioutil.ReadAll(response.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(content)).To(ContainSubstring("Hello World!"))
			})
		})

		context("when using optional utility buildpacks", func() {
			it.Before(func() {
				Expect(ioutil.WriteFile(filepath.Join(source, "Procfile"), []byte("web: echo Procfile command && dotnet /workspace/react-app.dll --urls http://0.0.0.0:${PORT:-8080}"), 0644)).To(Succeed())
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

				container, err = docker.Container.Run.
					WithEnv(map[string]string{"PORT": "8080"}).
					WithPublish("8080").
					WithPublishAll().
					Execute(image.ID)
				Expect(err).NotTo(HaveOccurred())

				Eventually(container).Should(BeAvailable())

				Expect(logs).To(ContainLines(ContainSubstring(".NET Core Runtime Buildpack")))
				Expect(logs).To(ContainLines(ContainSubstring("ASP.NET Core Buildpack")))
				Expect(logs).To(ContainLines(ContainSubstring(".NET Core SDK Buildpack")))
				Expect(logs).To(ContainLines(ContainSubstring("ICU Buildpack")))
				Expect(logs).To(ContainLines(ContainSubstring(".NET Execute Buildpack")))
				Expect(logs).To(ContainLines(ContainSubstring("Procfile Buildpack")))
				Expect(logs).To(ContainLines(ContainSubstring("Environment Variables Buildpack")))
				Expect(logs).To(ContainLines(ContainSubstring("Image Labels Buildpack")))
				Expect(logs).To(ContainLines(ContainSubstring("Watchexec Buildpack")))

				Expect(image.Buildpacks[8].Key).To(Equal("paketo-buildpacks/environment-variables"))
				Expect(image.Buildpacks[8].Layers["environment-variables"].Metadata["variables"]).To(Equal(map[string]interface{}{"SOME_VARIABLE": "some-value"}))
				Expect(image.Labels["some-label"]).To(Equal("some-value"))

				containerLogs, err := docker.Container.Logs.Execute(container.ID)
				Expect(err).NotTo(HaveOccurred())

				Expect(containerLogs.String()).To(ContainSubstring("Procfile command"))

				response, err := http.Get(fmt.Sprintf("http://localhost:%s", container.HostPort("8080")))
				Expect(err).NotTo(HaveOccurred())
				defer response.Body.Close()

				Expect(response.StatusCode).To(Equal(http.StatusOK))

				content, err := ioutil.ReadAll(response.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(content)).To(ContainSubstring("<title>react_app</title>"))
			})
		})
	})
}
