package client

import "github.com/bitsnap/bitsnap/packages/data-pipeline-general-accounting/cmd/bitsnap-accounting-freshbooks-codegen/codegen/client"

// GenerateListInvoiceProfiles generates freshbooks client code to fetch all InvoiceProfiles
//
// API: https://api.freshbooks.com/accounting/account/<accountid>/invoice_profiles/invoice_profiles
//
// Documentation: https://www.freshbooks.com/api/invoice_profiles
func GenerateListInvoiceProfiles() string {
	return client.GenerateList("InvoiceProfiles", "InvoiceProfile", "/invoice_profiles/invoice_profiles")
}

// GenerateRetrieveInvoiceProfile generates freshbooks client code to retrieve specific InvoiceProfile
//
// API: https://api.freshbooks.com/accounting/account/<accountid>/invoice_profiles/invoice_profiles/<profileid>
//
// Documentation: https://www.freshbooks.com/api/invoice_profiles
func GenerateRetrieveInvoiceProfile() string {
	return client.GenerateRetrieve("InvoiceProfile", "/invoice_profiles/invoice_profiles")
}

// GenerateInvoiceProfilesModel generates freshbooks InvoiceProfile domain model
//
// Documentation: https://www.freshbooks.com/api/invoice_profiles
func GenerateInvoiceProfilesModel() string {
	return client.GenerateModel("InvoiceProfile", "https://www.freshbooks.com/api/invoice_profiles")
}
