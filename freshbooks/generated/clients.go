package generated

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Client struct {
	AccountingSystemId     string     `json:"accounting_systemid"`      // unique identifier of business client exists on
	AllowLateFees          bool       `json:"allow_late_fees"`          // deprecated
	AllowLateNotifications bool       `json:"allow_late_notifications"` // deprecated
	BusPhone               string     `json:"bus_phone"`                // business phone number
	CompanyIndustry        string     `json:"company_industry"`         // description of industry client is in
	CompanySize            string     `json:"company_size"`             // size of client’s company
	CurrencyCode           string     `json:"currency_code"`            // 3-letter shortcode for preferred currency
	DirectLinkToken        string     `json:"direct_link_token"`        // deprecated
	Email                  string     `json:"email"`                    // email
	Fax                    string     `json:"fax"`                      // fax number
	FirstName              string     `json:"fname"`                    // first name
	HomePhone              string     `json:"home_phone"`               // home phone number
	Id                     int64      `json:"id"`                       // unique to this business id for client
	Language               string     `json:"language"`                 // shortcode indicating user language e.g. “en”
	LastActivity           string     `json:"last_activity"`            // last client activity action
	LastLogin              *time.Time `json:"last_login"`               // last login time
	Level                  int64      `json:"level"`                    // deprecated: indication of access level on system
	LastName               string     `json:"lname"`                    // last name
	MobPhone               string     `json:"mob_phone"`                // mobile phone number
	Note                   string     `json:"note"`                     // notes kept by admin about client
	Notified               bool       `json:"notified"`                 // deprecated
	NumLogins              int64      `json:"num_logins"`               // number of logins
	Organization           string     `json:"organization"`             // name for client’s business
	BillingCity            string     `json:"p_city"`                   // billing city
	BillingCode            string     `json:"p_code"`                   // billing postal code
	BillingCountry         string     `json:"p_country"`                // billing country
	BillingProvince        string     `json:"p_province"`               // billing province
	BillingStreet          string     `json:"p_street"`                 // billing street address
	BillingStreet2         string     `json:"p_street2"`                // billing street address second part
	PrefersEmail           bool       `json:"pref_email"`               // prefers email over ground mail
	PrefersGroundMail      bool       `json:"pref_gmail"`               // prefers ground mail over email
	ShippingCity           string     `json:"s_city"`                   // shipping address city
	ShippingCode           string     `json:"s_code"`                   // shipping postal code
	ShippingCountry        string     `json:"s_country"`                // shipping country
	ShippingProvince       string     `json:"s_province"`               // short form for province
	ShippingStreet         string     `json:"s_street"`                 // shipping street address
	ShippingStreet2        string     `json:"s_street2"`                // shipping address second street info
	SignupDate             *time.Time `json:"signup_date"`              // time of user signup
	StatementToken         string     `json:"statement_token"`          // deprecated
	Subdomain              string     `json:"subdomain"`                // unused in the new FreshBooks
	Updated                *time.Time `json:"updated"`                  // time of last modification to resource
	UserId                 int64      `json:"userid"`                   // duplicate of id
	Username               string     `json:"username"`                 // deprecated: username used by the client to log into FreshBooks Classic
	VatName                string     `json:"vat_name"`                 // Value Added Tax name
	VatNumber              string     `json:"vat_number"`               // Value Added Tax number
	VisState               int64      `json:"vis_state"`                // “visibility state”, active, deleted, or archived
}

func ListClientsPageSortBy(accountId string, page uint64, sort *SortBy) ([]Client, *Pagination, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/users/clients")
	if err != nil {
		return nil, nil, err
	}

	content, err := getQueryPage(client, parsedUrl, page, sort)
	if err != nil {
		return nil, nil, err
	}

	type ClientResult struct {
		Clients []Client `json:"clients"`
	}

	type ClientResponseContent struct {
		Result ClientResult `json:"result"`
		Pagination
	}

	type ClientResponse struct {
		Response ClientResponseContent `json:"response"`
	}

	entries := ClientResponse{
		Response: ClientResponseContent{
			Result: ClientResult{
				Clients: make([]Client, 0, 10),
			},
		},
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, nil, err
	}

	if len(entries.Response.Result.Clients) == 0 {
		return []Client{}, nil, nil
	}

	return entries.Response.Result.Clients, &entries.Response.Pagination, nil
}

func ListClientsPage(site string, page uint64) ([]Client, *Pagination, error) {
	return ListClientsPageSortBy(site, page, nil)
}

func RetrieveClientById(accountId, id string) (*Client, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/users/clients/" + id)
	if err != nil {
		return nil, err
	}

	content, err := getQueryPage(client, parsedUrl, -1, nil)
	if err != nil {
		return nil, err
	}

	type ClientResult struct {
		Client Client `json:"client"`
	}

	type ClientResponseContent struct {
		Result ClientResult `json:"result"`
	}

	type ClientResponse struct {
		Response ClientResponseContent `json:"response"`
	}

	entry := ClientResponse{}

	err = json.Unmarshal(content, &entry)
	if err != nil {
		return nil, err
	}

	return &entry.Response.Result.Client, nil
}
