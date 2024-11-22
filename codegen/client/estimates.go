package client

import "github.com/bitsnap/bitsnap/packages/data-pipeline-general-accounting/cmd/bitsnap-accounting-freshbooks-codegen/codegen/client"

// GenerateListEstimates generates freshbooks client code to fetch all Estimates
//
// API: https://api.freshbooks.com/accounting/account/<accountid>/estimates/estimates
//
// Documentation: https://www.freshbooks.com/api/estimates
func GenerateListEstimates() string {
	return client.GenerateList("Estimates", "Estimate", "/estimates/estimates")
}

// GenerateRetrieveEstimate generates freshbooks client code to retrieve specific Estimate
//
// API: https://api.freshbooks.com/accounting/account/<accountid>/estimates/estimates/<estimateid>
//
// Documentation: https://www.freshbooks.com/api/estimates
func GenerateRetrieveEstimate() string {
	return client.GenerateRetrieve("Estimate", "/estimates/estimates")
}

// GenerateEstimateModel generates freshbooks Estimate domain model
//
// Documentation: https://www.freshbooks.com/api/estimates
func GenerateEstimateModel() string {
	return client.GenerateModel("Estimate", "https://www.freshbooks.com/api/estimates")
}
