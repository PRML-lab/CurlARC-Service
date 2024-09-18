package usecase_test

import (
	"errors"
	"testing"

	"CurlARC/internal/domain/entity"
	"CurlARC/internal/usecase"
	"CurlARC/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTeam(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTeamRepo := mock.NewMockTeamRepository(ctrl)
	mockUserRepo := mock.NewMockUserRepository(ctrl)
	mockUserTeamRepo := mock.NewMockUserTeamRepository(ctrl)

	teamUsecase := usecase.NewTeamUsecase(
		mockTeamRepo,
		mockUserRepo,
		mockUserTeamRepo,
	)

	t.Run("正常系: チームが正常に作成される", func(t *testing.T) {
		team := entity.NewTeam("Team A")
		userId := "user-123"
		user := entity.NewUser(*entity.NewUserId(userId), "User A", "A@gmail.com")
		userTeam := entity.NewUserTeam(*entity.NewUserId(userId), *team.GetId(), entity.Member)

		mockTeamRepo.EXPECT().Save(team).Return(team, nil)
		mockUserRepo.EXPECT().FindById(userId).Return(user, nil)
		mockUserTeamRepo.EXPECT().Save(userTeam).Return(nil)

		createdTeam, err := teamUsecase.CreateTeam("Team A", userId)
		assert.NoError(t, err)
		assert.Equal(t, team, createdTeam)
	})

	t.Run("異常系: dbへの保存に失敗する", func(t *testing.T) {
		team := entity.NewTeam("Team B")
		userId := "user-123"

		mockTeamRepo.EXPECT().Save(team).Return(nil, errors.New("failed to save team"))

		createdTeam, err := teamUsecase.CreateTeam("Team B", userId)
		assert.Error(t, err)
		assert.Nil(t, createdTeam)
		assert.Equal(t, "failed to save team", err.Error())
	})

	t.Run("異常系: 作成者のユーザーが見つからない", func(t *testing.T) {
		team := entity.NewTeam("Team C")
		userId := "user-123"

		mockTeamRepo.EXPECT().Save(team).Return(team, nil)
		mockUserRepo.EXPECT().FindById(userId).Return(nil, errors.New("user not found"))

		_, err := teamUsecase.CreateTeam("Team C", userId)
		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
	})
}

func TestGetAllTeams(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTeamRepo := mock.NewMockTeamRepository(ctrl)
	mockUserRepo := mock.NewMockUserRepository(ctrl)
	mockUserTeamRepo := mock.NewMockUserTeamRepository(ctrl)

	teamUsecase := usecase.NewTeamUsecase(mockTeamRepo, mockUserRepo, mockUserTeamRepo)

	t.Run("正常系: チームが正常に取得される", func(t *testing.T) {
		teams := []*entity.Team{
			entity.NewTeam("Team A"),
			entity.NewTeam("Team B"),
		}

		mockTeamRepo.EXPECT().FindAll().Return(teams, nil)

		result, err := teamUsecase.GetAllTeams()
		assert.NoError(t, err)
		assert.Equal(t, teams, result)
	})

	t.Run("異常系: チームの取得に失敗する", func(t *testing.T) {
		mockTeamRepo.EXPECT().FindAll().Return(nil, errors.New("failed to get teams"))

		result, err := teamUsecase.GetAllTeams()
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "failed to get teams", err.Error())
	})
}

func TestUpdateTeam(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTeamRepo := mock.NewMockTeamRepository(ctrl)
	mockUserRepo := mock.NewMockUserRepository(ctrl)
	mockUserTeamRepo := mock.NewMockUserTeamRepository(ctrl)

	teamUsecase := usecase.NewTeamUsecase(mockTeamRepo, mockUserRepo, mockUserTeamRepo)

	team := entity.NewTeam("Team A")
	ToUpdateTeam := entity.NewTeam("Team A+")

	t.Run("正常系: チームが正常に更新される", func(t *testing.T) {

		mockTeamRepo.EXPECT().FindById("team-123").Return(team, nil)
		mockTeamRepo.EXPECT().Update(team).Return(nil)

		updatedTeam, err := teamUsecase.UpdateTeam("team-123", "Team A+")
		assert.NoError(t, err)
		assert.Equal(t, updatedTeam, ToUpdateTeam)
	})

	t.Run("異常系: チームが見つからない", func(t *testing.T) {
		mockTeamRepo.EXPECT().FindById("team-123").Return(nil, errors.New("team not found"))

		_, err := teamUsecase.UpdateTeam("team-123", "Team A")
		assert.Error(t, err)
		assert.Equal(t, "team not found", err.Error())
	})

	t.Run("異常系: チームの更新に失敗する", func(t *testing.T) {
		mockTeamRepo.EXPECT().FindById("team-123").Return(team, nil)
		mockTeamRepo.EXPECT().Update(team).Return(errors.New("failed to update team"))

		_, err := teamUsecase.UpdateTeam("team-123", "Team A")
		assert.Error(t, err)
		assert.Equal(t, "failed to update team", err.Error())
	})
}

