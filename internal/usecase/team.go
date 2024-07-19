package usecase

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/domain/repository"
)

type TeamUsecase interface {
	// CRUD
	CreateTeam(name string) error
	GetAllTeams() ([]*model.Team, error)
	GetTeam(id string) (*model.Team, error)
	UpdateTeam(id, name string) error
	DeleteTeam(id string) error

	// User関連
	AddMember(teamId, userId string) error
	RemoveMember(teamId, userId string) error
}

type teamUsecase struct {
	teamRepo     repository.TeamRepository
	userRepo     repository.UserRepository
	userTeamRepo repository.UserTeamRepository
}

func NewTeamUsecase(teamRepo repository.TeamRepository, userRepo repository.UserRepository, userTeamRepo repository.UserTeamRepository) TeamUsecase {
	return &teamUsecase{teamRepo: teamRepo, userRepo: userRepo, userTeamRepo: userTeamRepo}
}

func (usecase *teamUsecase) CreateTeam(name string) error {
	team := &model.Team{Name: name}
	_, err := usecase.teamRepo.Save(team)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *teamUsecase) GetAllTeams() ([]*model.Team, error) {
	teams, err := usecase.teamRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func (usecase *teamUsecase) GetTeam(id string) (*model.Team, error) {
	team, err := usecase.teamRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (usecase *teamUsecase) UpdateTeam(id, name string) error {
	team, err := usecase.teamRepo.FindById(id)
	if err != nil {
		return err
	}
	team.Name = name
	err = usecase.teamRepo.Update(team)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *teamUsecase) DeleteTeam(id string) error {
	err := usecase.teamRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *teamUsecase) AddMember(teamId, userId string) error {
	// Check existence of team and user
	_, err := usecase.teamRepo.FindById(teamId)
	if err != nil {
		return err
	}
	_, err = usecase.userRepo.FindById(userId)
	if err != nil {
		return err
	}

	// Add user to team
	err = usecase.userTeamRepo.Save(userId, teamId)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *teamUsecase) RemoveMember(teamId, userId string) error {
	// Check existence of team and user
	_, err := usecase.teamRepo.FindById(teamId)
	if err != nil {
		return err
	}
	_, err = usecase.userRepo.FindById(userId)
	if err != nil {
		return err
	}

	// Remove user from team
	err = usecase.userTeamRepo.Delete(userId, teamId)
	if err != nil {
		return err
	}

	return nil
}
