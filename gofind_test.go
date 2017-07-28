package main_test

import (
	"os"
	"os/exec"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("GoFind", func() {
	var (
		cmdErr     error
		cmd        *exec.Cmd
		session    *gexec.Session
		currentDir string
		params     []string
	)

	BeforeEach(func() {
		var err error
		currentDir, err = filepath.Abs(filepath.Dir(os.Args[0]))
		Expect(err).NotTo(HaveOccurred())
		os.Setenv("GOPATH", filepath.Join(currentDir, "assets"))
	})

	It("returns the correct repo path", func() {
		params = []string{"foobar"}
		cmd = exec.Command(binaryPath, params...)

		session, cmdErr = gexec.Start(cmd, os.Stdout, os.Stderr)
		Expect(cmdErr).NotTo(HaveOccurred())
		Eventually(session).Should(gexec.Exit(0))

		expectedRepoPath := filepath.Join(currentDir, "assets", "src", "github.com", "foouser", "foobar")
		Eventually(session.Out).Should(gbytes.Say(expectedRepoPath))
	})

	Context("When the query contains a substring", func() {
		It("returns the correct repo path", func() {
			params = []string{"foob"}
			cmd = exec.Command(binaryPath, params...)

			session, cmdErr = gexec.Start(cmd, os.Stdout, os.Stderr)
			Expect(cmdErr).NotTo(HaveOccurred())
			Eventually(session).Should(gexec.Exit(0))

			expectedRepoPath := filepath.Join(currentDir, "assets", "src", "github.com", "foouser", "foobar")
			Eventually(session.Out).Should(gbytes.Say(expectedRepoPath))
		})

		It("returns the correct repo path", func() {
			params = []string{"fob"}
			cmd = exec.Command(binaryPath, params...)

			session, cmdErr = gexec.Start(cmd, os.Stdout, os.Stderr)
			Expect(cmdErr).NotTo(HaveOccurred())
			Eventually(session).Should(gexec.Exit(0))

			expectedRepoPath := filepath.Join(currentDir, "assets", "src", "github.com", "foouser", "foobar")
			Eventually(session.Out).Should(gbytes.Say(expectedRepoPath))
		})

		It("returns the correct repo path", func() {
			params = []string{"fbr"}
			cmd = exec.Command(binaryPath, params...)

			session, cmdErr = gexec.Start(cmd, os.Stdout, os.Stderr)
			Expect(cmdErr).NotTo(HaveOccurred())
			Eventually(session).Should(gexec.Exit(0))

			expectedRepoPath := filepath.Join(currentDir, "assets", "src", "github.com", "foouser", "foobar")
			Eventually(session.Out).Should(gbytes.Say(expectedRepoPath))
		})
	})
})
