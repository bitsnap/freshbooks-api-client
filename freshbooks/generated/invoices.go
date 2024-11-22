package generated

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"
	"time"

	"github.com/bitsnap/freshbooks-api-client/models"
	"github.com/shopspring/decimal"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Invoice struct {
	InvoiceId           int64                        `json:"invoiceid"`                       // unique-to-this-system invoice id
	Id                  int64                        `json:"id"`                              // unique-to-this-system invoice id, duplicate of invoiceid
	AccountingSystemId  string                       `json:"accounting_systemid"`             // unique id for system
	AccountId           string                       `json:"accountid"`                       // unique id for system, repeat of accounting_systemid
	Amount              models.InvoicesAmount        `json:"amount"`                          // total amount of invoice. subfields: amount, code
	Paid                models.InvoicesPaid          `json:"paid"`                            // subfields: amount and code
	Outstanding         models.InvoicesOutstanding   `json:"outstanding"`                     // subfields: amount, code
	DiscountTotal       models.InvoicesDiscountTotal `json:"discount_total"`                  // subfields: amount and code
	CreatedAt           *time.Time                   `json:"created_at"`                      // Time the invoice was created, YYYY-MM-DD HH:MM:SS format
	CurrentOrganization string                       `json:"current_organization"`            // Name of organization being invoiced — denormalized from client
	DatePaid            *time.Time                   `json:"date_paid"`                       // date invoice was fully paid, YYYY-MM-DD format
	Description         string                       `json:"description"`                     // description of first line of invoice
	DueDate             *time.Time                   `json:"due_date"`                        // date invoice is marked as due by, YYYY-MM-DD format, calculated from due_offset_days. If due_offset_days is not set, it will default to the date of issue.
	Updated             *time.Time                   `json:"updated"`                         // time invoice last updated at, YYYY-MM-DD HH:MM:SS format
	OwnerId             int64                        `json:"ownerid"`                         // id of creator of invoice. 1 if business admin, other if created by e.g. a contractor
	EstimateId          int64                        `json:"estimateid"`                      // id of associated estimate, 0 if none
	SentId              int64                        `json:"sentid"`                          // userid of user who sent the invoice, typically 1 for admin
	Status              string                       `json:"status"`                          // See Invoice Status table.
	Parent              int64                        `json:"parent"`                          // id of object this invoice was generated from, 0 if none
	DisplayStatus       string                       `json:"display_status"`                  // Description of status shown in FreshBooks UI, one of ‘draft’, ‘created’, ‘sent’, ‘viewed’, or ‘outstanding’.
	AutobillStatus      string                       `json:"autobill_status"`                 // one of retry, failed, or success
	PaymentStatus       string                       `json:"payment_status"`                  // description of payment status. One of ‘unpaid’, ‘partial’, ‘paid’, and ‘auto-paid’. See the v3_status table on this page for descriptions of each.
	LastOrderStatus     string                       `json:"last_order_status"`               // describes status of last attempted payment
	DisputeStatus       string                       `json:"dispute_status"`                  // description of whether invoice has been disputed.
	DepositStatus       string                       `json:"deposit_status"`                  // description of deposits applied to invoice. One of ‘paid’, ‘unpaid’, ‘partial’, ‘none’, and ‘converted’.
	AutoBill            bool                         `json:"auto_bill"`                       // whether this invoice has a credit card saved
	V3Status            string                       `json:"v3_status"`                       // description of Invoice Status, see V3 Status Table.
	InvoiceNumber       string                       `json:"invoice_number"`                  // user-specified and visible invoice id
	CustomerId          int64                        `json:"customerid" validate:"required"`  // unique-to-this-system client-id
	CreateDate          string                       `json:"create_date" validate:"required"` // date invoice was created, YYYY-MM-DD (string) format
	GenerationDate      *time.Time                   `json:"generation_date"`                 // date invoice generated from object, null if it wasn’t, YYYY-MM-DD (string) if it was
	DiscountValue       *decimal.Decimal             `json:"discount_value"`                  // percent amount being discounted from the subtotal, decimal-string amount ranging from 0 to 100
	DiscountDescription string                       `json:"discount_description"`            // public note about discount
	PoNumber            string                       `json:"po_number"`                       // Reference number for address on invoice.
	Template            string                       `json:"template"`                        // (internal, deprecated) choice of rendering style
	CurrencyCode        string                       `json:"currency_code"`                   // three-letter currency code for invoice
	Language            string                       `json:"language"`                        // two-letter language code, e.g. “en”
	Terms               string                       `json:"terms"`                           // terms listed on invoice
	Notes               string                       `json:"notes"`                           // Notes listed on invoice
	Address             string                       `json:"address"`                         // First line of address on invoice
	ReturnUri           string                       `json:"return_uri"`                      // (deprecated)
	DepositAmount       models.InvoicesDepositAmount `json:"deposit_amount"`                  // amount required as deposit, null if none
	DepositPercentage   string                       `json:"deposit_percentage"`              // percent of the invoice’s value required as a deposit
	Gmail               bool                         `json:"gmail"`                           // whether to send via ground mail
	ShowAttachments     bool                         `json:"show_attachments"`                // whether attachments on invoice are rendered
	ExtArchive          int64                        `json:"ext_archive"`                     // (deprecated) 0 or 1 indicating archived or not
	VisState            uint8                        `json:"vis_state"`                       // “visibility state”. 0 for active, 1 for deleted, 2 for archived
	Street              string                       `json:"street"`                          // street for address on invoice
	Street2             string                       `json:"street2"`                         // second line of street for address on invoice.
	City                string                       `json:"city"`                            // city for address on invoice
	Province            string                       `json:"province"`                        // Province for address on invoice.
	Code                string                       `json:"code"`                            // zip code for address on invoice
	Country             string                       `json:"country"`                         // Country for address on invoice
	Organization        string                       `json:"organization"`                    // Name of organization being invoiced.
	FirstName           string                       `json:"fname"`                           // First name of Client on invoice
	LastName            string                       `json:"lname"`                           // Last name of client being invoiced
	VatName             string                       `json:"vat_name"`                        // Value Added Tax name if provided
	VatNumber           string                       `json:"vat_number"`                      // Value Added Tax number if provided
	DueOffsetDays       int64                        `json:"due_offset_days"`                 // Number of days from creation that invoice is due. If not set, the due date will default to the date of issue.
	Lines               []models.InvoicesLines       `json:"lines"`                           // Lines of the invoice
	Presentation        models.InvoicesPresentation  `json:"presentation"`                    // where invoice logo and styles are defined. See our postman collection for details.
}

