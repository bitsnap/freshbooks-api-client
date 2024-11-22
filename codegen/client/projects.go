package client

import "github.com/bitsnap/bitsnap/packages/data-pipeline-general-accounting/cmd/bitsnap-accounting-freshbooks-codegen/codegen/client"

// GenerateListProjects generates freshbooks client code to fetch all Projects
//
// API: https://api.freshbooks.com/projects/business/<business_id>/projects
//
// Documentation: https://www.freshbooks.com/api/project
func GenerateListProjects() string {
	return client.GenerateListFor("Projects", "Project", "/projects", "projects/business/", "businessId")
}

// GenerateRetrieveProject generates freshbooks client code to retrieve specific Project
//
// API: https://api.freshbooks.com/projects/business/<business_id>/project/<project_id>
//
// Documentation: https://www.freshbooks.com/api/project
func GenerateRetrieveProject() string {
	return client.GenerateRetrieveFor("Project", "/project", "projects/business/", "businessId")
}

// GenerateProjectsModel generates freshbooks Project domain model
//
// Documentation: https://www.freshbooks.com/api/project
func GenerateProjectsModel() string {
	return client.GenerateModel("Project", "https://www.freshbooks.com/api/project")
}
