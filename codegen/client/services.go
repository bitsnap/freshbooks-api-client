package client

import "github.com/bitsnap/bitsnap/packages/data-pipeline-general-accounting/cmd/bitsnap-accounting-freshbooks-codegen/codegen/client"

// GenerateListServices generates freshbooks client code to fetch all services
//
// API: https://api.freshbooks.com/comments/business/<businessid>/services
//
// Documentation: https://www.freshbooks.com/api/services
func GenerateListServices() string {
	return client.GenerateListFor("Services", "Service", "/services", "comments/business/", "businessId")
}

// GenerateRetrieveService generates freshbooks client code to retrieve specific service
//
// API: https://api.freshbooks.com/comments/business/<businessid>/service/<serviceid>
//
// Documentation: https://www.freshbooks.com/api/services
func GenerateRetrieveService() string {
	return client.GenerateRetrieveFor("Service", "/services", "comments/business/", "businessId")
}

// GenerateRetrieveServiceRate generates freshbooks client code to retrieve specific service rate
//
// API: https://api.freshbooks.com/comments/business/<businessid>/service/<serviceid>/rate
//
// Documentation: https://www.freshbooks.com/api/services
func GenerateRetrieveServiceRate() string {
	return client.GenerateRetrieveForWithEnding("ServiceRate", "/services", "comments/business/", "/rate", "businessId")
}

// GenerateServicesModel generates freshbooks Service domain model
//
// Documentation: https://www.freshbooks.com/api/services
func GenerateServicesModel() string {
	return client.GenerateModel("Service", "https://www.freshbooks.com/api/services")
}
