package response

type GetUserResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SignInResponse struct {
	Id 	string `json:"id"`
	Name 	string `json:"name"`
	Email 	string `json:"email"`
}