package main_test

import (
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gexec"
)

func TestGoFind(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoFind Test Suite")
}

var binaryPath string

var _ = SynchronizedBeforeSuite(func() []byte {
	srcPath := filepath.Join("github.com", "chewymeister", "gofind")
	binaryPath, err := gexec.Build(srcPath)
	Expect(err).NotTo(HaveOccurred())

	return []byte(binaryPath)
}, func(rawBinaryPath []byte) {
	binaryPath = string(rawBinaryPath)
})

var _ = SynchronizedAfterSuite(func() {}, func() {
	gexec.CleanupBuildArtifacts()
})
