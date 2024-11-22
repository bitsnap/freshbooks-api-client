package generated

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"
	"time"

	"github.com/bitsnap/freshbooks-api-client/models"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Item struct {
	Id                 int64                `json:"id"`                  // unique id of item within this business
	ItemId             int64                `json:"itemid"`              // duplicate of id
	AccountingSystemId string               `json:"accounting_systemid"` // unique id of business client exists on
	Description        string               `json:"description"`         // descriptive text for item
	Inventory          string               `json:"inventory"`           // decimal-string count of inventory
	Name               string               `json:"name"`                // descriptive name of item
	Qty                string               `json:"qty"`                 // decimal-string number to multiply unit cost by
	Sku                string               `json:"sku"`                 // id for a specific item or product, used in inventory management. Note the sku is not currently available in the FreshBooks UI.
	Tax1               int64                `json:"tax1"`                // id of default tax for the item
	Tax2               int64                `json:"tax2"`                // id of second default tax for the item
	UnitCost           models.ItemsUnitCost `json:"unit_cost"`           // subfields: amount and code
	Updated            *time.Time           `json:"updated"`             // date item was last updated
	VisState           uint8                `json:"vis_state"`           // 0 for active, 1 for deleted, 2 archived
}

func ListItemsPageSortBy(accountId string, page uint64, sort *SortBy) ([]Item, *Pagination, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/items/items")
	if err != nil {
		return nil, nil, err
	}

	content, err := getQueryPage(client, parsedUrl, page, sort)
	if err != nil {
		return nil, nil, err
	}

	type ItemResult struct {
		Items []Item `json:"items"`
	}

	type ItemResponseContent struct {
		Result ItemResult `json:"result"`
		Pagination
	}

	type ItemResponse struct {
		Response ItemResponseContent `json:"response"`
	}

	entries := ItemResponse{
		Response: ItemResponseContent{
			Result: ItemResult{
				Items: make([]Item, 0, 10),
			},
		},
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, nil, err
	}

	if len(entries.Response.Result.Items) == 0 {
		return []Item{}, nil, nil
	}

	return entries.Response.Result.Items, &entries.Response.Pagination, nil
}

func ListItemsPage(site string, page uint64) ([]Item, *Pagination, error) {
	return ListItemsPageSortBy(site, page, nil)
}

func RetrieveItemById(accountId, id string) (*Item, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/items/items/" + id)
	if err != nil {
		return nil, err
	}

	content, err := getQueryPage(client, parsedUrl, -1, nil)
	if err != nil {
		return nil, err
	}

	type ItemResult struct {
		Item Item `json:"item"`
	}

	type ItemResponseContent struct {
		Result ItemResult `json:"result"`
	}

	type ItemResponse struct {
		Response ItemResponseContent `json:"response"`
	}

	entry := ItemResponse{}

	err = json.Unmarshal(content, &entry)
	if err != nil {
		return nil, err
	}

	return &entry.Response.Result.Item, nil
}
