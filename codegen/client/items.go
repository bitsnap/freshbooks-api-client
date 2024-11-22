package client

import "github.com/bitsnap/bitsnap/packages/data-pipeline-general-accounting/cmd/bitsnap-accounting-freshbooks-codegen/codegen/client"

// GenerateListItems generates freshbooks client code to fetch all Items
//
// API: https://api.freshbooks.com/accounting/account/<accountid>/items/items
//
// Documentation: https://www.freshbooks.com/api/items
func GenerateListItems() string {
	return client.GenerateList("Items", "Item", "/items/items")
}

// GenerateRetrieveItem generates freshbooks client code to retrieve specific Item
//
// API: https://api.freshbooks.com/accounting/account/<accountid>/items/items/<id>
//
// Documentation: https://www.freshbooks.com/api/items
func GenerateRetrieveItem() string {
	return client.GenerateRetrieve("Item", "/items/items")
}

// GenerateItemsModel generates freshbooks Item domain model
//
// Documentation: https://www.freshbooks.com/api/items
func GenerateItemsModel() string {
	return client.GenerateModel("Item", "https://www.freshbooks.com/api/items")
}
