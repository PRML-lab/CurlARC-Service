package usecase

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/domain/repository"
	"errors"
)

type TeamUsecase interface {
	// CRUD
	CreateTeam(name, userId string) error
	GetAllTeams() ([]*model.Team, error)
	GetTeam(id string) (*model.Team, error)
	UpdateTeam(id, name string) error
	DeleteTeam(id string) error

	// User関連
	InviteUser(teamId, userId string) error
	AcceptInvitation(teamId, userId string) error
	RemoveMember(teamId, userId string) error
	GetTeamsByUserId(userId string) ([]*model.Team, error)
	GetMembersByTeamId(teamId string) ([]*model.User, error)
}

type teamUsecase struct {
	teamRepo     repository.TeamRepository
	userRepo     repository.UserRepository
	userTeamRepo repository.UserTeamRepository
}

func NewTeamUsecase(teamRepo repository.TeamRepository, userRepo repository.UserRepository, userTeamRepo repository.UserTeamRepository) TeamUsecase {
	return &teamUsecase{teamRepo: teamRepo, userRepo: userRepo, userTeamRepo: userTeamRepo}
}

func (usecase *teamUsecase) CreateTeam(name, userId string) error {
	team := &model.Team{Name: name}
	createdTeam, err := usecase.teamRepo.Save(team)
	if err != nil {
		return err
	}

	// user-team tableに保存
	_, err = usecase.userRepo.FindById(userId)
	if err != nil {
		return err
	}

	err = usecase.userTeamRepo.Save(userId, createdTeam.Id, "MEMBER")
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

func (usecase *teamUsecase) InviteUser(teamId, userId string) error {
	// Check existence of team and user
	_, err := usecase.teamRepo.FindById(teamId)
	if err != nil {
		return err
	}
	_, err = usecase.userRepo.FindById(userId)
	if err != nil {
		return err
	}
	// Check if the inviter is a member of the team
	isMember, err := usecase.userTeamRepo.IsMember(userId, teamId)
	if err != nil {
		return err
	}
	if !isMember {
		return errors.New("inviter is not a member of the team")
	}

	// Add user to team with "INVITED" state
	err = usecase.userTeamRepo.Save(userId, teamId, "INVITED")
	if err != nil {
		return err
	}

	// Send invitation email
	//////////////////////////

	return nil
}

func (usecase *teamUsecase) AcceptInvitation(teamId, userId string) error {
	// Check existence of team and user
	_, err := usecase.teamRepo.FindById(teamId)
	if err != nil {
		return err
	}
	_, err = usecase.userRepo.FindById(userId)
	if err != nil {
		return err
	}

	// Update state of user-team
	err = usecase.userTeamRepo.UpdateState(userId, teamId)
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

func (usecase *teamUsecase) GetTeamsByUserId(userId string) ([]*model.Team, error) {
	teamIds, err := usecase.userTeamRepo.FindTeamsByUserId(userId)
	if err != nil {
		return nil, err
	}

	var teams []*model.Team
	for _, teamId := range teamIds {
		team, err := usecase.teamRepo.FindById(teamId)
		if err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}

	return teams, nil
}

func (usecase *teamUsecase) GetMembersByTeamId(teamId string) ([]*model.User, error) {
	userIds, err := usecase.userTeamRepo.FindUsersByTeamId(teamId)
	if err != nil {
		return nil, err
	}

	var users []*model.User
	for _, userId := range userIds {
		user, err := usecase.userRepo.FindById(userId)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
