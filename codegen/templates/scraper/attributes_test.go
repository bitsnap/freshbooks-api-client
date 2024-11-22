package scraper_test

import (
	"strings"

	"github.com/bitsnap/bitsnap/packages/data-pipeline-general-accounting/cmd/bitsnap-accounting-freshbooks-codegen/codegen/scraper"
	"github.com/tommy351/goldga"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// golden files helper
type attributes []*scraper.Attribute

func (a attributes) String() string {
	sb := strings.Builder{}
	for _, attr := range a {
		sb.WriteString(attr.String())
		sb.WriteRune(',')
		sb.WriteRune(' ')
	}

	return sb.String()
}

var _ = Describe("Model attributes scraping", func() {
	When("traversing documentation pages", func() {
		It("should be able to find attribute tables with specified header text", func() {
			attrs, err := scraper.ScrapeAttributes("https://www.freshbooks.com/api/expenses")
			Expect(err).To(Not(HaveOccurred()))
			Expect(attributes(attrs)).NotTo(BeEmpty())
			Expect(attributes(attrs)).To(goldga.Match())
		})
	})
})
