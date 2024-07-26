package response

type Team struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type GetAllTeamsResponse []struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
