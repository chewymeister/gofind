package main_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("GoFind", func() {
	var (
		cmd     *exec.Cmd
		session *gexec.Session
	)

	BeforeEach(func() {
		params := []string{"cf-redis-broker"}
		cmd = exec.Command(binaryPath, params...)
	})

	It("should run the binary", func() {
		var err error
		session, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		// Expect(session.ExitCode()).To(Equal(0))
		Eventually(session).Should(gbytes.Say("cf-redis-broker"))
	})
})
