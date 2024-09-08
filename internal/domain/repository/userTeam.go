package repository

import "CurlARC/internal/domain/entity"

type UserTeamRepository interface {
	Save(userTeam entity.UserTeam) (*entity.UserTeam, error)
	FindUsersByTeamId(teamId string) ([]string, error)   // All users including INVITED users
	FindMembersByTeamId(teamId string) ([]string, error) // Only MEMBERS
	FindTeamsByUserId(userId string) ([]string, error)
	FindInvitedTeamsByUserId(userId string) ([]string, error) // Only INVITED teams
	UpdateState(userTeam entity.UserTeam) (*entity.UserTeam, error)
	Delete(userId, teamId string) error

	IsMember(userId, teamId string) (bool, error)
}