func TestDeleteTeam(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTeamRepo := mock.NewMockTeamRepository(ctrl)
	mockUserRepo := mock.NewMockUserRepository(ctrl)
	mockUserTeamRepo := mock.NewMockUserTeamRepository(ctrl)

	teamUsecase := usecase.NewTeamUsecase(mockTeamRepo, mockUserRepo, mockUserTeamRepo)

	team := entity.NewTeam("Team A")

	t.Run("正常系: チームが正常に削除される", func(t *testing.T) {

		mockTeamRepo.EXPECT().FindById("team-123").Return(team, nil)
		mockTeamRepo.EXPECT().Delete("team-123").Return(nil)

		err := teamUsecase.DeleteTeam("team-123")
		assert.NoError(t, err)
	})

	t.Run("異常系: チームが見つからない", func(t *testing.T) {
		mockTeamRepo.EXPECT().FindById("team-123").Return(nil, errors.New("team not found"))

		err := teamUsecase.DeleteTeam("team-123")
		assert.Error(t, err)
		assert.Equal(t, "team not found", err.Error())
	})

	t.Run("異常系: チームの削除に失敗する", func(t *testing.T) {
		mockTeamRepo.EXPECT().FindById("team-123").Return(team, nil)
		mockTeamRepo.EXPECT().Delete("team-123").Return(errors.New("failed to delete team"))

		err := teamUsecase.DeleteTeam("team-123")
		assert.Error(t, err)
		assert.Equal(t, "failed to delete team", err.Error())
	})
}

func TestInviteUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTeamRepo := mock.NewMockTeamRepository(ctrl)
	mockUserRepo := mock.NewMockUserRepository(ctrl)
	mockUserTeamRepo := mock.NewMockUserTeamRepository(ctrl)

	teamUsecase := usecase.NewTeamUsecase(mockTeamRepo, mockUserRepo, mockUserTeamRepo)

	team := entity.NewTeam("Team A")
	user := entity.NewUser(*entity.NewUserId("user-123"), "User A", "user-123@gmail.com")
	user1 := entity.NewUser(*entity.NewUserId("newcommer1"), "Newcommer 1", "newcommer1@gmail.com")
	user2 := entity.NewUser(*entity.NewUserId("newcommer2"), "Newcommer 2", "newcommer2@gmail.com")
	userTeam1 := entity.NewUserTeam(*user1.GetId(), *team.GetId(), entity.Invited)
	userTeam2 := entity.NewUserTeam(*user2.GetId(), *team.GetId(), entity.Invited)
	targetUserEmails := []string{
		"newcommer1@gmail.com",
		"newcommer2@gmail.com",
	}

	t.Run("正常系: ユーザーが正常に招待される", func(t *testing.T) {

		// チームと招待者ユーザーの存在確認
		mockTeamRepo.EXPECT().FindById("team-123").Return(team, nil)
		mockUserRepo.EXPECT().FindById("user-123").Return(user, nil)
		mockUserTeamRepo.EXPECT().IsMember("user-123", "team-123").Return(true, nil)

		// 招待対象ユーザーの存在確認と招待
		// 1人目
		mockUserRepo.EXPECT().FindByEmail("newcommer1@gmail.com").Return(user1, nil)
		mockUserTeamRepo.EXPECT().IsMember("newcommer1", "team-123").Return(false, nil)
		mockUserTeamRepo.EXPECT().Save(userTeam1).Return(nil)
		// 2人目
		mockUserRepo.EXPECT().FindByEmail("newcommer2@gmail.com").Return(user2, nil)
		mockUserTeamRepo.EXPECT().IsMember("newcommer2", "team-123").Return(false, nil)
		mockUserTeamRepo.EXPECT().Save(userTeam2).Return(nil)

		err := teamUsecase.InviteUsers("team-123", "user-123", targetUserEmails)
		assert.NoError(t, err)
	})

	t.Run("異常系: チームが見つからない", func(t *testing.T) {
		mockTeamRepo.EXPECT().FindById("teamId_noexists").Return(nil, errors.New("team not found"))

		err := teamUsecase.InviteUsers("teamId_noexists", "user-123", nil)
		assert.Error(t, err)
		assert.Equal(t, "team not found", err.Error())
	})

	t.Run("異常系: ユーザーが見つからない", func(t *testing.T) {
		mockTeamRepo.EXPECT().FindById("team-123").Return(team, nil)
		mockUserRepo.EXPECT().FindById("user-123").Return(nil, errors.New("user not found"))

		err := teamUsecase.InviteUsers("team-123", "user-123", nil)
		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
	})

	t.Run("異常系: 招待者がチームのメンバーではない", func(t *testing.T) {
		mockTeamRepo.EXPECT().FindById("team-123").Return(team, nil)
		mockUserRepo.EXPECT().FindById("user-123").Return(user, nil)
		mockUserTeamRepo.EXPECT().IsMember("user-123", "team-123").Return(false, nil)

		err := teamUsecase.InviteUsers("team-123", "user-123", nil)
		assert.Error(t, err)
		assert.Equal(t, "inviter is not a member of the team", err.Error())
	})
}

