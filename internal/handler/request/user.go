package request

import "github.com/lib/pq"

type SignUpRequest struct {
	IdToken string `json:"id_token"`
	Name    string `json:"name"`
	Email   string `json:"email"`
}

type SignInRequest struct {
	IdToken string `json:"id_token"`
}

type GetUserRequest struct {
	Id string `json:"id"`
}

type UpdateUserRequest struct {
	Name    string         `json:"name"`
	Email   string         `json:"email"`
	TeamIds pq.StringArray `json:"team_ids"`
}

type DeleteUserRequest struct {
	Id string `json:"id"`
}
