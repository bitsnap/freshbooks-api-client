package generated

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"
	"time"

	"github.com/bitsnap/freshbooks-api-client/models"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Expense struct {
	AccountId          int64                     `json:"accountid"`                      // id of expense account if applicable
	AccountName        string                    `json:"account_name"`                   // name of related account if applicable
	AccountingSystemId string                    `json:"accounting_systemid"`            // unique id for system
	Amount             models.ExpensesAmount     `json:"amount"`                         // Total amount of the expense
	BackgroundJobId    int64                     `json:"background_jobid"`               // (internal) id for related background job if applicable
	BankName           string                    `json:"bank_name"`                      // name of bank expense was imported from, if applicable
	CategoryId         int64                     `json:"categoryid" validate:"required"` // id of related expense category
	ClientId           int64                     `json:"clientid"`                       // id of related client if applicable
	CompoundedTax      bool                      `json:"compounded_tax"`                 // Deprecated. true/false tax2 was a compound tax
	Date               *time.Time                `json:"date"`                           // date of expense, YYYY-MM-DD format
	ExpenseId          int64                     `json:"expenseid"`                      // unique to this business id for expense
	ExtInvoiceId       int64                     `json:"ext_invoiceid"`                  // id of related contractor invoice if applicable
	ExtSystemId        int64                     `json:"ext_systemid"`                   // id of related contractor system if applicable
	HasReceipt         bool                      `json:"has_receipt"`                    // true/false has receipt attached
	Id                 int64                     `json:"id"`                             // duplicate of expenseid
	InvoiceId          int64                     `json:"invoiceid"`                      // id of related invoice if applicable
	Isduplicate        bool                      `json:"isduplicate"`                    // true/false is duplicated expense
	MarkupPercent      string                    `json:"markup_percent"`                 // string-decimal, note of percent to mark expense up
	Notes              string                    `json:"notes"`                          // notes about expense
	ProfileId          int64                     `json:"profileid"`                      // id of related profile if applicable
	ProjectId          int64                     `json:"projectid"`                      // id of related project if applicable
	StaffId            int64                     `json:"staffid"`                        // id of related staff member if applicable
	Status             int64                     `json:"status"`                         // values from expense status table
	TaxAmount1         models.ExpensesTaxAmount1 `json:"taxAmount1"`                     // The total for first tax applied to the subtotal amount of the expense
	TaxName1           string                    `json:"taxName1"`                       // name of first tax
	TaxPercent1        string                    `json:"taxPercent1"`                    // string-decimal tax amount – indicates the maximum tax percentage for this expense, this does not add tax to the expense, instead use taxAmount1
	TaxAmount2         models.ExpensesTaxAmount2 `json:"taxAmount2"`                     // The total for second tax applied to the subtotal amount of the expense
	TaxName2           string                    `json:"taxName2"`                       // name of second tax
	TaxPercent2        string                    `json:"taxPercent2"`                    // string-decimal tax amount for second tax – indicates the maximum tax percentage for this expense, this does not add tax to the expense, instead use taxAmount2
	TransactionId      int64                     `json:"transactionid"`                  // id of related transaction if applicable
	Updated            *time.Time                `json:"updated"`                        // time invoice last updated at, YYYY-MM-DD HH:MM:SS format
	Vendor             string                    `json:"vendor"`                         // name of vendor
	VisState           uint8                     `json:"vis_state"`                      // 0 for active, 1 for deleted
}

func ListExpensesPageSortBy(accountId string, page uint64, sort *SortBy) ([]Expense, *Pagination, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/expenses/expenses")
	if err != nil {
		return nil, nil, err
	}

	content, err := getQueryPage(client, parsedUrl, page, sort)
	if err != nil {
		return nil, nil, err
	}

	type ExpenseResult struct {
		Expenses []Expense `json:"expenses"`
	}

	type ExpenseResponseContent struct {
		Result ExpenseResult `json:"result"`
		Pagination
	}

	type ExpenseResponse struct {
		Response ExpenseResponseContent `json:"response"`
	}

	entries := ExpenseResponse{
		Response: ExpenseResponseContent{
			Result: ExpenseResult{
				Expenses: make([]Expense, 0, 10),
			},
		},
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, nil, err
	}

	if len(entries.Response.Result.Expenses) == 0 {
		return []Expense{}, nil, nil
	}

	return entries.Response.Result.Expenses, &entries.Response.Pagination, nil
}

func ListExpensesPage(site string, page uint64) ([]Expense, *Pagination, error) {
	return ListExpensesPageSortBy(site, page, nil)
}

func RetrieveExpenseById(accountId, id string) (*Expense, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/expenses/expenses/" + id)
	if err != nil {
		return nil, err
	}

	content, err := getQueryPage(client, parsedUrl, -1, nil)
	if err != nil {
		return nil, err
	}

	type ExpenseResult struct {
		Expense Expense `json:"expense"`
	}

	type ExpenseResponseContent struct {
		Result ExpenseResult `json:"result"`
	}

	type ExpenseResponse struct {
		Response ExpenseResponseContent `json:"response"`
	}

	entry := ExpenseResponse{}

	err = json.Unmarshal(content, &entry)
	if err != nil {
		return nil, err
	}

	return &entry.Response.Result.Expense, nil
}
