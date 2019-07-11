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

func TestUrlsHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Urls Handler Suite")
}
var _ = Describe("Urls Handler", func() {

	BeforeSuite(func() {
		os.Chdir(build.Default.GOPATH + "/src/github.com/ildomm/linx_challenge/")
		config.Setup()
		db.Setup()
	})

	BeforeEach(func() {
		db.CleanDatabases()
	})

	Context("Redirect", func() {
		It("success, it has to redirect", func() {
			Skip("validate reflection type return")
		})
	})

	Context("Delete", func() {
		It("success, it has to erase all data", func() {
			Skip("validate reflection type return")
		})
	})


})
