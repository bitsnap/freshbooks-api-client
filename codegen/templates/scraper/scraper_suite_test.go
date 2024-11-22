package scraper_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestScraperSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Freshbooks scraper Suite")
}
