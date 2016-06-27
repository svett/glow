package glow_test

import (
	"os"

	"github.com/svett/glow"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var _ = Describe("DB", func() {
	It("create a gorm.DB successfully", func() {
		os.Setenv("GORM_DIALECT", "sqlite3")
		os.Setenv("GORM_CONNECTION", "glow.db")

		db, err := glow.OpenDB()
		Expect(db).NotTo(BeNil())
		Expect(err).NotTo(HaveOccurred())
	})

	Context("when GORM_DIALECT variable is not set", func() {
		BeforeEach(func() {
			os.Unsetenv("GORM_DIALECT")
			os.Setenv("GORM_CONNECTION", "glow.db")
		})

		It("returns the error", func() {
			db, err := glow.OpenDB()
			Expect(db).To(BeNil())
			Expect(err).To(MatchError("GORM_DIALECT variable is not set"))
		})
	})

	Context("when GORM_CONNECTION variable is not set", func() {
		BeforeEach(func() {
			os.Setenv("GORM_DIALECT", "sqlite3")
			os.Unsetenv("GORM_CONNECTION")
		})

		It("returns the error", func() {
			db, err := glow.OpenDB()
			Expect(db).To(BeNil())
			Expect(err).To(MatchError("GORM_CONNECTION variable is not set"))
		})
	})
})
