package repository

import "CurlARC/internal/domain/entity"

type TeamRepository interface {
	Save(team *entity.Team) (*entity.Team, error)
	FindAll() ([]*entity.Team, error)
	FindById(id string) (*entity.Team, error)
	Update(team *entity.Team) (*entity.Team, error)
	Delete(id string) error
}
