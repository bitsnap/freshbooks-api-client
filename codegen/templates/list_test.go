package templates_test

import (
	"github.com/bitsnap/bitsnap/packages/data-pipeline-general-accounting/cmd/bitsnap-accounting-freshbooks-codegen/codegen/client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("List client endpoint code generation", func() {
	When("generating list endpoints client", func() {
		It("should be able to generate list methods", func() {
			output := client.GenerateList("Entitlements", "Entitlement", "")
			Expect(output).To(goldga.Match())
		})
	})
})
