package client

import "github.com/bitsnap/bitsnap/packages/data-pipeline-general-accounting/cmd/bitsnap-accounting-freshbooks-codegen/codegen/client"

// GenerateListBillVendors generates freshbooks client code to fetch all BillVendors
//
// API: https://api.freshbooks.com/accounting/account/<accountId>/bill_vendors/bill_vendors
//
// Documentation: https://www.freshbooks.com/api/vendors
func GenerateListBillVendors() string {
	return client.GenerateList("BillVendors", "BillVendor", "/bill_vendors/bill_vendors")
}

// GenerateBillVendorsModel generates freshbooks BillVendor domain model
//
// Documentation: https://www.freshbooks.com/api/vendors
func GenerateBillVendorsModel() string {
	return client.GenerateModel("BillVendor", "https://www.freshbooks.com/api/vendors")
}
