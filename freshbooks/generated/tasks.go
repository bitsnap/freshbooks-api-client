package generated

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"
	"time"

	"github.com/bitsnap/freshbooks-api-client/models"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Task struct {
	Id          int64            `json:"id"`          // the unique id for the task
	TaskId      int64            `json:"taskid"`      // duplicate of id
	Name        string           `json:"name"`        // the name of the task
	Description string           `json:"description"` // the description of the task
	Rate        models.TasksRate `json:"rate"`        // subfields: amount, code
	Billable    bool             `json:"billable"`    // true if the task is billable
	Tax1        int64            `json:"tax1"`        // id of tax on task
	Tax2        int64            `json:"tax2"`        // id of tax on task
	Updated     *time.Time       `json:"updated"`     // the Date/Time the task was last updated
	VisState    int64            `json:"vis_state"`   // 0 marks the task as active, 1 if inactive
}

func ListTasksPageSortBy(accountId string, page uint64, sort *SortBy) ([]Task, *Pagination, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/projects/projects")
	if err != nil {
		return nil, nil, err
	}

	content, err := getQueryPage(client, parsedUrl, page, sort)
	if err != nil {
		return nil, nil, err
	}

	type TaskResult struct {
		Tasks []Task `json:"tasks"`
	}

	type TaskResponseContent struct {
		Result TaskResult `json:"result"`
		Pagination
	}

	type TaskResponse struct {
		Response TaskResponseContent `json:"response"`
	}

	entries := TaskResponse{
		Response: TaskResponseContent{
			Result: TaskResult{
				Tasks: make([]Task, 0, 10),
			},
		},
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, nil, err
	}

	if len(entries.Response.Result.Tasks) == 0 {
		return []Task{}, nil, nil
	}

	return entries.Response.Result.Tasks, &entries.Response.Pagination, nil
}

func ListTasksPage(site string, page uint64) ([]Task, *Pagination, error) {
	return ListTasksPageSortBy(site, page, nil)
}

func RetrieveTaskById(accountId, id string) (*Task, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/accounting/account/" + accountId + "/projects/projects/" + id)
	if err != nil {
		return nil, err
	}

	content, err := getQueryPage(client, parsedUrl, -1, nil)
	if err != nil {
		return nil, err
	}

	type TaskResult struct {
		Task Task `json:"task"`
	}

	type TaskResponseContent struct {
		Result TaskResult `json:"result"`
	}

	type TaskResponse struct {
		Response TaskResponseContent `json:"response"`
	}

	entry := TaskResponse{}

	err = json.Unmarshal(content, &entry)
	if err != nil {
		return nil, err
	}

	return &entry.Response.Result.Task, nil
}
