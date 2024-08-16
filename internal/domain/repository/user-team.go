package repository

type UserTeamRepository interface {
	Save(userId, teamId, state string) error
	FindUsersByTeamId(teamId string) ([]string, error) // All users including INVITED users
	FindMembersByTeamId(teamId string) ([]string, error) // Only MEMBERS
	FindTeamsByUserId(userId string) ([]string, error)
	FindInvitedTeamsByUserId(userId string) ([]string, error) // Only INVITED teams
	UpdateState(userId, teamId string) error
	Delete(userId, teamId string) error

	IsMember(userId, teamId string) (bool, error)
}