func TestAcceptInvitation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTeamRepo := mock.NewMockTeamRepository(ctrl)
	mockUserRepo := mock.NewMockUserRepository(ctrl)
	mockUserTeamRepo := mock.NewMockUserTeamRepository(ctrl)

	teamUsecase := usecase.NewTeamUsecase(mockTeamRepo, mockUserRepo, mockUserTeamRepo)

	team := entity.NewTeam("Team A")
	user := entity.NewUser(*entity.NewUserId("user-123"), "User A", "user-123@gmail.com")
	userTeam := entity.NewUserTeam(*user.GetId(), *team.GetId(), entity.Invited)

	t.Run("正常系: 招待を受け入れる", func(t *testing.T) {
		mockTeamRepo.EXPECT().FindById("team-123").Return(team, nil)
		mockUserRepo.EXPECT().FindById("user-123").Return(user, nil)
		mockUserTeamRepo.EXPECT().UpdateState(userTeam).Return(nil)

		err := teamUsecase.AcceptInvitation("team-123", "user-123")
		assert.NoError(t, err)
	})

	t.Run("異常系: チームが見つからない", func(t *testing.T) {
		mockTeamRepo.EXPECT().FindById("team-123").Return(nil, errors.New("team not found"))

		err := teamUsecase.AcceptInvitation("team-123", "user-123")
		assert.Error(t, err)
		assert.Equal(t, "team not found", err.Error())
	})

	t.Run("異常系: ユーザーが見つからない", func(t *testing.T) {
		mockTeamRepo.EXPECT().FindById("team-123").Return(team, nil)
		mockUserRepo.EXPECT().FindById("user-123").Return(nil, errors.New("user not found"))

		err := teamUsecase.AcceptInvitation("team-123", "user-123")
		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
	})
}

func TestRemoveMember(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTeamRepo := mock.NewMockTeamRepository(ctrl)
	mockUserRepo := mock.NewMockUserRepository(ctrl)
	mockUserTeamRepo := mock.NewMockUserTeamRepository(ctrl)

	teamUsecase := usecase.NewTeamUsecase(mockTeamRepo, mockUserRepo, mockUserTeamRepo)

	team := entity.NewTeam("Team A")
	user := entity.NewUser(*entity.NewUserId("user-123"), "User A", "user-123@gmail.com")

	t.Run("正常系: メンバーが正常に削除される", func(t *testing.T) {
		mockTeamRepo.EXPECT().FindById("team-123").Return(team, nil)
		mockUserRepo.EXPECT().FindById("user-123").Return(user, nil)
		mockUserTeamRepo.EXPECT().Delete("user-123", "team-123").Return(nil)

		err := teamUsecase.RemoveMember("team-123", "user-123")
		assert.NoError(t, err)
	})

	t.Run("異常系: チームが見つからない", func(t *testing.T) {
		mockTeamRepo.EXPECT().FindById("team-123").Return(nil, errors.New("team not found"))

		err := teamUsecase.RemoveMember("team-123", "user-123")
		assert.Error(t, err)
		assert.Equal(t, "team not found", err.Error())
	})

	t.Run("異常系: ユーザーが見つからない", func(t *testing.T) {
		mockTeamRepo.EXPECT().FindById("team-123").Return(team, nil)
		mockUserRepo.EXPECT().FindById("user-123").Return(nil, errors.New("user not found"))

		err := teamUsecase.RemoveMember("team-123", "user-123")
		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
	})
}

