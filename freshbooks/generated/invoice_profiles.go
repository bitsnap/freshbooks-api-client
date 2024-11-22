package generated

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"
	"time"

	"github.com/bitsnap/freshbooks-api-client/models"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type InvoiceProfile struct {
	Id                int64                               `json:"id"`                  // the unique id for the invoice profile
	OwnerId           int64                               `json:"ownerid"`             // id of creator of invoice. 1 if business admin, other if created by e.g. a contractor
	Updated           *time.Time                          `json:"updated"`             // the date the invoice profile was updated
	DiscountTotal     models.InvoiceProfilesDiscountTotal `json:"discount_total"`      // subfields: amount, code
	Amount            string                              `json:"amount"`              // the amount of the discount
	Code              string                              `json:"code"`                // three-letter currency code for the discount
	OccurrencesToDate int64                               `json:"occurrences_to_date"` // number of invoices that have been generated
	ProfileId         int64                               `json:"profileid"`           // same as id
	AutoBill          *bool                               `json:"auto_bill*"`          // Whether this invoice has a credit card saved
}

func ListInvoiceProfilesPageSortBy(accountId string, page uint64, sort *SortBy) ([]InvoiceProfile, *Pagination, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/invoice_profiles/invoice_profiles")
	if err != nil {
		return nil, nil, err
	}

	content, err := getQueryPage(client, parsedUrl, page, sort)
	if err != nil {
		return nil, nil, err
	}

	type InvoiceProfileResult struct {
		InvoiceProfiles []InvoiceProfile `json:"invoice_profiles"`
	}

	type InvoiceProfileResponseContent struct {
		Result InvoiceProfileResult `json:"result"`
		Pagination
	}

	type InvoiceProfileResponse struct {
		Response InvoiceProfileResponseContent `json:"response"`
	}

	entries := InvoiceProfileResponse{
		Response: InvoiceProfileResponseContent{
			Result: InvoiceProfileResult{
				InvoiceProfiles: make([]InvoiceProfile, 0, 10),
			},
		},
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, nil, err
	}

	if len(entries.Response.Result.InvoiceProfiles) == 0 {
		return []InvoiceProfile{}, nil, nil
	}

	return entries.Response.Result.InvoiceProfiles, &entries.Response.Pagination, nil
}

func ListInvoiceProfilesPage(site string, page uint64) ([]InvoiceProfile, *Pagination, error) {
	return ListInvoiceProfilesPageSortBy(site, page, nil)
}

func RetrieveInvoiceProfileById(accountId, id string) (*InvoiceProfile, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/invoice_profiles/invoice_profiles/" + id)
	if err != nil {
		return nil, err
	}

	content, err := getQueryPage(client, parsedUrl, -1, nil)
	if err != nil {
		return nil, err
	}

	type InvoiceProfileResult struct {
		InvoiceProfile InvoiceProfile `json:"invoice_profile"`
	}

	type InvoiceProfileResponseContent struct {
		Result InvoiceProfileResult `json:"result"`
	}

	type InvoiceProfileResponse struct {
		Response InvoiceProfileResponseContent `json:"response"`
	}

	entry := InvoiceProfileResponse{}

	err = json.Unmarshal(content, &entry)
	if err != nil {
		return nil, err
	}

	return &entry.Response.Result.InvoiceProfile, nil
}
