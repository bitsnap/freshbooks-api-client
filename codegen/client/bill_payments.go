package client

import (
	"github.com/bitsnap/bitsnap/packages/data-pipeline-general-accounting/cmd/bitsnap-accounting-freshbooks-codegen/codegen/client"
)

// GenerateListBillPayments generates freshbooks client code to fetch all bill payments
//
// API: https://api.freshbooks.com/accounting/account/<accountId>/bill_payments/bill_payments
//
// Documentation: https://www.freshbooks.com/api/bill_payments
func GenerateListBillPayments() string {
	return client.GenerateList("BillPayments", "BillPayment", "/bill_payments/bill_payments")
}

// GenerateRetrieveBillPayment generates freshbooks client code to retrieve specific client
//
// API:https://api.freshbooks.com/accounting/account/<accountId>/bill_payments/bill_payments/<payment-id>
//
// Documentation: https://www.freshbooks.com/api/bill_payments
func GenerateRetrieveBillPayment() string {
	return client.GenerateRetrieve("BillPayment", "/bill_payments/bill_payments")
}

// GenerateBillPaymentsModel generates freshbooks BillPayment domain model
//
// Documentation:
func GenerateBillPaymentsModel() string {
	return client.GenerateModel("BillPayment", "https://www.freshbooks.com/api/bill_payments")
}
