package response

type AuthorizeResponse struct {
	User        User   `json:"user"`
	AccessToken string `json:"access_token"`
}

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
