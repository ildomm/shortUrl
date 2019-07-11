package handlers_test

import (
	"github.com/ildomm/linx_challenge/config"
	"github.com/ildomm/linx_challenge/db"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go/build"
	"os"
	"testing"
)

func TestStatsHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stats Handler Suite")
}
var _ = Describe("Stats Handler", func() {

	BeforeSuite(func() {
		os.Chdir(build.Default.GOPATH + "/src/github.com/ildomm/linx_challenge/")
		config.Setup()
		db.Setup()
	})

	BeforeEach(func() {
		db.CleanDatabases()
	})

	Context("Url", func() {
		It("success, it has to provide url stats", func() {
			Skip("validate reflection type return")
		})
	})

	Context("user", func() {
		It("success, it has to provide user stats", func() {
			Skip("validate reflection type return")
		})
	})

	Context("system", func() {
		It("success, it has to provide system stats", func() {
			Skip("validate reflection type return")
		})
	})

})
