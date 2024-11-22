package generated

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type TeamMember struct {
	Uuid                   string     `json:"uuid"`                     // Unique id of the team member
	FirstName              string     `json:"first_name"`               // First name
	MiddleName             string     `json:"middle_name"`              // Middle name
	LastName               string     `json:"last_name"`                // Last name
	Email                  string     `json:"email"`                    // Email
	JobTitle               string     `json:"job_title"`                // Job Title
	Street1                string     `json:"street_1"`                 // Street address
	Street2                string     `json:"street_2"`                 // Second line of the street address
	City                   string     `json:"city"`                     // City
	Province               string     `json:"province"`                 // Province/State
	Country                string     `json:"country"`                  // Country
	PostalCode             string     `json:"postal_code"`              // Postal/ZIP Code
	PhoneNumber            string     `json:"phone_number"`             // Phone Number
	BusinessId             int64      `json:"business_id"`              // The id of this business
	BusinessRoleName       string     `json:"business_role_name"`       // Team member’s role in the business
	Active                 bool       `json:"active"`                   // Whether the member is active or not
	IdentityId             int64      `json:"identity_id"`              // The identity_id of the team member if they have a FreshBooks account
	InvitationDateAccepted *time.Time `json:"invitation_date_accepted"` // The date/time the team member accepted their invitation
	CreatedAt              *time.Time `json:"created_at"`               // The date/time the team member was created
	UpdatedAt              *time.Time `json:"updated_at"`               // The date/time the team member was last modified
}

func ListTeamMembersPageSortBy(businessId string, page uint64, sort *SortBy) ([]TeamMember, *Pagination, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/auth/api/v1/businesses/" + businessId + "/team_members")
	if err != nil {
		return nil, nil, err
	}

	content, err := getQueryPage(client, parsedUrl, page, sort)
	if err != nil {
		return nil, nil, err
	}

	type TeamMemberResult struct {
		TeamMembers []TeamMember `json:"team_members"`
	}

	type TeamMemberResponseContent struct {
		Result TeamMemberResult `json:"result"`
		Pagination
	}

	type TeamMemberResponse struct {
		Response TeamMemberResponseContent `json:"response"`
	}

	entries := TeamMemberResponse{
		Response: TeamMemberResponseContent{
			Result: TeamMemberResult{
				TeamMembers: make([]TeamMember, 0, 10),
			},
		},
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, nil, err
	}

	if len(entries.Response.Result.TeamMembers) == 0 {
		return []TeamMember{}, nil, nil
	}

	return entries.Response.Result.TeamMembers, &entries.Response.Pagination, nil
}

func ListTeamMembersPage(site string, page uint64) ([]TeamMember, *Pagination, error) {
	return ListTeamMembersPageSortBy(site, page, nil)
}

func RetrieveTeamMemberById(businessId, id string) (*TeamMember, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/auth/api/v1/businesses/" + businessId + "/team_members/" + id)
	if err != nil {
		return nil, err
	}

	content, err := getQueryPage(client, parsedUrl, -1, nil)
	if err != nil {
		return nil, err
	}

	type TeamMemberResult struct {
		TeamMember TeamMember `json:"team_member"`
	}

	type TeamMemberResponseContent struct {
		Result TeamMemberResult `json:"result"`
	}

	type TeamMemberResponse struct {
		Response TeamMemberResponseContent `json:"response"`
	}

	entry := TeamMemberResponse{}

	err = json.Unmarshal(content, &entry)
	if err != nil {
		return nil, err
	}

	return &entry.Response.Result.TeamMember, nil
}
