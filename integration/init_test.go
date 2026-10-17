package integration_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"testing"
	"time"

	"github.com/paketo-buildpacks/occam"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"

	. "github.com/onsi/gomega"
)

var (
	dotnetCoreBuildpack string
	ubiNodejsExtension  string
	config              struct {
		UbiNodejsExtension string `json:"ubi-nodejs-extension"`
	}
)

func TestIntegration(t *testing.T) {
	Expect := NewWithT(t).Expect

	output, err := exec.Command("bash", "-c", "../scripts/package.sh --version 1.2.3").CombinedOutput()
	Expect(err).NotTo(HaveOccurred(), string(output))

	dotnetCoreBuildpack, err = filepath.Abs("../build/buildpackage.cnb")
	Expect(err).NotTo(HaveOccurred())

	file, err := os.Open("../integration.json")
	Expect(err).NotTo(HaveOccurred())

	Expect(json.NewDecoder(file).Decode(&config)).To(Succeed())
	Expect(file.Close()).To(Succeed())

	pack := occam.NewPack()

	builder, err := pack.Builder.Inspect.Execute()
	Expect(err).NotTo(HaveOccurred())

	isUbiBuilder := regexp.MustCompile(`ubi8`).MatchString(builder.BuilderName)

	if isUbiBuilder {
		Expect(occam.NewDocker().Pull.Execute(config.UbiNodejsExtension)).To(Succeed())
		ubiNodejsExtension = config.UbiNodejsExtension
	}

	SetDefaultEventuallyTimeout(10 * time.Second)

	suite := spec.New("Integration", spec.Parallel(), spec.Report(report.Terminal{}))
	suite("FDD", testFDD)
	suite("FDE", testFDE)
	suite("SelfContained", testSelfContained)
	suite("Source", testSource)
	suite("ReproducibleBuilds", testReproducibleBuilds)
	suite.Run(t)
}
