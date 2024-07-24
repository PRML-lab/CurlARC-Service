package infra

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/domain/repository"
)

type TeamRepository struct {
	SqlHandler
}

func NewTeamRepository(sqlHandler SqlHandler) repository.TeamRepository {
	teamRepository := TeamRepository{SqlHandler: sqlHandler}
	return &teamRepository
}

func (teamRepo *TeamRepository) Save(team *model.Team) (*model.Team, error) {
	result := teamRepo.SqlHandler.Conn.Create(team)
	if result.Error != nil {
		return team, result.Error
	}
	return team, nil
}

func (teamRepo *TeamRepository) FindAll() ([]*model.Team, error) {
	teams := []*model.Team{}
	result := teamRepo.SqlHandler.Conn.Find(&teams)
	if result.Error != nil {
		return nil, result.Error
	}
	return teams, nil
}

func (teamRepo *TeamRepository) FindById(id string) (*model.Team, error) {
	team := new(model.Team)
	result := teamRepo.SqlHandler.Conn.Where("id = ?", id).First(team)
	if result.Error != nil {
		return nil, result.Error
	}
	return team, nil
}

func (teamRepo *TeamRepository) Update(team *model.Team) error {
	result := teamRepo.SqlHandler.Conn.Save(team)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (teamRepo *TeamRepository) Delete(id string) error {
	result := teamRepo.SqlHandler.Conn.Where("id = ?", id).Delete(&model.Team{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
