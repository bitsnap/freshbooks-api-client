package templates

import (
	"embed"
	"strings"

	. "github.com/bitsnap/bitsnap/packages/go-util"
)

//go:embed list_convenience.go.tpl
var listConvTemplateFS embed.FS

var listConvTemplate = MustParseTemplate(listConvTemplateFS, "list_convenience.go.tpl")

func GenerateListConvenienceMethods() string {
	type Method struct {
		Name         string
		NameSingular string
	}

	type ListConvenience struct {
		Methods []Method
	}

	output := strings.Builder{}
	err := listConvTemplate.Execute(&output, ListConvenience{
		Methods: []Method{{
			"BillPayments",
			"BillPayment",
		}, {
			"BillVendors",
			"BillVendor",
		}, {
			"Bills",
			"Bill",
		}, {
			"Clients",
			"Client",
		}, {
			"CreditNotes",
			"CreditNote",
		}, {
			"Estimates",
			"Estimate",
		}, {
			"Expenses",
			"Expense",
		}, {
			"InvoiceProfiles",
			"InvoiceProfile",
		}, {
			"Items",
			"Item",
		}, {
			"OtherIncomes",
			"OtherIncome",
		}, {
			"PaymentOptions",
			"PaymentOption",
		}, {
			"Projects",
			"Project",
		}, {
			"Services",
			"Service",
		}, {
			"Staffs",
			"Staff",
		}, {
			"Tasks",
			"Task",
		}, {
			"Taxes",
			"Tax",
		}, {
			"TeamMembers",
			"TeamMember",
		}, {
			"TimeEntries",
			"TimeEntry",
		}},
	})
	if err != nil {
		panic(err)
	}

	return output.String()
}
