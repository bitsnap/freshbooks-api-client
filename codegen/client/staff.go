package client

import "github.com/bitsnap/bitsnap/packages/data-pipeline-general-accounting/cmd/bitsnap-accounting-freshbooks-codegen/codegen/client"

// GenerateListStaffs generates freshbooks client code to fetch all staffs
//
// API: https://api.freshbooks.com/accounting/account/<accountid>/users/staffs
//
// Documentation: https://www.freshbooks.com/api/staffs
func GenerateListStaffs() string {
	return client.GenerateList("Staffs", "Staff", "/users/staffs")
}

// GenerateRetrieveStaff generates freshbooks client code to retrieve specific staff
//
// API: https://api.freshbooks.com/accounting/account/<accountid>/users/staffs/<id>
//
// Documentation: https://www.freshbooks.com/api/staff
func GenerateRetrieveStaff() string {
	return client.GenerateRetrieve("Staff", "/users/staffs")
}

// GenerateStaffsModel generates freshbooks Staff domain model
//
// Documentation: https://www.freshbooks.com/api/staff
func GenerateStaffsModel() string {
	return client.GenerateModel("Staff", "https://www.freshbooks.com/api/staff")
}
