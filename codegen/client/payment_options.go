package client

import "github.com/bitsnap/bitsnap/packages/data-pipeline-general-accounting/cmd/bitsnap-accounting-freshbooks-codegen/codegen/client"

// GenerateRetrievePaymentOptions generates freshbooks invoice code to retrieve invoice PaymentOptions
//
// API: https://api.freshbooks.com/payments/account/<accountid>/payment_options?entity_type=invoice
//
// Documentation: https://www.freshbooks.com/api/online-payments
func GenerateRetrievePaymentOptions() string {
	return client.GenerateList("PaymentOptions", "PaymentOption", "/payment_options?entity_type=invoice")
}
