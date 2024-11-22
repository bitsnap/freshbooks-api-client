package generated

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"
	"time"

	"github.com/bitsnap/freshbooks-api-client/models"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Payment struct {
	AccountingSystemId     string                `json:"accounting_systemid"`      // unique identifier of business client exists on
	Amount                 models.PaymentsAmount `json:"amount"`                   // amount paid on invoice
	BulkPaymentId          int64                 `json:"bulk_paymentid"`           //
	ClientId               int64                 `json:"clientid"`                 // id of client who made the payment
	CreditId               int64                 `json:"creditid"`                 // id of related credit
	Date                   string                `json:"date"`                     // date the payment was made, YYYY-MM-DD format
	FromCredit             bool                  `json:"from_credit"`              // whether or not the payment was converted from a Credit on a Client’s account
	Gateway                string                `json:"gateway"`                  // the payment processor used, if any
	Id                     int64                 `json:"id"`                       // unique id (across this business) for payment
	InvoiceId              int64                 `json:"invoiceid"`                // id of related invoice
	LogId                  int64                 `json:"logid"`                    // duplicate of id
	Note                   string                `json:"note"`                     // notes on payment, often used for credit card reference number
	OrderId                int64                 `json:"orderid"`                  // id of related orderid
	OverpaymentId          int64                 `json:"overpaymentid"`            // id of related overpayment if relevant
	SendClientNotification bool                  `json:"send_client_notification"` // whether to send the client a notification of this payment
	TransactionId          int64                 `json:"transactionid"`            // deprecated
	Type                   string                `json:"type"`                     // “Check”, “Credit”, “Cash”, etc.
	Updated                *time.Time            `json:"updated"`                  // date payment was last updated, YYYY-MM-DD
	VisState               uint8                 `json:"vis_state"`                // 0 for active, 1 for deleted
}

func ListPaymentsPageSortBy(accountId string, page uint64, sort *SortBy) ([]Payment, *Pagination, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/payments/payments")
	if err != nil {
		return nil, nil, err
	}

	content, err := getQueryPage(client, parsedUrl, page, sort)
	if err != nil {
		return nil, nil, err
	}

	type PaymentResult struct {
		Payments []Payment `json:"payments"`
	}

	type PaymentResponseContent struct {
		Result PaymentResult `json:"result"`
		Pagination
	}

	type PaymentResponse struct {
		Response PaymentResponseContent `json:"response"`
	}

	entries := PaymentResponse{
		Response: PaymentResponseContent{
			Result: PaymentResult{
				Payments: make([]Payment, 0, 10),
			},
		},
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, nil, err
	}

	if len(entries.Response.Result.Payments) == 0 {
		return []Payment{}, nil, nil
	}

	return entries.Response.Result.Payments, &entries.Response.Pagination, nil
}

func ListPaymentsPage(site string, page uint64) ([]Payment, *Pagination, error) {
	return ListPaymentsPageSortBy(site, page, nil)
}

func RetrievePaymentById(accountId, id string) (*Payment, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/payments/payments/" + id)
	if err != nil {
		return nil, err
	}

	content, err := getQueryPage(client, parsedUrl, -1, nil)
	if err != nil {
		return nil, err
	}

	type PaymentResult struct {
		Payment Payment `json:"payment"`
	}

	type PaymentResponseContent struct {
		Result PaymentResult `json:"result"`
	}

	type PaymentResponse struct {
		Response PaymentResponseContent `json:"response"`
	}

	entry := PaymentResponse{}

	err = json.Unmarshal(content, &entry)
	if err != nil {
		return nil, err
	}

	return &entry.Response.Result.Payment, nil
}
