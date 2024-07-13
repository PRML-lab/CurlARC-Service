package response

import "github.com/lib/pq"

type SignInResponse struct {
	Jwt     string         `json:"jwt"`
	Id      string         `json:"id"`
	Name    string         `json:"name"`
	Email   string         `json:"email"`
	TeamIds pq.StringArray `json:"team_ids"`
}

type GetUserResponse struct {
	Id      string         `json:"id"`
	Name    string         `json:"name"`
	Email   string         `json:"email"`
	TeamIds pq.StringArray `json:"team_ids"`
}
