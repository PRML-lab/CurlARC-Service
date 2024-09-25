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

func (userTeam *UserTeam) FromDomain(userTeamEntity *entity.UserTeam) {
	userTeam.UserId = userTeamEntity.GetUserId().Value()
	userTeam.TeamId = userTeamEntity.GetTeamId().Value()
	userTeam.State = string(userTeamEntity.GetState())
}

func (userTeam *UserTeam) ToDomain() *entity.UserTeam {
	return entity.NewUserTeam(
		*entity.NewUserId(userTeam.UserId),
		*entity.NewTeamId(userTeam.TeamId),
		entity.UserTeamState(userTeam.State),
	)
}

func (userTeamRepo *UserTeamRepository) Save(userTeam *entity.UserTeam) (*entity.UserTeam, error) {
	var dbUserTeam UserTeam
	dbUserTeam.FromDomain(userTeam)

	if err := userTeamRepo.SqlHandler.Conn.Create(&dbUserTeam).Error; err != nil {
		return nil, err
	}

	return dbUserTeam.ToDomain(), nil
}

func (userTeamRepo *UserTeamRepository) FindUsersByTeamId(teamId string) ([]string, error) {
	var userTeams []*UserTeam
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
	var userTeams []*UserTeam
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
	var userTeams []*UserTeam
	result := userTeamRepo.SqlHandler.Conn.Where("user_id = ?", userId).Find(&userTeams)

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
	var userTeams []*UserTeam
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

func (userTeamRepo *UserTeamRepository) UpdateState(userTeam *entity.UserTeam) (*entity.UserTeam, error) {
	var dbUserTeam UserTeam
	dbUserTeam.FromDomain(userTeam)

	result := userTeamRepo.SqlHandler.Conn.Model(&dbUserTeam).
		Where("user_id = ? AND team_id = ?", dbUserTeam.UserId, dbUserTeam.TeamId).
		Update("state", dbUserTeam.State)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("user team not found")
	}

	return dbUserTeam.ToDomain(), nil
}

func (userTeamRepo *UserTeamRepository) Delete(userId, teamId string) error {
	result := userTeamRepo.SqlHandler.Conn.Delete(&UserTeam{}, "user_id = ? AND team_id = ?", userId, teamId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (userTeamRepo *UserTeamRepository) IsMember(userId, teamId string) (bool, error) {
	var userTeam UserTeam
	result := userTeamRepo.SqlHandler.Conn.First(&userTeam, "user_id = ? AND team_id = ?", userId, teamId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if result.Error != nil {
		return false, result.Error
	}

	return userTeam.State == "MEMBER", nil
}
