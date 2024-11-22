package client

import "github.com/bitsnap/bitsnap/packages/data-pipeline-general-accounting/cmd/bitsnap-accounting-freshbooks-codegen/codegen/client"

// GenerateListExpenses generates freshbooks client code to fetch all Expenses
//
// API: https://api.freshbooks.com/accounting/account/<accountid>/expenses/expenses
//
// Documentation: https://www.freshbooks.com/api/expenses
func GenerateListExpenses() string {
	return client.GenerateList("Expenses", "Expense", "/expenses/expenses")
}

// GenerateRetrieveExpense generates freshbooks client code to retrieve specific Expense
//
// API: https://api.freshbooks.com/accounting/account/<accountid>/expenses/expenses/<expenseid>
//
// Documentation: https://www.freshbooks.com/api/expenses
func GenerateRetrieveExpense() string {
	return client.GenerateRetrieve("Expense", "/expenses/expenses")
}

// GenerateExpenseModel generates freshbooks Expense domain model
//
// Documentation: https://www.freshbooks.com/api/expenses
func GenerateExpenseModel() string {
	return client.GenerateModel("Expense", "https://www.freshbooks.com/api/expenses")
}
