package client

import "github.com/bitsnap/bitsnap/packages/data-pipeline-general-accounting/cmd/bitsnap-accounting-freshbooks-codegen/codegen/client"

// GenerateListOtherIncomes generates freshbooks client code to fetch all OtherIncomes
//
// API: https://api.freshbooks.com/accounting/account/<accountid>/other_incomes/other_incomes/
//
// Documentation: https://www.freshbooks.com/api/other_income
func GenerateListOtherIncomes() string {
	return client.GenerateList("OtherIncomes", "OtherIncome", "/other_incomes/other_incomes")
}

// GenerateRetrieveOtherIncome generates freshbooks client code to retrieve specific OtherIncome
//
// API: https://api.freshbooks.com/accounting/account/<accountid>/other_incomes/other_incomes/<incomeid>
//
// Documentation: https://www.freshbooks.com/api/other_income
func GenerateRetrieveOtherIncome() string {
	return client.GenerateRetrieve("OtherIncome", "/other_incomes/other_incomes")
}

// GenerateOtherIncomesModel generates freshbooks OtherIncome domain model
//
// Documentation: https://www.freshbooks.com/api/other_income
func GenerateOtherIncomesModel() string {
	return client.GenerateModel("OtherIncome", "https://www.freshbooks.com/api/other_income")
}
