package generated

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/bitsnap/freshbooks-api-client/generated/enums"
	"github.com/bitsnap/freshbooks-api-client/generated/models"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type CreditNote struct {
	CreditNumber string                  `json:"credit_number"` // name of the credit note
	CurrencyCode string                  `json:"currency_code"` // abbreviated currency code identified type of currency
	CreateDate   *time.Time              `json:"create_date"`   // date the credit was created
	Notes        string                  `json:"notes"`         // notes associated with the creation of the credit note.
	CreditType   enums.CreditsCreditType `json:"credit_type"`   // defines type of credit (i.e. “goodwill” or “prepayment”
	Terms        string                  `json:"terms"`         // terms of the credit note
	Language     enums.CreditsLanguage   `json:"language"`      // language the credit is received in. ie. “en”
	Lines        []models.CreditsLines   `json:"lines"`         // line item of a credit. Each lines array that is created provides an additional line item on the credit.
	ClientId     string                  `json:"clientid"`      // identifier for client the note is associated with. For information on how to find the clientid, check the “Client” section.
}

func ListCreditNotesPageSortBy(accountId string, page uint64, sort *SortBy) ([]CreditNote, *Pagination, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/credit_notes/credit_notes")
	if err != nil {
		return nil, nil, err
	}

	content, err := getQueryPage(client, parsedUrl, page, sort)
	if err != nil {
		return nil, nil, err
	}

	type CreditNoteResult struct {
		CreditNotes []CreditNote `json:"credit_notes"`
	}

	type CreditNoteResponseContent struct {
		Result CreditNoteResult `json:"result"`
		Pagination
	}

	type CreditNoteResponse struct {
		Response CreditNoteResponseContent `json:"response"`
	}

	entries := CreditNoteResponse{
		Response: CreditNoteResponseContent{
			Result: CreditNoteResult{
				CreditNotes: make([]CreditNote, 0, 10),
			},
		},
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, nil, err
	}

	if len(entries.Response.Result.CreditNotes) == 0 {
		return []CreditNote{}, nil, nil
	}

	return entries.Response.Result.CreditNotes, &entries.Response.Pagination, nil
}

func ListCreditNotesPage(site string, page uint64) ([]CreditNote, *Pagination, error) {
	return ListCreditNotesPageSortBy(site, page, nil)
}

func RetrieveCreditNoteById(accountId, id string) (*CreditNote, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/credit_notes/credit_notes/" + id)
	if err != nil {
		return nil, err
	}

	content, err := getQueryPage(client, parsedUrl, -1, nil)
	if err != nil {
		return nil, err
	}

	type CreditNoteResult struct {
		CreditNote CreditNote `json:"credit_note"`
	}

	type CreditNoteResponseContent struct {
		Result CreditNoteResult `json:"result"`
	}

	type CreditNoteResponse struct {
		Response CreditNoteResponseContent `json:"response"`
	}

	entry := CreditNoteResponse{}

	err = json.Unmarshal(content, &entry)
	if err != nil {
		return nil, err
	}

	return &entry.Response.Result.CreditNote, nil
}
