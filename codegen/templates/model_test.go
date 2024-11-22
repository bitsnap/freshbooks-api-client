package templates_test

import (
	"github.com/bitsnap/bitsnap/packages/data-pipeline-general-accounting/cmd/bitsnap-accounting-freshbooks-codegen/codegen/client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Model attributes code generation", func() {
	When("generating models from scraped documentation pages", func() {
		It("should be able to generate models from templates", func() {
			output := client.GenerateModel("Entitlement", "https://www.freshbooks.com/api/expenses")
			Expect(output).To(goldga.Match())
		})
	})
})
