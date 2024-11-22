package main

import (
	"github.com/bitsnap/bitsnap/packages/data-pipeline-general-accounting/cmd/bitsnap-accounting-freshbooks-codegen/codegen"
	. "github.com/bitsnap/bitsnap/packages/go-util"
	"strings"

	"github.com/bitsnap/bitsnap/packages/data-pipeline-general-accounting/cmd/bitsnap-accounting-freshbooks-codegen/codegen/client"

	"github.com/bitsnap/bitsnap/packages/data-pipeline-general-accounting/cmd/bitsnap-accounting-freshbooks-codegen/freshbooks"
	goutil "github.com/bitsnap/bitsnap/packages/go-util"
)

var content = map[string][]func() string{
	"bill_payments.go": {
		freshbooks.GenerateBillPaymentsModel,
		freshbooks.GenerateListBillPayments,
		freshbooks.GenerateRetrieveBillPayment,
	},
	"bills.go": {
		freshbooks.GenerateBillsModel,
		freshbooks.GenerateListBills,
		freshbooks.GenerateRetrieveBill,
	},
	"bill_vendors.go": {
		freshbooks.GenerateBillVendorsModel,
		freshbooks.GenerateListBillVendors,
	},
	//"business_memberships.go": {
	//},
	"clients.go": {
		freshbooks.GenerateClientsModel,
		freshbooks.GenerateListClients,
		freshbooks.GenerateRetrieveClient,
	},
	"credit_notes.go": {
		freshbooks.GenerateCreditNotesModel,
		freshbooks.GenerateListCreditNotes,
		freshbooks.GenerateRetrieveCreditNote,
	},
	"estimates.go": {
		freshbooks.GenerateEstimateModel,
		freshbooks.GenerateListEstimates,
		freshbooks.GenerateRetrieveEstimate,
	},
	"expenses.go": {
		freshbooks.GenerateExpenseModel,
		freshbooks.GenerateListExpenses,
		freshbooks.GenerateRetrieveExpense,
	},
	//"gateways.go":         {
	//
	//},
	"invoice_profiles.go": {
		freshbooks.GenerateInvoiceProfilesModel,
		freshbooks.GenerateListInvoiceProfiles,
		freshbooks.GenerateRetrieveInvoiceProfile,
	},
	"invoices.go": {
		freshbooks.GenerateInvoicesModel,
		freshbooks.GenerateListInvoices,
		freshbooks.GenerateRetrieveInvoice,
	},
	"items.go": {
		freshbooks.GenerateItemsModel,
		freshbooks.GenerateListItems,
		freshbooks.GenerateRetrieveItem,
	},
	"other_incomes.go": {
		freshbooks.GenerateOtherIncomesModel,
		freshbooks.GenerateListOtherIncomes,
		freshbooks.GenerateRetrieveOtherIncome,
	},
	"payment_options.go": {
		freshbooks.GenerateRetrievePaymentOptions,
	},
	"payments.go": {
		freshbooks.GeneratePaymentsModel,
		freshbooks.GenerateListPayments,
		freshbooks.GenerateRetrievePayment,
	},
	"projects.go": {
		freshbooks.GenerateProjectsModel,
		freshbooks.GenerateListProjects,
		freshbooks.GenerateRetrieveProject,
	},
	"services.go": {
		freshbooks.GenerateServicesModel,
		freshbooks.GenerateListServices,
		freshbooks.GenerateRetrieveService,
		freshbooks.GenerateRetrieveServiceRate,
	},
	"staff.go": {
		freshbooks.GenerateStaffsModel,
		freshbooks.GenerateListStaffs,
		freshbooks.GenerateRetrieveStaff,
	},
	"tasks.go": {
		freshbooks.GenerateTasksModel,
		freshbooks.GenerateListTasks,
		freshbooks.GenerateRetrieveTask,
	},
	"taxes.go": {
		freshbooks.GenerateTaxModel,
		freshbooks.GenerateListTaxes,
		freshbooks.GenerateRetrieveTax,
	},
	"team_members.go": {
		freshbooks.GenerateTeamMemberModel,
		freshbooks.GenerateListTeamMembers,
		freshbooks.GenerateRetrieveTeamMember,
	},
	"time_entries.go": {
		freshbooks.GenerateTimeEntryModel,
		freshbooks.GenerateListTimeEntries,
		freshbooks.GenerateRetrieveTimeEntry,
	},
	"list_convenience.go": {
		client.GenerateListConvenienceMethods,
	},
}

func header(content string) string {
	enums := "\t\"github.com/bitsnap/freshbooks-api-client/enums\"\n"
	models := "\t\"github.com/bitsnap/freshbooks-api-client/models\"\n"
	decimal := "\t\"github.com/shopspring/decimal\"\n"
	time := "\t\"time\"\n"

	if !strings.Contains(content, "models.") {
		models = ""
	}

	if !strings.Contains(content, "enums.") {
		enums = ""
	}

	if !strings.Contains(content, "time.") {
		time = ""
	}

	if !strings.Contains(content, "decimal.") {
		decimal = ""
	}

	return `package freshbooks

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
` + enums + models + decimal + time + `
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

`
}

func GenerateInto(targetDir string) {
	goutil.GenerateInto(targetDir, header, content)
}

func main() {
	dir := TargetCodegenDir()
	syncLogger := SetGlobalLogger()
	defer syncLogger()

	Logger().Infow("starting bitsnap freshbooks codegen",
		"dir", dir,
	)

	codegen.GenerateInto(dir)
}
