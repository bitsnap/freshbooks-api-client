package generated

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"
	"time"

	"github.com/shopspring/decimal"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Staff struct {
	Fax                string           `json:"fax"`                 // fax number
	Rate               *decimal.Decimal `json:"rate"`                // rate this staff is billed at
	NumLogins          int64            `json:"num_logins"`          // how many times staff has logged in
	ApiToken           string           `json:"api_token"`           // token used for deprecated FreshBooks Classic API
	Id                 int64            `json:"id"`                  // unique to this business id for staff
	Note               string           `json:"note"`                // notes about staff
	DisplayName        string           `json:"display_name"`        // name chosen by staff member to display
	LastName           string           `json:"lname"`               // last name
	MobPhone           string           `json:"mob_phone"`           // mobile phone number
	LastLogin          *time.Time       `json:"last_login"`          // date staff account was last logged in to, YYYY-MM-DD HH:MM:SS format
	HomePhone          string           `json:"home_phone"`          // home phone number
	Email              string           `json:"email"`               // email address for staff
	Username           string           `json:"username"`            // username for staff; randomly assigned if none specified at creation time
	Updated            *time.Time       `json:"updated"`             // date staff object was last updated, YYYY-MM-DD HH:MM:SS format
	BillingProvince    string           `json:"p_province"`          // billing address province
	BillingCity        string           `json:"p_city"`              // billing address city
	BillingCode        string           `json:"p_code"`              // billing address zip code
	BillingCountry     string           `json:"p_country"`           // billing address country
	AccountingSystemId string           `json:"accounting_systemid"` // unique identifier of business staff exists on
	BusPhone           string           `json:"bus_phone"`           // business phone number
	SignupDate         *time.Time       `json:"signup_date"`         // date staff account was created, YYYY-MM-DD HH:MM:SS format
	Language           string           `json:"language"`            // staff’s selected language
	Level              int64            `json:"level"`               // deprecated description of access level
	UserId             int64            `json:"userid"`              // duplicate of id
	BillingStreet2     string           `json:"p_street2"`           // billing address secondary street info
	VisState           int64            `json:"vis_state"`           // “visibility state”, active, deleted, or archived
	FirstName          string           `json:"fname"`               // first name
	Organization       string           `json:"organization"`        // organization staff member is affiliated with
	BillingStreet      string           `json:"p_street"`            // billing address primary street info
	CurrencyCode       string           `json:"currency_code"`       // 3-letter shortcode for preferred currency
}

func ListStaffsPageSortBy(accountId string, page uint64, sort *SortBy) ([]Staff, *Pagination, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/users/staffs")
	if err != nil {
		return nil, nil, err
	}

	content, err := getQueryPage(client, parsedUrl, page, sort)
	if err != nil {
		return nil, nil, err
	}

	type StaffResult struct {
		Staffs []Staff `json:"staffs"`
	}

	type StaffResponseContent struct {
		Result StaffResult `json:"result"`
		Pagination
	}

	type StaffResponse struct {
		Response StaffResponseContent `json:"response"`
	}

	entries := StaffResponse{
		Response: StaffResponseContent{
			Result: StaffResult{
				Staffs: make([]Staff, 0, 10),
			},
		},
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, nil, err
	}

	if len(entries.Response.Result.Staffs) == 0 {
		return []Staff{}, nil, nil
	}

	return entries.Response.Result.Staffs, &entries.Response.Pagination, nil
}

func ListStaffsPage(site string, page uint64) ([]Staff, *Pagination, error) {
	return ListStaffsPageSortBy(site, page, nil)
}

func RetrieveStaffById(accountId, id string) (*Staff, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/users/staffs/" + id)
	if err != nil {
		return nil, err
	}

	content, err := getQueryPage(client, parsedUrl, -1, nil)
	if err != nil {
		return nil, err
	}

	type StaffResult struct {
		Staff Staff `json:"staff"`
	}

	type StaffResponseContent struct {
		Result StaffResult `json:"result"`
	}

	type StaffResponse struct {
		Response StaffResponseContent `json:"response"`
	}

	entry := StaffResponse{}

	err = json.Unmarshal(content, &entry)
	if err != nil {
		return nil, err
	}

	return &entry.Response.Result.Staff, nil
}
