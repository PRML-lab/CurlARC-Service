package repository

type UserTeamRepository interface {
	Save(userId, teamId string) error
	Delete(userId, teamId string) error
	FindUsersByTeamId(teamId string) ([]string, error)
	FindTeamsByUserId(userId string) ([]string, error)
}
