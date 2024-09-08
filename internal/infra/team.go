package infra

import (
	"CurlARC/internal/domain/entity"
	"CurlARC/internal/domain/repository"
)

type TeamRepository struct {
	SqlHandler
}

func NewTeamRepository(sqlHandler SqlHandler) repository.TeamRepository {
	teamRepository := TeamRepository{SqlHandler: sqlHandler}
	return &teamRepository
}

func (teamRepo *TeamRepository) Save(team *entity.Team) (*entity.Team, error) {
	result := teamRepo.SqlHandler.Conn.Create(team)
	if result.Error != nil {
		return team, result.Error
	}
	return team, nil
}

func (teamRepo *TeamRepository) FindAll() ([]*entity.Team, error) {
	teams := []*entity.Team{}
	result := teamRepo.SqlHandler.Conn.Find(&teams)
	if result.Error != nil {
		return nil, result.Error
	}
	return teams, nil
}

func (teamRepo *TeamRepository) FindById(id string) (*entity.Team, error) {
	team := new(entity.Team)
	result := teamRepo.SqlHandler.Conn.Where("id = ?", id).First(team)
	if result.Error != nil {
		return nil, result.Error
	}
	return team, nil
}

func (teamRepo *TeamRepository) Update(team *entity.Team) error {
	result := teamRepo.SqlHandler.Conn.Save(team)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (teamRepo *TeamRepository) Delete(id string) error {
	result := teamRepo.SqlHandler.Conn.Where("id = ?", id).Delete(&entity.Team{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
