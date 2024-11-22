package client

import "github.com/bitsnap/bitsnap/packages/data-pipeline-general-accounting/cmd/bitsnap-accounting-freshbooks-codegen/codegen/client"

// GenerateListTaxes generates freshbooks client code to fetch all taxes
//
// API: https://api.freshbooks.com/accounting/account/<accountid>/taxes/taxes
//
// Documentation: https://www.freshbooks.com/api/taxes
func GenerateListTaxes() string {
	return client.GenerateList("Taxes", "Tax", "/taxes/taxes")
}

// GenerateRetrieveTax generates freshbooks client code to retrieve specific payment
//
// API: https://api.freshbooks.com/accounting/account/<accountid>/taxes/taxes/<id>
//
// Documentation: https://www.freshbooks.com/api/taxes
func GenerateRetrieveTax() string {
	return client.GenerateRetrieve("Tax", "/taxes/taxes")
}

// GenerateTaxModel generates freshbooks Tax domain model
//
// Documentation: https://www.freshbooks.com/api/taxes
func GenerateTaxModel() string {
	return client.GenerateModel("Tax", "https://www.freshbooks.com/api/taxes")
}
