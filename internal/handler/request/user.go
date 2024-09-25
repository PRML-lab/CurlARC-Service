package request

// type SignUpRequest struct {
// 	IdToken string `json:"id_token"`
// 	Name    string `json:"name"`
// 	Email   string `json:"email"`
// }

type AuthorizeRequest struct {
	IdToken string `json:"id_token"`
}

type GetUserRequest struct {
	Id string `json:"id"`
}

type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type DeleteUserRequest struct {
	Id string `json:"id"`
}
