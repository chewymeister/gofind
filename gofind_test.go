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
		gofindCmd *exec.Cmd
		params    []string
	)

	BeforeEach(func() {
		testGoPath := filepath.Join(currentDir(), "assets")
		os.Setenv("GOPATH", testGoPath)
	})

	It("returns the correct repo path", func() {
		params = []string{"foobar"}
		gofindCmd = exec.Command(binaryPath, params...)

		session := executeCommand(gofindCmd)

		expectedRepoPath := repoPath("foobar")
		Eventually(session.Out).Should(gbytes.Say(expectedRepoPath))
	})

	Context("When the query contains a substring", func() {
		It("returns the correct repo path", func() {
			params = []string{"fooba"}
			gofindCmd = exec.Command(binaryPath, params...)

			session := executeCommand(gofindCmd)

			expectedRepoPath := repoPath("foobar")
			Eventually(session.Out).Should(gbytes.Say(expectedRepoPath))
		})

		It("returns the correct repo path", func() {
			params = []string{"foob"}
			gofindCmd = exec.Command(binaryPath, params...)

			session := executeCommand(gofindCmd)

			expectedRepoPath := repoPath("foobar")
			Eventually(session.Out).Should(gbytes.Say(expectedRepoPath))
		})

		It("returns the correct repo path", func() {
			params = []string{"fbr"}
			gofindCmd = exec.Command(binaryPath, params...)

			session := executeCommand(gofindCmd)

			expectedRepoPath := repoPath("foobar")
			Eventually(session.Out).Should(gbytes.Say(expectedRepoPath))
		})
	})
})

func repoPath(repoName string) string {
	return filepath.Join(currentDir(), "assets", "src", "simple", repoName)
}

func currentDir() string {
	currentdirPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	Expect(err).NotTo(HaveOccurred())
	return currentdirPath
}

func executeCommand(cmd *exec.Cmd) *gexec.Session {
	session, cmdErr := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(cmdErr).NotTo(HaveOccurred())
	Eventually(session, "5s").Should(gexec.Exit(0))

	return session
}
