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

func (team *Team) FromDomain(domain *entity.Team) {
	team.Id = domain.GetId().Value()
	team.Name = domain.GetName()
}

func (team *Team) ToDomain() *entity.Team {
	return entity.NewTeamFromDB(team.Id, team.Name)
}

////////////////////////////////////////
// Team Repository Implementation
////////////////////////////////////////

func (r *TeamRepository) Save(team *entity.Team) (*entity.Team, error) {
	var dbTeam Team
	dbTeam.FromDomain(team)

	if err := r.SqlHandler.Conn.Create(&dbTeam).Error; err != nil {
		return nil, err
	}

	return dbTeam.ToDomain(), nil
}

func (r *TeamRepository) FindAll() ([]*entity.Team, error) {
	var teams []Team
	if err := r.SqlHandler.Conn.Find(&teams).Error; err != nil {
		return nil, err
	}

	var teamsEntity []*entity.Team
	for _, team := range teams {
		teamsEntity = append(teamsEntity, team.ToDomain())
	}

	return teamsEntity, nil
}

func (r *TeamRepository) FindById(id string) (*entity.Team, error) {
	var team Team
	if err := r.SqlHandler.Conn.First(&team, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return team.ToDomain(), nil
}

func (r *TeamRepository) Update(team *entity.Team) (*entity.Team, error) {
	var dbTeam Team
	dbTeam.FromDomain(team)
	if err := r.SqlHandler.Conn.Save(&dbTeam).Error; err != nil {
		return nil, err
	}
	return dbTeam.ToDomain(), nil
}

func (r *TeamRepository) Delete(id string) error {
	if err := r.SqlHandler.Conn.Delete(&Team{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
