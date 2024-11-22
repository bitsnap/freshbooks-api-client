package client

import "github.com/bitsnap/bitsnap/packages/data-pipeline-general-accounting/cmd/bitsnap-accounting-freshbooks-codegen/codegen/client"

// GenerateListBills generates freshbooks client code to fetch all Bills
//
// API: https://api.freshbooks.com/accounting/account/<accountId>/bills/bills/
//
// Documentation: https://www.freshbooks.com/api/bills
func GenerateListBills() string {
	return client.GenerateList("Bills", "Bill", "/bills/bills")
}

// GenerateRetrieveBill generates freshbooks client code to retrieve specific Bill
//
// API: https://api.freshbooks.com/accounting/account/<accountId>/bills/bills/<billId>
//
// Documentation: https://www.freshbooks.com/api/bills
func GenerateRetrieveBill() string {
	return client.GenerateRetrieve("Bill", "/bills/bills")
}

// GenerateBillsModel generates freshbooks Bill domain model
//
// Documentation: https://www.freshbooks.com/api/bills
func GenerateBillsModel() string {
	return client.GenerateModel("Bill", "https://www.freshbooks.com/api/bills")
}
