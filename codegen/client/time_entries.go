package client

import "github.com/bitsnap/bitsnap/packages/data-pipeline-general-accounting/cmd/bitsnap-accounting-freshbooks-codegen/codegen/client"

// GenerateListTimeEntries generates freshbooks client code to fetch all time entries
//
// API: https://api.freshbooks.com/timetracking/business/<business_id>/time_entries
//
// Documentation: https://www.freshbooks.com/api/time_entries
func GenerateListTimeEntries() string {
	return client.GenerateListFor("TimeEntries", "TimeEntry", "/time_entries", "timetracking/business/", "businessId")
}

// GenerateRetrieveTimeEntry generates freshbooks client code to retrieve specific time entry
//
// API: https://api.freshbooks.com/timetracking/business/<business_id>/time_entries/5095
//
// Documentation: https://www.freshbooks.com/api/time_entries
func GenerateRetrieveTimeEntry() string {
	return client.GenerateRetrieveFor("TimeEntry", "/time_entries", "timetracking/business/", "businessId")
}

// GenerateTimeEntryModel generates freshbooks TimeEntry domain model
//
// Documentation: https://www.freshbooks.com/api/time_entries
func GenerateTimeEntryModel() string {
	return client.GenerateModel("TimeEntry", "https://www.freshbooks.com/api/time_entries")
}
