package generated

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Service struct {
	BusinessId     int64  `json:"business_id"`              // unique id for business
	Id             int64  `json:"id"`                       // unique id of the service
	Name           string `json:"name" validate:"required"` // descriptive name of service
	Billable       bool   `json:"billable"`                 // whether the service is billable to clients or not
	ProjectDefault bool   `json:"project_default"`          // whether the service will be automatically added to new projects in the FreshBooks UI (will not be automatically added in API calls)
	VisState       uint8  `json:"vis_state"`                // 0 for active, 1 for deleted
}

func ListServicesPageSortBy(businessId string, page uint64, sort *SortBy) ([]Service, *Pagination, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/comments/business/" + businessId + "/services")
	if err != nil {
		return nil, nil, err
	}

	content, err := getQueryPage(client, parsedUrl, page, sort)
	if err != nil {
		return nil, nil, err
	}

	type ServiceResult struct {
		Services []Service `json:"services"`
	}

	type ServiceResponseContent struct {
		Result ServiceResult `json:"result"`
		Pagination
	}

	type ServiceResponse struct {
		Response ServiceResponseContent `json:"response"`
	}

	entries := ServiceResponse{
		Response: ServiceResponseContent{
			Result: ServiceResult{
				Services: make([]Service, 0, 10),
			},
		},
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, nil, err
	}

	if len(entries.Response.Result.Services) == 0 {
		return []Service{}, nil, nil
	}

	return entries.Response.Result.Services, &entries.Response.Pagination, nil
}

func ListServicesPage(site string, page uint64) ([]Service, *Pagination, error) {
	return ListServicesPageSortBy(site, page, nil)
}

func RetrieveServiceById(businessId, id string) (*Service, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/comments/business/" + businessId + "/services/" + id)
	if err != nil {
		return nil, err
	}

	content, err := getQueryPage(client, parsedUrl, -1, nil)
	if err != nil {
		return nil, err
	}

	type ServiceResult struct {
		Service Service `json:"service"`
	}

	type ServiceResponseContent struct {
		Result ServiceResult `json:"result"`
	}

	type ServiceResponse struct {
		Response ServiceResponseContent `json:"response"`
	}

	entry := ServiceResponse{}

	err = json.Unmarshal(content, &entry)
	if err != nil {
		return nil, err
	}

	return &entry.Response.Result.Service, nil
}

func RetrieveServiceRateById(businessId, id string) (*ServiceRate, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/comments/business/" + businessId + "/services/" + id + "/rate")
	if err != nil {
		return nil, err
	}

	content, err := getQueryPage(client, parsedUrl, -1, nil)
	if err != nil {
		return nil, err
	}

	type ServiceRateResult struct {
		ServiceRate ServiceRate `json:"service_rate"`
	}

	type ServiceRateResponseContent struct {
		Result ServiceRateResult `json:"result"`
	}

	type ServiceRateResponse struct {
		Response ServiceRateResponseContent `json:"response"`
	}

	entry := ServiceRateResponse{}

	err = json.Unmarshal(content, &entry)
	if err != nil {
		return nil, err
	}

	return &entry.Response.Result.ServiceRate, nil
}
