package client

import "github.com/bitsnap/bitsnap/packages/data-pipeline-general-accounting/cmd/bitsnap-accounting-freshbooks-codegen/codegen/client"

// GenerateListCreditNotes generates freshbooks client code to fetch all CreditNotes
//
// API: https://api.freshbooks.com/accounting/account/<accountid>/credit_notes/credit_notes
//
// Documentation: https://www.freshbooks.com/api/credits
func GenerateListCreditNotes() string {
	return client.GenerateList("CreditNotes", "CreditNote", "/credit_notes/credit_notes")
}

// GenerateRetrieveCreditNote generates freshbooks client code to retrieve specific CreditNote
//
// API:https://api.freshbooks.com/accounting/account/<accountId>/credit_notes/credit_notes/<credit-id>
//
// Documentation: https://www.freshbooks.com/api/credits
func GenerateRetrieveCreditNote() string {
	return client.GenerateRetrieve("CreditNote", "/credit_notes/credit_notes")
}

// GenerateCreditNotesModel generates freshbooks CreditNote domain model
//
// Documentation: https://www.freshbooks.com/api/credits
func GenerateCreditNotesModel() string {
	return client.GenerateModel("CreditNote", "https://www.freshbooks.com/api/credits")
}
