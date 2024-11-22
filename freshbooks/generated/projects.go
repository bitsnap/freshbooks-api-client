package generated

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"
	"time"

	"github.com/bitsnap/freshbooks-api-client/models"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Project struct {
	DueDate            *time.Time             `json:"due_date"`                  // date of projected completion.
	FixedPrice         string                 `json:"fixed_price"`               // used for flat-rate projects. Represents the amount being charged to the client for the project
	Group              models.ProjectGroup    `json:"group"`                     // subfields: pending_invitations, id, members
	PendingInvitations []string               `json:"pending_invitations"`       // the pending invites to employees and contractors within the project
	Id                 int64                  `json:"id"`                        // Unique id of group membership
	Members            models.ProjectMembers  `json:"members"`                   // subfields: first_name, last_name, role, identity_id, active, company, id, email
	FirstName          string                 `json:"first_name"`                // first name of the identity
	LastName           string                 `json:"last_name"`                 // last name of the identity
	Role               string                 `json:"role"`                      // named role identity holds in group
	IdentityId         int64                  `json:"identity_id"`               // The unique id for the identity
	Active             bool                   `json:"active"`                    // whether the project is active or not
	Company            string                 `json:"company"`                   // the name of the business
	Email              string                 `json:"email"`                     // email of identity
	Description        string                 `json:"description"`               // description of project
	Complete           bool                   `json:"complete"`                  // whether the project is completed or not
	Title              string                 `json:"title" validate:"required"` // the project’s title
	ProjectType        string                 `json:"project_type"`              // type of project: fixed_price, hourly_rate
	Budget             int64                  `json:"budget"`                    // budget for project
	UpdatedAt          *time.Time             `json:"updated_at"`                // the time the project was last updated
	Services           models.ProjectServices `json:"services"`                  // subfields: business_id, name, id
	BusinessId         int64                  `json:"business_id"`               // unique id for business
	Name               string                 `json:"name"`                      // the name of the service
	Rate               string                 `json:"rate"`                      // the hourly rate of the project. Only applies to hourly_rate projects
	ClientId           int64                  `json:"client_id"`                 // unique id of the client being billed for the project
	CreatedAt          *time.Time             `json:"created_at"`                // the date/time the project was created
	LoggedDuration     int64                  `json:"logged_duration"`           // the time logged for the project in seconds
	BillingMethod      string                 `json:"billing_method"`            // the method of payment for the project
	Internal           bool                   `json:"internal"`                  // clarifies that the project is internally within the company (client is the company)
}

func ListProjectsPageSortBy(businessId string, page uint64, sort *SortBy) ([]Project, *Pagination, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/projects/business/" + businessId + "/projects")
	if err != nil {
		return nil, nil, err
	}

	content, err := getQueryPage(client, parsedUrl, page, sort)
	if err != nil {
		return nil, nil, err
	}

	type ProjectResult struct {
		Projects []Project `json:"projects"`
	}

	type ProjectResponseContent struct {
		Result ProjectResult `json:"result"`
		Pagination
	}

	type ProjectResponse struct {
		Response ProjectResponseContent `json:"response"`
	}

	entries := ProjectResponse{
		Response: ProjectResponseContent{
			Result: ProjectResult{
				Projects: make([]Project, 0, 10),
			},
		},
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, nil, err
	}

	if len(entries.Response.Result.Projects) == 0 {
		return []Project{}, nil, nil
	}

	return entries.Response.Result.Projects, &entries.Response.Pagination, nil
}

func ListProjectsPage(site string, page uint64) ([]Project, *Pagination, error) {
	return ListProjectsPageSortBy(site, page, nil)
}

func RetrieveProjectById(businessId, id string) (*Project, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/projects/business/" + businessId + "/project/" + id)
	if err != nil {
		return nil, err
	}

	content, err := getQueryPage(client, parsedUrl, -1, nil)
	if err != nil {
		return nil, err
	}

	type ProjectResult struct {
		Project Project `json:"project"`
	}

	type ProjectResponseContent struct {
		Result ProjectResult `json:"result"`
	}

	type ProjectResponse struct {
		Response ProjectResponseContent `json:"response"`
	}

	entry := ProjectResponse{}

	err = json.Unmarshal(content, &entry)
	if err != nil {
		return nil, err
	}

	return &entry.Response.Result.Project, nil
}
