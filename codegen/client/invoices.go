package client

import "github.com/bitsnap/bitsnap/packages/data-pipeline-general-accounting/cmd/bitsnap-accounting-freshbooks-codegen/codegen/client"

// GenerateListInvoices generates freshbooks invoice code to fetch all invoices
//
// API: https://api.freshbooks.com/accounting/account/<accountId>/invoices/invoices
//
// Documentation: https://www.freshbooks.com/api/invoices
func GenerateListInvoices() string {
	return client.GenerateList("Invoices", "Invoice", "/invoices/invoices")
}

// GenerateRetrieveInvoice generates freshbooks invoice code to retrieve specific invoice
//
// API: https://api.freshbooks.com/accounting/account/<accountId>/invoices/invoices/<invoiceId>
//
// Documentation: https://www.freshbooks.com/api/invoices
func GenerateRetrieveInvoice() string {
	return client.GenerateRetrieve("Invoice", "/invoices/invoices")
}

// GenerateInvoicesModel generates freshbooks Invoice domain model
//
// Documentation: https://www.freshbooks.com/api/invoices
func GenerateInvoicesModel() string {
	return client.GenerateModel("Invoice", "https://www.freshbooks.com/api/invoices")
}
