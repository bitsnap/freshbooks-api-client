package generated

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"
	"time"

	"github.com/bitsnap/freshbooks-api-client/models"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type BillVendor struct {
	AccountNumber           string                           `json:"account_number"`             // account number of the vendor
	City                    string                           `json:"city"`                       // city of vendor
	Country                 string                           `json:"country"`                    // country of vendor
	CreatedAt               *time.Time                       `json:"created_at"`                 // Read-only. time the vendor was created, YYYY-MM-DD HH:MM:SS format
	CurrencyCode            string                           `json:"currency_code"`              // default three-letter currency code for vendor
	Is1099                  bool                             `json:"is_1099"`                    // set true if vendor is a 1099 contractor
	Language                string                           `json:"language"`                   // two-letter language code, e.g. “en”
	Note                    string                           `json:"note"`                       //
	OutstandingBalance      models.VendorsOutstandingBalance `json:"outstanding_balance"`        // Read-only. subfields: amount and code
	OverdueBalance          models.VendorsOverdueBalance     `json:"overdue_balance"`            // Read-only. subfields: amount and code
	Phone                   string                           `json:"phone"`                      // phone number
	PostalCode              string                           `json:"postal_code"`                // postal code
	PrimaryContactEmail     string                           `json:"primary_contact_email"`      // vendor primary email
	PrimaryContactFirstName string                           `json:"primary_contact_first_name"` // vendor primary first name
	PrimaryContactLastName  string                           `json:"primary_contact_last_name"`  // vendor primary last name
	Province                string                           `json:"province"`                   // province
	Street                  string                           `json:"street"`                     // street address
	Street2                 string                           `json:"street2"`                    // street address 2nd part
	TaxDefaults             models.VendorsTaxDefaults        `json:"tax_defaults"`               // array of “bill_vendor_tax_defaults” object, see below
	UpdatedAt               *time.Time                       `json:"updated_at"`                 // Read-only. time of last modification to resource
	VendorName              string                           `json:"vendor_name"`                // vendor name
	VendorId                int64                            `json:"vendorid"`                   // unique identifier for vendor
	VisState                string                           `json:"vis_state"`                  // visibility state, possible values are 0, 1, 2
	Website                 string                           `json:"website"`                    // vendor website address
}

func ListBillVendorsPageSortBy(accountId string, page uint64, sort *SortBy) ([]BillVendor, *Pagination, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/bill_vendors/bill_vendors")
	if err != nil {
		return nil, nil, err
	}

	content, err := getQueryPage(client, parsedUrl, page, sort)
	if err != nil {
		return nil, nil, err
	}

	type BillVendorResult struct {
		BillVendors []BillVendor `json:"bill_vendors"`
	}

	type BillVendorResponseContent struct {
		Result BillVendorResult `json:"result"`
		Pagination
	}

	type BillVendorResponse struct {
		Response BillVendorResponseContent `json:"response"`
	}

	entries := BillVendorResponse{
		Response: BillVendorResponseContent{
			Result: BillVendorResult{
				BillVendors: make([]BillVendor, 0, 10),
			},
		},
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, nil, err
	}

	if len(entries.Response.Result.BillVendors) == 0 {
		return []BillVendor{}, nil, nil
	}

	return entries.Response.Result.BillVendors, &entries.Response.Pagination, nil
}

func ListBillVendorsPage(site string, page uint64) ([]BillVendor, *Pagination, error) {
	return ListBillVendorsPageSortBy(site, page, nil)
}
