package generated

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"time"
)

func ListBillPayments(id string) ([]BillPayment, error) {
	return ListAll(id, DefaultSortBy(), ListBillPaymentsPageSortBy)
}

func ListBillPaymentsCreatedSince(id string, createdSince *time.Time) ([]BillPayment, error) {
	return ListAll(id, &SortBy{
		Name:      "date",
		Asc:       true,
		AfterDate: createdSince,
	}, ListBillPaymentsPageSortBy)
}

func ListBillPaymentsUpdatedSince(accountId string, updatedSince *time.Time) ([]BillPayment, error) {
	return ListAll(accountId, &SortBy{
		Name:         "updated",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListBillPaymentsPageSortBy)
}

func ListBillVendors(id string) ([]BillVendor, error) {
	return ListAll(id, DefaultSortBy(), ListBillVendorsPageSortBy)
}

func ListBillVendorsCreatedSince(id string, createdSince *time.Time) ([]BillVendor, error) {
	return ListAll(id, &SortBy{
		Name:      "date",
		Asc:       true,
		AfterDate: createdSince,
	}, ListBillVendorsPageSortBy)
}

func ListBillVendorsUpdatedSince(accountId string, updatedSince *time.Time) ([]BillVendor, error) {
	return ListAll(accountId, &SortBy{
		Name:         "updated",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListBillVendorsPageSortBy)
}

func ListBills(id string) ([]Bill, error) {
	return ListAll(id, DefaultSortBy(), ListBillsPageSortBy)
}

func ListBillsCreatedSince(id string, createdSince *time.Time) ([]Bill, error) {
	return ListAll(id, &SortBy{
		Name:      "date",
		Asc:       true,
		AfterDate: createdSince,
	}, ListBillsPageSortBy)
}

func ListBillsUpdatedSince(accountId string, updatedSince *time.Time) ([]Bill, error) {
	return ListAll(accountId, &SortBy{
		Name:         "updated",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListBillsPageSortBy)
}

func ListClients(id string) ([]Client, error) {
	return ListAll(id, DefaultSortBy(), ListClientsPageSortBy)
}

func ListClientsCreatedSince(id string, createdSince *time.Time) ([]Client, error) {
	return ListAll(id, &SortBy{
		Name:      "date",
		Asc:       true,
		AfterDate: createdSince,
	}, ListClientsPageSortBy)
}

func ListClientsUpdatedSince(accountId string, updatedSince *time.Time) ([]Client, error) {
	return ListAll(accountId, &SortBy{
		Name:         "updated",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListClientsPageSortBy)
}

func ListCreditNotes(id string) ([]CreditNote, error) {
	return ListAll(id, DefaultSortBy(), ListCreditNotesPageSortBy)
}

func ListCreditNotesCreatedSince(id string, createdSince *time.Time) ([]CreditNote, error) {
	return ListAll(id, &SortBy{
		Name:      "date",
		Asc:       true,
		AfterDate: createdSince,
	}, ListCreditNotesPageSortBy)
}

func ListCreditNotesUpdatedSince(accountId string, updatedSince *time.Time) ([]CreditNote, error) {
	return ListAll(accountId, &SortBy{
		Name:         "updated",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListCreditNotesPageSortBy)
}

func ListEstimates(id string) ([]Estimate, error) {
	return ListAll(id, DefaultSortBy(), ListEstimatesPageSortBy)
}

func ListEstimatesCreatedSince(id string, createdSince *time.Time) ([]Estimate, error) {
	return ListAll(id, &SortBy{
		Name:      "date",
		Asc:       true,
		AfterDate: createdSince,
	}, ListEstimatesPageSortBy)
}

func ListEstimatesUpdatedSince(accountId string, updatedSince *time.Time) ([]Estimate, error) {
	return ListAll(accountId, &SortBy{
		Name:         "updated",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListEstimatesPageSortBy)
}

func ListExpenses(id string) ([]Expense, error) {
	return ListAll(id, DefaultSortBy(), ListExpensesPageSortBy)
}

func ListExpensesCreatedSince(id string, createdSince *time.Time) ([]Expense, error) {
	return ListAll(id, &SortBy{
		Name:      "date",
		Asc:       true,
		AfterDate: createdSince,
	}, ListExpensesPageSortBy)
}

func ListExpensesUpdatedSince(accountId string, updatedSince *time.Time) ([]Expense, error) {
	return ListAll(accountId, &SortBy{
		Name:         "updated",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListExpensesPageSortBy)
}

func ListInvoiceProfiles(id string) ([]InvoiceProfile, error) {
	return ListAll(id, DefaultSortBy(), ListInvoiceProfilesPageSortBy)
}

func ListInvoiceProfilesCreatedSince(id string, createdSince *time.Time) ([]InvoiceProfile, error) {
	return ListAll(id, &SortBy{
		Name:      "date",
		Asc:       true,
		AfterDate: createdSince,
	}, ListInvoiceProfilesPageSortBy)
}

func ListInvoiceProfilesUpdatedSince(accountId string, updatedSince *time.Time) ([]InvoiceProfile, error) {
	return ListAll(accountId, &SortBy{
		Name:         "updated",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListInvoiceProfilesPageSortBy)
}

func ListItems(id string) ([]Item, error) {
	return ListAll(id, DefaultSortBy(), ListItemsPageSortBy)
}

func ListItemsCreatedSince(id string, createdSince *time.Time) ([]Item, error) {
	return ListAll(id, &SortBy{
		Name:      "date",
		Asc:       true,
		AfterDate: createdSince,
	}, ListItemsPageSortBy)
}

func ListItemsUpdatedSince(accountId string, updatedSince *time.Time) ([]Item, error) {
	return ListAll(accountId, &SortBy{
		Name:         "updated",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListItemsPageSortBy)
}

func ListOtherIncomes(id string) ([]OtherIncome, error) {
	return ListAll(id, DefaultSortBy(), ListOtherIncomesPageSortBy)
}

func ListOtherIncomesCreatedSince(id string, createdSince *time.Time) ([]OtherIncome, error) {
	return ListAll(id, &SortBy{
		Name:      "date",
		Asc:       true,
		AfterDate: createdSince,
	}, ListOtherIncomesPageSortBy)
}

func ListOtherIncomesUpdatedSince(accountId string, updatedSince *time.Time) ([]OtherIncome, error) {
	return ListAll(accountId, &SortBy{
		Name:         "updated",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListOtherIncomesPageSortBy)
}

func ListPaymentOptions(id string) ([]PaymentOption, error) {
	return ListAll(id, DefaultSortBy(), ListPaymentOptionsPageSortBy)
}

func ListPaymentOptionsCreatedSince(id string, createdSince *time.Time) ([]PaymentOption, error) {
	return ListAll(id, &SortBy{
		Name:      "date",
		Asc:       true,
		AfterDate: createdSince,
	}, ListPaymentOptionsPageSortBy)
}

func ListPaymentOptionsUpdatedSince(accountId string, updatedSince *time.Time) ([]PaymentOption, error) {
	return ListAll(accountId, &SortBy{
		Name:         "updated",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListPaymentOptionsPageSortBy)
}

func ListProjects(id string) ([]Project, error) {
	return ListAll(id, DefaultSortBy(), ListProjectsPageSortBy)
}

func ListProjectsCreatedSince(id string, createdSince *time.Time) ([]Project, error) {
	return ListAll(id, &SortBy{
		Name:      "date",
		Asc:       true,
		AfterDate: createdSince,
	}, ListProjectsPageSortBy)
}

func ListProjectsUpdatedSince(accountId string, updatedSince *time.Time) ([]Project, error) {
	return ListAll(accountId, &SortBy{
		Name:         "updated",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListProjectsPageSortBy)
}

func ListServices(id string) ([]Service, error) {
	return ListAll(id, DefaultSortBy(), ListServicesPageSortBy)
}

func ListServicesCreatedSince(id string, createdSince *time.Time) ([]Service, error) {
	return ListAll(id, &SortBy{
		Name:      "date",
		Asc:       true,
		AfterDate: createdSince,
	}, ListServicesPageSortBy)
}

func ListServicesUpdatedSince(accountId string, updatedSince *time.Time) ([]Service, error) {
	return ListAll(accountId, &SortBy{
		Name:         "updated",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListServicesPageSortBy)
}

func ListStaffs(id string) ([]Staff, error) {
	return ListAll(id, DefaultSortBy(), ListStaffsPageSortBy)
}

func ListStaffsCreatedSince(id string, createdSince *time.Time) ([]Staff, error) {
	return ListAll(id, &SortBy{
		Name:      "date",
		Asc:       true,
		AfterDate: createdSince,
	}, ListStaffsPageSortBy)
}

func ListStaffsUpdatedSince(accountId string, updatedSince *time.Time) ([]Staff, error) {
	return ListAll(accountId, &SortBy{
		Name:         "updated",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListStaffsPageSortBy)
}

func ListTasks(id string) ([]Task, error) {
	return ListAll(id, DefaultSortBy(), ListTasksPageSortBy)
}

func ListTasksCreatedSince(id string, createdSince *time.Time) ([]Task, error) {
	return ListAll(id, &SortBy{
		Name:      "date",
		Asc:       true,
		AfterDate: createdSince,
	}, ListTasksPageSortBy)
}

func ListTasksUpdatedSince(accountId string, updatedSince *time.Time) ([]Task, error) {
	return ListAll(accountId, &SortBy{
		Name:         "updated",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListTasksPageSortBy)
}

func ListTaxes(id string) ([]Tax, error) {
	return ListAll(id, DefaultSortBy(), ListTaxesPageSortBy)
}

func ListTaxesCreatedSince(id string, createdSince *time.Time) ([]Tax, error) {
	return ListAll(id, &SortBy{
		Name:      "date",
		Asc:       true,
		AfterDate: createdSince,
	}, ListTaxesPageSortBy)
}

func ListTaxesUpdatedSince(accountId string, updatedSince *time.Time) ([]Tax, error) {
	return ListAll(accountId, &SortBy{
		Name:         "updated",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListTaxesPageSortBy)
}

func ListTeamMembers(id string) ([]TeamMember, error) {
	return ListAll(id, DefaultSortBy(), ListTeamMembersPageSortBy)
}

func ListTeamMembersCreatedSince(id string, createdSince *time.Time) ([]TeamMember, error) {
	return ListAll(id, &SortBy{
		Name:      "date",
		Asc:       true,
		AfterDate: createdSince,
	}, ListTeamMembersPageSortBy)
}

func ListTeamMembersUpdatedSince(accountId string, updatedSince *time.Time) ([]TeamMember, error) {
	return ListAll(accountId, &SortBy{
		Name:         "updated",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListTeamMembersPageSortBy)
}

func ListTimeEntries(id string) ([]TimeEntry, error) {
	return ListAll(id, DefaultSortBy(), ListTimeEntriesPageSortBy)
}

func ListTimeEntriesCreatedSince(id string, createdSince *time.Time) ([]TimeEntry, error) {
	return ListAll(id, &SortBy{
		Name:      "date",
		Asc:       true,
		AfterDate: createdSince,
	}, ListTimeEntriesPageSortBy)
}

func ListTimeEntriesUpdatedSince(accountId string, updatedSince *time.Time) ([]TimeEntry, error) {
	return ListAll(accountId, &SortBy{
		Name:         "updated",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListTimeEntriesPageSortBy)
}
