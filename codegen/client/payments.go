package client

import "github.com/bitsnap/bitsnap/packages/data-pipeline-general-accounting/cmd/bitsnap-accounting-freshbooks-codegen/codegen/client"

// GenerateListPayments generates freshbooks client code to fetch all payments
//
// API: https://api.freshbooks.com/accounting/account/<accountId>/payments/payments
//
// Documentation: https://www.freshbooks.com/api/payments
func GenerateListPayments() string {
	return client.GenerateList("Payments", "Payment", "/payments/payments")
}

// GenerateRetrievePayment generates freshbooks client code to retrieve specific payment
//
// API: https://api.freshbooks.com/accounting/account/<accountid>/payments/payments/<id>
//
// Documentation: https://www.freshbooks.com/api/payments
func GenerateRetrievePayment() string {
	return client.GenerateRetrieve("Payment", "/payments/payments")
}

// GeneratePaymentsModel generates freshbooks Payment domain model
//
// Documentation: https://www.freshbooks.com/api/payments
func GeneratePaymentsModel() string {
	return client.GenerateModel("Payment", "https://www.freshbooks.com/api/payments")
}
