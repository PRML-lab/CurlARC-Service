package entity

type UserTeamState string

const (
	Invited UserTeamState = "INVITED"
	Member  UserTeamState = "MEMBER"
)

type UserTeam struct {
	userId UserId
	teamId TeamId
	state  UserTeamState
}

func NewUserTeam(userId UserId, teamId TeamId, state UserTeamState) *UserTeam {
	return &UserTeam{
		userId: userId,
		teamId: teamId,
		state:  state,
	}
}
