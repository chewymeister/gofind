package main_test

import (
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("GoFind", func() {
	var (
		cmdErr  error
		cmd     *exec.Cmd
		session *gexec.Session
	)

	BeforeEach(func() {
		params := []string{"cf-redis-broker"}
		cmd = exec.Command(binaryPath, params...)
	})

	It("should run the binary", func() {
		session, cmdErr = gexec.Start(cmd, os.Stdout, os.Stderr)
		Expect(cmdErr).NotTo(HaveOccurred())
		Eventually(session).Should(gexec.Exit(0))

		Eventually(session.Out).Should(gbytes.Say("cf-redis-broker"))
	})
})
