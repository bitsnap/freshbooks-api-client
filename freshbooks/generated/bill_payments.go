package generated

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"
	"time"

	"github.com/bitsnap/freshbooks-api-client/models"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type BillPayment struct {
	Id                 int64                     `json:"id"`                   // uniquely identifies the payment associated with the business.
	Amount             models.BillPaymentsAmount `json:"amount"`               // subfields: amount and code
	BillId             int64                     `json:"billid"`               // id of the bill to create the payment for
	MatchedWithExpense bool                      `json:"matched_with_expense"` // internal use. If the Bill has been reconciled with a bank-imported expense.
	PaidDate           *time.Time                `json:"paid_date"`            // date the payment was made, YYYY-MM-DD format
	PaymentType        string                    `json:"payment_type"`         // “Check”, “Credit”, “Cash”, etc.
	Note               string                    `json:"note"`                 // notes on payment
	VisState           uint8                     `json:"vis_state"`            // 0 for active, 1 for deleted, 2 for archived
}

func ListBillPaymentsPageSortBy(accountId string, page uint64, sort *SortBy) ([]BillPayment, *Pagination, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/bill_payments/bill_payments")
	if err != nil {
		return nil, nil, err
	}

	content, err := getQueryPage(client, parsedUrl, page, sort)
	if err != nil {
		return nil, nil, err
	}

	type BillPaymentResult struct {
		BillPayments []BillPayment `json:"bill_payments"`
	}

	type BillPaymentResponseContent struct {
		Result BillPaymentResult `json:"result"`
		Pagination
	}

	type BillPaymentResponse struct {
		Response BillPaymentResponseContent `json:"response"`
	}

	entries := BillPaymentResponse{
		Response: BillPaymentResponseContent{
			Result: BillPaymentResult{
				BillPayments: make([]BillPayment, 0, 10),
			},
		},
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, nil, err
	}

	if len(entries.Response.Result.BillPayments) == 0 {
		return []BillPayment{}, nil, nil
	}

	return entries.Response.Result.BillPayments, &entries.Response.Pagination, nil
}

func ListBillPaymentsPage(site string, page uint64) ([]BillPayment, *Pagination, error) {
	return ListBillPaymentsPageSortBy(site, page, nil)
}

func RetrieveBillPaymentById(accountId, id string) (*BillPayment, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/bill_payments/bill_payments/" + id)
	if err != nil {
		return nil, err
	}

	content, err := getQueryPage(client, parsedUrl, -1, nil)
	if err != nil {
		return nil, err
	}

	type BillPaymentResult struct {
		BillPayment BillPayment `json:"bill_payment"`
	}

	type BillPaymentResponseContent struct {
		Result BillPaymentResult `json:"result"`
	}

	type BillPaymentResponse struct {
		Response BillPaymentResponseContent `json:"response"`
	}

	entry := BillPaymentResponse{}

	err = json.Unmarshal(content, &entry)
	if err != nil {
		return nil, err
	}

	return &entry.Response.Result.BillPayment, nil
}
