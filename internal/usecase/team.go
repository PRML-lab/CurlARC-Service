package usecase

import (
	"CurlARC/internal/domain/entity"
	"CurlARC/internal/domain/repository"
	"errors"
	"fmt"
)

type TeamUsecase interface {
	// CRUD
	CreateTeam(name, userId string) (*entity.Team, error)
	GetAllTeams() ([]*entity.Team, error)
	UpdateTeam(id, name string) (*entity.Team, error)
	DeleteTeam(id string) error

	// User関連
	InviteUsers(teamId, userId string, targetUserEmails []string) error
	AcceptInvitation(teamId, userId string) error
	RemoveMember(teamId, userId string) error
	GetTeamsByUserId(userId string) ([]*entity.Team, error)
	GetInvitedTeams(userId string) ([]*entity.Team, error)
	GetMembersByTeamId(teamId string) ([]*entity.User, error)
}

type teamUsecase struct {
	teamRepo     repository.TeamRepository
	userRepo     repository.UserRepository
	userTeamRepo repository.UserTeamRepository
}

func NewTeamUsecase(teamRepo repository.TeamRepository, userRepo repository.UserRepository, userTeamRepo repository.UserTeamRepository) TeamUsecase {
	return &teamUsecase{teamRepo: teamRepo, userRepo: userRepo, userTeamRepo: userTeamRepo}
}

func (usecase *teamUsecase) CreateTeam(name, userId string) (*entity.Team, error) {
	// Check existence of user
	_, err := usecase.userRepo.FindById(userId)
	if err != nil {
		return nil, err
	}

	// Create entities
	team := entity.NewTeam(name)
	userTeam := entity.NewUserTeam(*entity.NewUserId(userId), *team.GetId(), "MEMBER")

	// 永続化
	savedTeam, err := usecase.teamRepo.Save(team)
	if err != nil {
		return nil, err
	}
	_, err = usecase.userTeamRepo.Save(userTeam)
	if err != nil {
		return nil, err
	}
	return savedTeam, nil
}

func (usecase *teamUsecase) GetAllTeams() ([]*entity.Team, error) {
	teams, err := usecase.teamRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func (usecase *teamUsecase) UpdateTeam(id, name string) (*entity.Team, error) {
	team, err := usecase.teamRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	team.SetName(name)
	updatedTeam, err := usecase.teamRepo.Update(team)
	if err != nil {
		return nil, err
	}

	return updatedTeam, nil
}

func (usecase *teamUsecase) DeleteTeam(id string) error {
	// Check existence of team
	_, err := usecase.teamRepo.FindById(id)
	if err != nil {
		return err
	}

	err = usecase.teamRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *teamUsecase) InviteUsers(teamId, userId string, targetUserEmails []string) error {
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

	var inviteErrors []error

	for _, targetEmail := range targetUserEmails {
		// Check existence of target user
		targetUser, err := usecase.userRepo.FindByEmail(targetEmail)
		if err != nil {
			inviteErrors = append(inviteErrors, fmt.Errorf("target user %s not found: %v", targetEmail, err))
			continue
		}

		// Check if the target user is already a member of the team
		isMember, err = usecase.userTeamRepo.IsMember(targetUser.GetId().Value(), teamId)
		if err != nil {
			inviteErrors = append(inviteErrors, fmt.Errorf("error checking membership for user %s: %v", targetEmail, err))
			continue
		}
		if isMember {
			inviteErrors = append(inviteErrors, fmt.Errorf("target user %s is already a member of the team", targetEmail))
			continue
		}

		// Add user to team with "INVITED" state
		userTeam := entity.NewUserTeam(*entity.NewUserId(targetUser.GetId().Value()), *entity.NewTeamId(teamId), "INVITED")
		_, err = usecase.userTeamRepo.Save(userTeam)
		if err != nil {
			inviteErrors = append(inviteErrors, fmt.Errorf("error inviting user %s: %v", targetEmail, err))
			continue
		}

		// Send invitation email
		// Note: Add email sending logic here if needed
	}

	if len(inviteErrors) > 0 {
		return fmt.Errorf("one or more invitations failed: %v", inviteErrors)
	}

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

	userTeam := entity.NewUserTeam(*entity.NewUserId(userId), *entity.NewTeamId(teamId), "MEMBER")

	// Update state of user-team
	_, err = usecase.userTeamRepo.UpdateState(userTeam)
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

func (usecase *teamUsecase) GetTeamsByUserId(userId string) ([]*entity.Team, error) {
	teamIds, err := usecase.userTeamRepo.FindTeamsByUserId(userId)
	if err != nil {
		return nil, err
	}

	var teams []*entity.Team
	for _, teamId := range teamIds {
		team, err := usecase.teamRepo.FindById(teamId)
		if err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}

	return teams, nil
}

func (usecase *teamUsecase) GetInvitedTeams(userId string) ([]*entity.Team, error) {
	teamIds, err := usecase.userTeamRepo.FindInvitedTeamsByUserId(userId)
	if err != nil {
		return nil, err
	}

	var teams []*entity.Team
	for _, teamId := range teamIds {
		team, err := usecase.teamRepo.FindById(teamId)
		if err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}

	return teams, nil
}

func (usecase *teamUsecase) GetMembersByTeamId(teamId string) ([]*entity.User, error) {
	userIds, err := usecase.userTeamRepo.FindMembersByTeamId(teamId)
	if err != nil {
		return nil, err
	}

	var users []*entity.User
	for _, userId := range userIds {
		user, err := usecase.userRepo.FindById(userId)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
