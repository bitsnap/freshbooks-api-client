package client

import "github.com/bitsnap/bitsnap/packages/data-pipeline-general-accounting/cmd/bitsnap-accounting-freshbooks-codegen/codegen/client"

// GenerateListTasks generates freshbooks client code to fetch all Tasks
//
// API: https://api.freshbooks.com/accounting/account/<account_id>/projects/projects
//
// Documentation: https://www.freshbooks.com/api/task
func GenerateListTasks() string {
	return client.GenerateList("Tasks", "Task", "/projects/projects")
}

// GenerateRetrieveTask generates freshbooks client code to retrieve specific Task
//
// API: https://api.freshbooks.com/accounting/account/<account_id>/projects/tasks/<task_id>
//
// Documentation: https://www.freshbooks.com/api/task
func GenerateRetrieveTask() string {
	return client.GenerateRetrieve("Task", "/projects/projects")
}

// GenerateTasksModel generates freshbooks Task domain model
//
// Documentation: https://www.freshbooks.com/api/task
func GenerateTasksModel() string {
	return client.GenerateModel("Task", "https://www.freshbooks.com/api/tasks")
}
