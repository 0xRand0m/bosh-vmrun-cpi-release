package vm_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

func TestVm(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "VM Suite")
}

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
