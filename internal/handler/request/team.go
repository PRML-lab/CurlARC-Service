package request

type CreateTeamRequest struct {
	Name string `json:"name"`
}

type UpdateTeamRequest struct {
	Name string `json:"name"`
}

type InviteUsersRequest struct {
	TargetUserEmails []string `json:"target_user_emails"`
}