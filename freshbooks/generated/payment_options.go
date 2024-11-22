package generated

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func ListPaymentOptionsPageSortBy(accountId string, page uint64, sort *SortBy) ([]PaymentOption, *Pagination, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/payment_options?entity_type=invoice")
	if err != nil {
		return nil, nil, err
	}

	content, err := getQueryPage(client, parsedUrl, page, sort)
	if err != nil {
		return nil, nil, err
	}

	type PaymentOptionResult struct {
		PaymentOptions []PaymentOption `json:"payment_options"`
	}

	type PaymentOptionResponseContent struct {
		Result PaymentOptionResult `json:"result"`
		Pagination
	}

	type PaymentOptionResponse struct {
		Response PaymentOptionResponseContent `json:"response"`
	}

	entries := PaymentOptionResponse{
		Response: PaymentOptionResponseContent{
			Result: PaymentOptionResult{
				PaymentOptions: make([]PaymentOption, 0, 10),
			},
		},
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, nil, err
	}

	if len(entries.Response.Result.PaymentOptions) == 0 {
		return []PaymentOption{}, nil, nil
	}

	return entries.Response.Result.PaymentOptions, &entries.Response.Pagination, nil
}

func ListPaymentOptionsPage(site string, page uint64) ([]PaymentOption, *Pagination, error) {
	return ListPaymentOptionsPageSortBy(site, page, nil)
}
