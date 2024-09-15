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

// getter

func (u *UserTeam) GetUserId() *UserId {
	return &u.userId
}

func (u *UserTeam) GetTeamId() *TeamId {
	return &u.teamId
}

func (u *UserTeam) GetState() UserTeamState {
	return u.state
}

// setter

func (u *UserTeam) SetState(state UserTeamState) {
	u.state = state
}