func TestGetTeamsByUserId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTeamRepo := mock.NewMockTeamRepository(ctrl)
	mockUserRepo := mock.NewMockUserRepository(ctrl)
	mockUserTeamRepo := mock.NewMockUserTeamRepository(ctrl)

	teamUsecase := usecase.NewTeamUsecase(mockTeamRepo, mockUserRepo, mockUserTeamRepo)

	teams := []*entity.Team{
		entity.NewTeam("Team A"),
		entity.NewTeam("Team B"),
	}
	userId := "user-123"

	t.Run("正常系: ユーザーが所属するチームが正常に取得される", func(t *testing.T) {

		mockUserTeamRepo.EXPECT().FindTeamsByUserId(userId).Return([]string{"team-123", "team-456"}, nil)
		mockTeamRepo.EXPECT().FindById("team-123").Return(teams[0], nil)
		mockTeamRepo.EXPECT().FindById("team-456").Return(teams[1], nil)

		result, err := teamUsecase.GetTeamsByUserId(userId)
		assert.NoError(t, err)
		assert.Equal(t, teams, result)
	})

	t.Run("異常系: チームの取得に失敗する", func(t *testing.T) {
		userId := "user-123"

		mockUserTeamRepo.EXPECT().FindTeamsByUserId(userId).Return(nil, errors.New("failed to get teams"))

		result, err := teamUsecase.GetTeamsByUserId(userId)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "failed to get teams", err.Error())
	})
}

func TestGetInvitedTeams(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTeamRepo := mock.NewMockTeamRepository(ctrl)
	mockUserRepo := mock.NewMockUserRepository(ctrl)
	mockUserTeamRepo := mock.NewMockUserTeamRepository(ctrl)

	teamUsecase := usecase.NewTeamUsecase(mockTeamRepo, mockUserRepo, mockUserTeamRepo)

	teams := []*entity.Team{
		entity.NewTeam("Team A"),
		entity.NewTeam("Team B"),
	}
	userId := "user-123"

	t.Run("正常系: ユーザーが招待されているチームが正常に取得される", func(t *testing.T) {

		mockUserTeamRepo.EXPECT().FindInvitedTeamsByUserId(userId).Return([]string{"team-123", "team-456"}, nil)
		mockTeamRepo.EXPECT().FindById("team-123").Return(teams[0], nil)
		mockTeamRepo.EXPECT().FindById("team-456").Return(teams[1], nil)

		result, err := teamUsecase.GetInvitedTeams(userId)
		assert.NoError(t, err)
		assert.Equal(t, teams, result)
	})

	t.Run("異常系: チームの取得に失敗する", func(t *testing.T) {
		userId := "user-123"

		mockUserTeamRepo.EXPECT().FindInvitedTeamsByUserId(userId).Return(nil, errors.New("failed to get teams"))

		result, err := teamUsecase.GetInvitedTeams(userId)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "failed to get teams", err.Error())
	})
}

func TestGetMembersByTeamId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTeamRepo := mock.NewMockTeamRepository(ctrl)
	mockUserRepo := mock.NewMockUserRepository(ctrl)
	mockUserTeamRepo := mock.NewMockUserTeamRepository(ctrl)

	teamUsecase := usecase.NewTeamUsecase(mockTeamRepo, mockUserRepo, mockUserTeamRepo)
	teamId := "team-123"
	userIds := []string{"user-123", "user-456"}
	users := []*entity.User{
		entity.NewUser(*entity.NewUserId("user-123"), "User A", "user-123@gmail.com"),
		entity.NewUser(*entity.NewUserId("user-456"), "User B", "user-456@gmail.com"),
	}

	t.Run("正常系: チームのメンバーが正常に取得される", func(t *testing.T) {

		mockUserTeamRepo.EXPECT().FindMembersByTeamId(teamId).Return(userIds, nil)
		mockUserRepo.EXPECT().FindById("user-123").Return(users[0], nil)
		mockUserRepo.EXPECT().FindById("user-456").Return(users[1], nil)
		result, err := teamUsecase.GetMembersByTeamId(teamId)
		assert.NoError(t, err)
		assert.Equal(t, users, result)
	})

	t.Run("異常系: メンバーの取得に失敗する", func(t *testing.T) {
		teamId := "team-123"

		mockUserTeamRepo.EXPECT().FindMembersByTeamId(teamId).Return(nil, errors.New("failed to get members"))

		result, err := teamUsecase.GetMembersByTeamId(teamId)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "failed to get members", err.Error())
	})
}
