package infra

import (
	"CurlARC/internal/domain/entity"
	"CurlARC/internal/domain/repository"
	"errors"

	"gorm.io/gorm"
)

type UserTeamRepository struct {
	SqlHandler
}

func NewUserTeamRepository(sqlHandler SqlHandler) repository.UserTeamRepository {
	userTeamRepository := UserTeamRepository{SqlHandler: sqlHandler}
	return &userTeamRepository
}

type UserTeam struct {
	UserId string `gorm:"primaryKey"`
	TeamId string `gorm:"primaryKey"`
	State  string `gorm:"type:varchar(100)"`
}

func (userTeamRepo *UserTeamRepository) Save(userId, teamId string, state entity.UserTeamState) error {
	userTeam := &entity.UserTeam{UserId: userId, TeamId: teamId, State: state}
	result := userTeamRepo.SqlHandler.Conn.Create(userTeam)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (userTeamRepo *UserTeamRepository) FindUsersByTeamId(teamId string) ([]string, error) {

	var userTeams []*entity.UserTeam
	result := userTeamRepo.SqlHandler.Conn.Where("team_id = ?", teamId).Find(&userTeams)

	if result.Error != nil {
		return nil, result.Error
	}

	var userIds []string
	for _, userTeam := range userTeams {
		userIds = append(userIds, userTeam.UserId)
	}

	return userIds, nil
}

func (userTeamRepo *UserTeamRepository) FindMembersByTeamId(teamId string) ([]string, error) {

	var userTeams []*entity.UserTeam
	result := userTeamRepo.SqlHandler.Conn.Where("team_id = ? AND state = ?", teamId, "MEMBER").Find(&userTeams)

	if result.Error != nil {
		return nil, result.Error
	}

	var userIds []string
	for _, userTeam := range userTeams {
		userIds = append(userIds, userTeam.UserId)
	}

	return userIds, nil

}

func (userTeamRepo *UserTeamRepository) FindTeamsByUserId(userId string) ([]string, error) {

	var userTeams []*entity.UserTeam
	result := userTeamRepo.SqlHandler.Conn.Where("user_id = ? AND state = ?", userId, "MEMBER").Find(&userTeams)
	if result.Error != nil {
		return nil, result.Error
	}

	var teamIds []string
	for _, userTeam := range userTeams {
		teamIds = append(teamIds, userTeam.TeamId)
	}

	return teamIds, nil
}

func (userTeamRepo *UserTeamRepository) FindInvitedTeamsByUserId(userId string) ([]string, error) {

	var userTeams []*entity.UserTeam
	result := userTeamRepo.SqlHandler.Conn.Where("user_id = ? AND state = ?", userId, "INVITED").Find(&userTeams)
	if result.Error != nil {
		return nil, result.Error
	}

	var teamIds []string
	for _, userTeam := range userTeams {
		teamIds = append(teamIds, userTeam.TeamId)
	}

	return teamIds, nil
}

func (userTeamRepo *UserTeamRepository) UpdateState(userId, teamId string) error {
	userTeam := &entity.UserTeam{UserId: userId, TeamId: teamId}
	result := userTeamRepo.SqlHandler.Conn.Model(userTeam).Update("state", "MEMBER")
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (userTeamRepo *UserTeamRepository) Delete(userId, teamId string) error {
	userTeam := &entity.UserTeam{UserId: userId, TeamId: teamId}
	result := userTeamRepo.SqlHandler.Conn.Delete(userTeam)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (userTeamRepo *UserTeamRepository) IsMember(userId, teamId string) (bool, error) {
	userTeam := &entity.UserTeam{}
	result := userTeamRepo.SqlHandler.Conn.Where("user_id = ? AND team_id = ? AND state = ?", userId, teamId, "MEMBER").First(userTeam)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}
