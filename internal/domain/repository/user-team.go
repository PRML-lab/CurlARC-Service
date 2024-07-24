package repository

type UserTeamRepository interface {
	Save(userId, teamId, state string) error
	FindUsersByTeamId(teamId string) ([]string, error)
	FindTeamsByUserId(userId string) ([]string, error)
	UpdateState(userId, teamId string) error
	Delete(userId, teamId string) error

	IsMember(userId, teamId string) (bool, error)
}
