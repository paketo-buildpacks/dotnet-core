package integration_test

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/paketo-buildpacks/packit/v2/pexec"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"

	. "github.com/onsi/gomega"
)

var (
	dotnetCoreBuildpack string
	builder             struct {
		Local struct {
			Stack struct {
				ID string `json:"id"`
			} `json:"stack"`
		} `json:"local_info"`
	}
)

func TestIntegration(t *testing.T) {
	Expect := NewWithT(t).Expect

	output, err := exec.Command("bash", "-c", "../scripts/package.sh --version 1.2.3").CombinedOutput()
	Expect(err).NotTo(HaveOccurred(), string(output))

	dotnetCoreBuildpack, err = filepath.Abs("../build/buildpackage.cnb")
	Expect(err).NotTo(HaveOccurred())

	SetDefaultEventuallyTimeout(10 * time.Second)

	buf := bytes.NewBuffer(nil)
	cmd := pexec.NewExecutable("pack")
	Expect(cmd.Execute(pexec.Execution{
		Args:   []string{"builder", "inspect", "--output", "json"},
		Stdout: buf,
		Stderr: buf,
	})).To(Succeed(), buf.String())

	Expect(json.Unmarshal(buf.Bytes(), &builder)).To(Succeed(), buf.String())

	suite := spec.New("Integration", spec.Parallel(), spec.Report(report.Terminal{}))
	suite("FDD", testFDD)
	suite("FDE", testFDE)
	suite("SelfContained", testSelfContained)
	suite("Source", testSource)
	suite("ReproducibleBuilds", testReproducibleBuilds)
	suite.Run(t)
}
