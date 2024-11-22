package generated

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"
	"time"

	"github.com/bitsnap/freshbooks-api-client/models"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Bill struct {
	Amount             models.BillsAmount         `json:"amount"`              // Read-only. Computed from lines. subfields: amount and code
	Attachment         models.BillsAttachment     `json:"attachment"`          //
	BillNumber         string                     `json:"bill_number"`         // reference to vendor bill number
	BillPayments       []models.BillsBillPayments `json:"bill_payments"`       // Bill Payments made against the bill
	CreatedAt          *time.Time                 `json:"created_at"`          // Read-only. time the invoice was created, YYYY-MM-DD HH:MM:SS format
	CurrencyCode       string                     `json:"currency_code"`       // three-letter currency code
	DueDate            *time.Time                 `json:"due_date"`            // date for which the bill is due for payment
	DueOffsetDays      int64                      `json:"due_offset_days"`     // number of days from the issue date that the invoice needs to be set to due
	Id                 int64                      `json:"id"`                  // uniquely identifies the bill associated with the business.
	IssueDate          *time.Time                 `json:"issue_date"`          // date when the bill was issued by the vendor
	Language           string                     `json:"language"`            // two-letter language code, e.g. “en”
	Lines              []models.BillsLines        `json:"lines"`               // array of bill line items, refer to bill line items object below
	Outstanding        models.BillsOutstanding    `json:"outstanding"`         // Read-only. subfields: amount and code
	OverallCategory    string                     `json:"overall_category"`    // Read-only. If multiple categories are selected in the bill lines, then overall_category is Split. Otherwise, it will be the selected category.
	OverallDescription string                     `json:"overall_description"` // Read-only. First non-null value of bill line descriptions
	Paid               models.BillsPaid           `json:"paid"`                // Read-only. subfields: amount and code
	Status             string                     `json:"status"`              // Read-only. Status of the bill: “unpaid”, “overdue”, “partial”, “paid”
	TaxAmount          models.BillsTaxAmount      `json:"tax_amount"`          // Read-only. Computed from lines. subfields: amount and code
	TotalAmount        models.BillsTotalAmount    `json:"total_amount"`        // Read-only. Computed from lines. subfields: amount and code
	UpdatedAt          string                     `json:"updated_at"`          // Read-only. last time the resource was updated. YYYY-MM-DD HH:MM:SS format
	VendorId           int64                      `json:"vendorid"`            //
	VisState           uint8                      `json:"vis_state"`           // 0 for active, 1 for deleted, 2 for archived
}

func ListBillsPageSortBy(accountId string, page uint64, sort *SortBy) ([]Bill, *Pagination, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/bills/bills")
	if err != nil {
		return nil, nil, err
	}

	content, err := getQueryPage(client, parsedUrl, page, sort)
	if err != nil {
		return nil, nil, err
	}

	type BillResult struct {
		Bills []Bill `json:"bills"`
	}

	type BillResponseContent struct {
		Result BillResult `json:"result"`
		Pagination
	}

	type BillResponse struct {
		Response BillResponseContent `json:"response"`
	}

	entries := BillResponse{
		Response: BillResponseContent{
			Result: BillResult{
				Bills: make([]Bill, 0, 10),
			},
		},
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, nil, err
	}

	if len(entries.Response.Result.Bills) == 0 {
		return []Bill{}, nil, nil
	}

	return entries.Response.Result.Bills, &entries.Response.Pagination, nil
}

func ListBillsPage(site string, page uint64) ([]Bill, *Pagination, error) {
	return ListBillsPageSortBy(site, page, nil)
}

func RetrieveBillById(accountId, id string) (*Bill, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/bills/bills/" + id)
	if err != nil {
		return nil, err
	}

	content, err := getQueryPage(client, parsedUrl, -1, nil)
	if err != nil {
		return nil, err
	}

	type BillResult struct {
		Bill Bill `json:"bill"`
	}

	type BillResponseContent struct {
		Result BillResult `json:"result"`
	}

	type BillResponse struct {
		Response BillResponseContent `json:"response"`
	}

	entry := BillResponse{}

	err = json.Unmarshal(content, &entry)
	if err != nil {
		return nil, err
	}

	return &entry.Response.Result.Bill, nil
}
