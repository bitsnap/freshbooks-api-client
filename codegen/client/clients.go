package client

import "github.com/bitsnap/bitsnap/packages/data-pipeline-general-accounting/cmd/bitsnap-accounting-freshbooks-codegen/codegen/client"

// GenerateListClients generates freshbooks client code to fetch all clients
//
// API: https://api.freshbooks.com/accounting/account/<accountId>/users/clients
//
// Documentation: https://www.freshbooks.com/api/clients
func GenerateListClients() string {
	return client.GenerateList("Clients", "Client", "/users/clients")
}

// GenerateRetrieveClient generates freshbooks client code to retrieve specific client
//
// API:https://api.freshbooks.com/accounting/account/<accountId>/users/clients/<client-id>
//
// Documentation: https://www.freshbooks.com/api/clients
func GenerateRetrieveClient() string {
	return client.GenerateRetrieve("Client", "/users/clients")
}

// GenerateClientsModel generates freshbooks Client domain model
//
// Documentation: https://www.freshbooks.com/api/clients
func GenerateClientsModel() string {
	return client.GenerateModel("Client", "https://www.freshbooks.com/api/clients")
}
