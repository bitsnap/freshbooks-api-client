package generated

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"
	"time"

	"github.com/shopspring/decimal"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Tax struct {
	AccountingSystemId string           `json:"accounting_systemid"`      // unique identifier of business tax exists on
	Updated            *time.Time       `json:"updated"`                  // date object was last updated, YYYY-MM-DD HH:MM:SS format
	Name               string           `json:"name" validate:"required"` // identifiable name for your tax, e.g. “GST”
	Number             string           `json:"number"`                   // an external number that identifies your tax submission
	TaxId              int64            `json:"taxid"`                    // unique identifier of this tax within this business
	Amount             *decimal.Decimal `json:"amount"`                   // string-decimal representing percentage value of tax
	Compound           bool             `json:"compound"`                 // compound taxes are calculated on top of primary taxes
	Id                 int64            `json:"id"`                       // unique id to look this tax up later
}

func ListTaxesPageSortBy(accountId string, page uint64, sort *SortBy) ([]Tax, *Pagination, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/taxes/taxes")
	if err != nil {
		return nil, nil, err
	}

	content, err := getQueryPage(client, parsedUrl, page, sort)
	if err != nil {
		return nil, nil, err
	}

	type TaxResult struct {
		Taxes []Tax `json:"taxes"`
	}

	type TaxResponseContent struct {
		Result TaxResult `json:"result"`
		Pagination
	}

	type TaxResponse struct {
		Response TaxResponseContent `json:"response"`
	}

	entries := TaxResponse{
		Response: TaxResponseContent{
			Result: TaxResult{
				Taxes: make([]Tax, 0, 10),
			},
		},
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, nil, err
	}

	if len(entries.Response.Result.Taxes) == 0 {
		return []Tax{}, nil, nil
	}

	return entries.Response.Result.Taxes, &entries.Response.Pagination, nil
}

func ListTaxesPage(site string, page uint64) ([]Tax, *Pagination, error) {
	return ListTaxesPageSortBy(site, page, nil)
}

func RetrieveTaxById(accountId, id string) (*Tax, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/taxes/taxes/" + id)
	if err != nil {
		return nil, err
	}

	content, err := getQueryPage(client, parsedUrl, -1, nil)
	if err != nil {
		return nil, err
	}

	type TaxResult struct {
		Tax Tax `json:"tax"`
	}

	type TaxResponseContent struct {
		Result TaxResult `json:"result"`
	}

	type TaxResponse struct {
		Response TaxResponseContent `json:"response"`
	}

	entry := TaxResponse{}

	err = json.Unmarshal(content, &entry)
	if err != nil {
		return nil, err
	}

	return &entry.Response.Result.Tax, nil
}