func ListInvoicesPageSortBy(accountId string, page uint64, sort *SortBy) ([]Invoice, *Pagination, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/invoices/invoices")
	if err != nil {
		return nil, nil, err
	}

	content, err := getQueryPage(client, parsedUrl, page, sort)
	if err != nil {
		return nil, nil, err
	}

	type InvoiceResult struct {
		Invoices []Invoice `json:"invoices"`
	}

	type InvoiceResponseContent struct {
		Result InvoiceResult `json:"result"`
		Pagination
	}

	type InvoiceResponse struct {
		Response InvoiceResponseContent `json:"response"`
	}

	entries := InvoiceResponse{
		Response: InvoiceResponseContent{
			Result: InvoiceResult{
				Invoices: make([]Invoice, 0, 10),
			},
		},
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, nil, err
	}

	if len(entries.Response.Result.Invoices) == 0 {
		return []Invoice{}, nil, nil
	}

	return entries.Response.Result.Invoices, &entries.Response.Pagination, nil
}

func ListInvoicesPage(site string, page uint64) ([]Invoice, *Pagination, error) {
	return ListInvoicesPageSortBy(site, page, nil)
}

func RetrieveInvoiceById(accountId, id string) (*Invoice, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/invoices/invoices/" + id)
	if err != nil {
		return nil, err
	}

	content, err := getQueryPage(client, parsedUrl, -1, nil)
	if err != nil {
		return nil, err
	}

	type InvoiceResult struct {
		Invoice Invoice `json:"invoice"`
	}

	type InvoiceResponseContent struct {
		Result InvoiceResult `json:"result"`
	}

	type InvoiceResponse struct {
		Response InvoiceResponseContent `json:"response"`
	}

	entry := InvoiceResponse{}

	err = json.Unmarshal(content, &entry)
	if err != nil {
		return nil, err
	}

	return &entry.Response.Result.Invoice, nil
}
