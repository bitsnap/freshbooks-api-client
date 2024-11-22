package generated

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type TimeEntry struct {
	Billable   bool   `json:"billable"`                       // True for entries that can be automatically added to an invoice
	Billed     bool   `json:"billed"`                         // True for entries that have already been added to an invoice or manually marked as billed
	ClientId   int64  `json:"client_id"`                      // The unique identifier of the client to be billed for this entry
	Duration   int64  `json:"duration" validate:"required"`   // The length of time in seconds
	Internal   bool   `json:"internal"`                       // True if the time entry is not assigned to a client
	IsLogged   bool   `json:"is_logged" validate:"required"`  // False if the time entry is being created from a running timer
	Note       string `json:"note"`                           // A short description of the work being done during the time
	ProjectId  int64  `json:"project_id"`                     // The unique identifier of the project worked on during the time
	ServiceId  int64  `json:"service_id"`                     // The unique identifier of the project service worked on during the time
	StartedAt  int64  `json:"started_at" validate:"required"` // The date/time in UTC when the work started
	IdentityId int64  `json:"identity_id"`                    // The unique identifier of the teammate or user logging the time entry
}

func ListTimeEntriesPageSortBy(businessId string, page uint64, sort *SortBy) ([]TimeEntry, *Pagination, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/timetracking/business/" + businessId + "/time_entries")
	if err != nil {
		return nil, nil, err
	}

	content, err := getQueryPage(client, parsedUrl, page, sort)
	if err != nil {
		return nil, nil, err
	}

	type TimeEntryResult struct {
		TimeEntries []TimeEntry `json:"time_entries"`
	}

	type TimeEntryResponseContent struct {
		Result TimeEntryResult `json:"result"`
		Pagination
	}

	type TimeEntryResponse struct {
		Response TimeEntryResponseContent `json:"response"`
	}

	entries := TimeEntryResponse{
		Response: TimeEntryResponseContent{
			Result: TimeEntryResult{
				TimeEntries: make([]TimeEntry, 0, 10),
			},
		},
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, nil, err
	}

	if len(entries.Response.Result.TimeEntries) == 0 {
		return []TimeEntry{}, nil, nil
	}

	return entries.Response.Result.TimeEntries, &entries.Response.Pagination, nil
}

func ListTimeEntriesPage(site string, page uint64) ([]TimeEntry, *Pagination, error) {
	return ListTimeEntriesPageSortBy(site, page, nil)
}

func RetrieveTimeEntryById(businessId, id string) (*TimeEntry, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/timetracking/business/" + businessId + "/time_entries/" + id)
	if err != nil {
		return nil, err
	}

	content, err := getQueryPage(client, parsedUrl, -1, nil)
	if err != nil {
		return nil, err
	}

	type TimeEntryResult struct {
		TimeEntry TimeEntry `json:"time_entry"`
	}

	type TimeEntryResponseContent struct {
		Result TimeEntryResult `json:"result"`
	}

	type TimeEntryResponse struct {
		Response TimeEntryResponseContent `json:"response"`
	}

	entry := TimeEntryResponse{}

	err = json.Unmarshal(content, &entry)
	if err != nil {
		return nil, err
	}

	return &entry.Response.Result.TimeEntry, nil
}
