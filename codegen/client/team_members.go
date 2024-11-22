package client

import "github.com/bitsnap/bitsnap/packages/data-pipeline-general-accounting/cmd/bitsnap-accounting-freshbooks-codegen/codegen/client"

// GenerateListTeamMembers generates freshbooks client code to fetch all team members
//
// API: https://api.freshbooks.com/auth/api/v1/businesses/<business_id>/team_members
//
// Documentation: https://www.freshbooks.com/api/team-members
func GenerateListTeamMembers() string {
	return client.GenerateListFor("TeamMembers", "TeamMember", "/team_members", "auth/api/v1/businesses/", "businessId")
}

// GenerateRetrieveTeamMember generates freshbooks client code to retrieve specific team member
//
// API: https://api.freshbooks.com/auth/api/v1/businesses/<business_id>/team_members/<team_member_uuid>
//
// Documentation: https://www.freshbooks.com/api/team-members
func GenerateRetrieveTeamMember() string {
	return client.GenerateRetrieveFor("TeamMember", "/team_members", "auth/api/v1/businesses/", "businessId")
}

// GenerateTeamMemberModel generates freshbooks team member domain model
//
// Documentation: https://www.freshbooks.com/api/team-members
func GenerateTeamMemberModel() string {
	return client.GenerateModel("TeamMember", "https://www.freshbooks.com/api/team-members")
}
