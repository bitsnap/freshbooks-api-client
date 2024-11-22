package generated

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"
	"time"

	"github.com/bitsnap/freshbooks-api-client/models"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Estimate struct {
	EstimateId          string                        `json:"estimateid"`                      // Unique-to-this-system estimate id
	Id                  int64                         `json:"id"`                              // Unique-to-this-system estimate id, duplicate of estimateid
	AccountingSystemId  string                        `json:"accounting_systemid"`             // Unique id for system
	UiStatus            string                        `json:"ui_status"`                       // See Estimate UI Status table
	Status              string                        `json:"status"`                          // See Estimate Status table
	Amount              string                        `json:"amount"`                          // Total amount of estimate, to two decimal places
	Code                string                        `json:"code"`                            // Three-letter currency code
	DiscountTotal       models.EstimatesDiscountTotal `json:"discount_total"`                  // subfields: Amount and code
	Description         string                        `json:"description"`                     // Description of first line of estimate
	CurrentOrganization string                        `json:"current_organization"`            // Name of organization being estimated — denormalized from client
	Invoiced            bool                          `json:"invoiced"`                        // Indicator of whether this estimate has been converted to an invoice that was sent
	OwnerId             int64                         `json:"ownerid"`                         // Id of creator of estimate. 1 if business admin, other if created by e.g. a contractor
	SentId              int64                         `json:"sentid"`                          // Userid of user who sent the estimate, typically 1 for admin
	CreatedAt           *time.Time                    `json:"created_at"`                      // Time the estimate was created, YYYY-MM-DD HH:MM:SS format
	Updated             *time.Time                    `json:"updated"`                         // Time estimate last updated at, YYYY-MM-DD HH:MM:SS format
	DisplayStatus       string                        `json:"display_status"`                  // Description of status shown in FreshBooks UI, one of ‘draft’, ‘sent’, or ‘viewed’
	ReplyStatus         string                        `json:"reply_status"`                    // (Deprecated) Description of status shown in Classic FreshBooks’ UI, one of ‘replied’ or ‘resolved’
	EstimateNumber      string                        `json:"estimate_number"`                 // User-specified and visible estimate id
	CustomerId          int64                         `json:"customerid" validate:"required"`  // Unique-to-this-system client-id
	CreateDate          *time.Time                    `json:"create_date" validate:"required"` // Date estimate was created, YYYY-MM-DD format
	DiscountValue       string                        `json:"discount_value"`                  // Decimal-string amount
	PoNumber            string                        `json:"po_number"`                       // Post Office box number for address on estimate
	Template            string                        `json:"template"`                        // (Internal, deprecated) choice of rendering style
	CurrencyCode        string                        `json:"currency_code"`                   // Three-letter currency code for estimate
	Language            string                        `json:"language"`                        // Two-letter language code, e.g. “en”
	Terms               string                        `json:"terms"`                           // Terms listed on estimate
	Notes               string                        `json:"notes"`                           // Notes listed on estimate
	Address             string                        `json:"address"`                         // First line of address on estimate
	ExtArchive          int64                         `json:"ext_archive"`                     // (deprecated) 0 or 1 indicating archived or not
	VisState            uint8                         `json:"vis_state"`                       // 0 for active, 1 for deleted
	Street              string                        `json:"street"`                          // Street for address on estimate
	Street2             string                        `json:"street2"`                         // Second line of street for address on estimate
	City                string                        `json:"city"`                            // City for address on estimate
	Province            string                        `json:"province"`                        // Province for address on estimate
	Country             string                        `json:"country"`                         // Country for address on estimate
	Organization        string                        `json:"organization"`                    // Name of organization being estimated
	FirstName           string                        `json:"fname"`                           // First name of Client on estimate
	LastName            string                        `json:"lname"`                           // Last name of client being estimated
	VatName             string                        `json:"vat_name"`                        // Value Added Tax name if provided
	VatNumber           string                        `json:"vat_number"`                      // Value Added Tax number if provided
	Lines               []models.EstimatesLines       `json:"lines"`                           // Lines of the estimate
}

func ListEstimatesPageSortBy(accountId string, page uint64, sort *SortBy) ([]Estimate, *Pagination, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/estimates/estimates")
	if err != nil {
		return nil, nil, err
	}

	content, err := getQueryPage(client, parsedUrl, page, sort)
	if err != nil {
		return nil, nil, err
	}

	type EstimateResult struct {
		Estimates []Estimate `json:"estimates"`
	}

	type EstimateResponseContent struct {
		Result EstimateResult `json:"result"`
		Pagination
	}

	type EstimateResponse struct {
		Response EstimateResponseContent `json:"response"`
	}

	entries := EstimateResponse{
		Response: EstimateResponseContent{
			Result: EstimateResult{
				Estimates: make([]Estimate, 0, 10),
			},
		},
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, nil, err
	}

	if len(entries.Response.Result.Estimates) == 0 {
		return []Estimate{}, nil, nil
	}

	return entries.Response.Result.Estimates, &entries.Response.Pagination, nil
}

func ListEstimatesPage(site string, page uint64) ([]Estimate, *Pagination, error) {
	return ListEstimatesPageSortBy(site, page, nil)
}

func RetrieveEstimateById(accountId, id string) (*Estimate, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/estimates/estimates/" + id)
	if err != nil {
		return nil, err
	}

	content, err := getQueryPage(client, parsedUrl, -1, nil)
	if err != nil {
		return nil, err
	}

	type EstimateResult struct {
		Estimate Estimate `json:"estimate"`
	}

	type EstimateResponseContent struct {
		Result EstimateResult `json:"result"`
	}

	type EstimateResponse struct {
		Response EstimateResponseContent `json:"response"`
	}

	entry := EstimateResponse{}

	err = json.Unmarshal(content, &entry)
	if err != nil {
		return nil, err
	}

	return &entry.Response.Result.Estimate, nil
}
