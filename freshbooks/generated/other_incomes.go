package generated

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"
	"time"

	"github.com/bitsnap/freshbooks-api-client/models"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type OtherIncome struct {
	IncomeId     string                    `json:"incomeid"`      // unique identifier of this other income entry within the business
	Amount       models.OtherIncomeAmount  `json:"amount"`        // subfields: amount and code
	Code         string                    `json:"code"`          // three-letter currency code
	CategoryName string                    `json:"category_name"` // Options include: advertising, in_person_sales, online_sales, rentals, other
	CreatedAt    string                    `json:"created_at"`    // Time the other income entry was created, YYYY-MM-DD HH:MM:SS format
	Date         bool                      `json:"date"`          // The date the income was received, YYYY-MM-DD format
	Note         models.OtherIncomeNote    `json:"note"`          // notes on the income
	PaymentType  string                    `json:"payment_type"`  // Optional type of payment made. “Check”, “Credit”, “Cash”, etc.
	Source       string                    `json:"source"`        // Source of external income. E.g. Shopify, Etsy, Farmers’ Market
	Taxes        []models.OtherIncomeTaxes `json:"taxes"`         // subfields: amount and name
	Name         string                    `json:"name"`          // name of the tax
	UpdatedAt    *time.Time                `json:"updated_at"`    // Time the other income entry was last updated, YYYY-MM-DD HH:MM:SS format
	VisState     uint8                     `json:"vis_state"`     // 0 for active,1 for deleted,2 for archived(more info on vis_state)
}

func ListOtherIncomesPageSortBy(accountId string, page uint64, sort *SortBy) ([]OtherIncome, *Pagination, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/other_incomes/other_incomes")
	if err != nil {
		return nil, nil, err
	}

	content, err := getQueryPage(client, parsedUrl, page, sort)
	if err != nil {
		return nil, nil, err
	}

	type OtherIncomeResult struct {
		OtherIncomes []OtherIncome `json:"other_incomes"`
	}

	type OtherIncomeResponseContent struct {
		Result OtherIncomeResult `json:"result"`
		Pagination
	}

	type OtherIncomeResponse struct {
		Response OtherIncomeResponseContent `json:"response"`
	}

	entries := OtherIncomeResponse{
		Response: OtherIncomeResponseContent{
			Result: OtherIncomeResult{
				OtherIncomes: make([]OtherIncome, 0, 10),
			},
		},
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, nil, err
	}

	if len(entries.Response.Result.OtherIncomes) == 0 {
		return []OtherIncome{}, nil, nil
	}

	return entries.Response.Result.OtherIncomes, &entries.Response.Pagination, nil
}

func ListOtherIncomesPage(site string, page uint64) ([]OtherIncome, *Pagination, error) {
	return ListOtherIncomesPageSortBy(site, page, nil)
}

func RetrieveOtherIncomeById(accountId, id string) (*OtherIncome, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/other_incomes/other_incomes/" + id)
	if err != nil {
		return nil, err
	}

	content, err := getQueryPage(client, parsedUrl, -1, nil)
	if err != nil {
		return nil, err
	}

	type OtherIncomeResult struct {
		OtherIncome OtherIncome `json:"other_income"`
	}

	type OtherIncomeResponseContent struct {
		Result OtherIncomeResult `json:"result"`
	}

	type OtherIncomeResponse struct {
		Response OtherIncomeResponseContent `json:"response"`
	}

	entry := OtherIncomeResponse{}

	err = json.Unmarshal(content, &entry)
	if err != nil {
		return nil, err
	}

	return &entry.Response.Result.OtherIncome, nil
}
