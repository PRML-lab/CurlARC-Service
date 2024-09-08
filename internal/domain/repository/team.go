package repository

import entity "CurlARC/internal/domain/entity/team"

type TeamRepository interface {
	Save(team *entity.Team) error
	FindAll() ([]*entity.Team, error)
	FindById(id string) (*entity.Team, error)
	Update(team *entity.Team) error
	Delete(id string) error
}
