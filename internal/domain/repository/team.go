package repository

import "CurlARC/internal/domain/model"

type TeamRepository interface {
	Save(team *model.Team) (*model.Team, error)
	FindAll() ([]*model.Team, error)
	FindById(id string) (*model.Team, error)
	Update(team *model.Team) error
	Delete(id string) error
}
