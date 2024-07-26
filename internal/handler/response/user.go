package response

type SignInResponse struct {
	Jwt   string `json:"jwt"`
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GetUserResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
