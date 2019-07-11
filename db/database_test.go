package db_test

import (
	"github.com/ildomm/linx_challenge/config"
	"github.com/ildomm/linx_challenge/db"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go/build"
	"os"
	"testing"
)

func TestDatabase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Database Suite")
}

var _ = Describe("Connection", func() {
	BeforeSuite(func() {
		os.Chdir(build.Default.GOPATH + "/src/github.com/ildomm/linx_challenge/")
		config.Setup()
		db.Setup()
	})

	BeforeEach(func() {
	})

	Context("Evolve", func() {
		It("able to reset", func() {
			defer GinkgoRecover()
			db.CleanDatabases()
			Expect(db.CountTable("users")).To(Equal(int(0)))
			Expect(db.CountTable("urls")).To(Equal(int(0)))
		})
	})

	Context("Connect", func() {

		It("has DATABASE instance", func() {
			defer GinkgoRecover()
			Expect(db.Setup()).ToNot(BeNil())
		})

	})
})
