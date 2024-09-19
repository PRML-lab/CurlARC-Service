package usecase_test

import (
	"CurlARC/internal/domain/entity"
	"CurlARC/internal/usecase"
	"CurlARC/mock"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateRecord(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRecordRepo := mock.NewMockRecordRepository(ctrl)
	mockUserTEamRepo := mock.NewMockUserTeamRepository(ctrl)
	mockTeamRepo := mock.NewMockTeamRepository(ctrl)

	recordUsecase := usecase.NewRecordUsecase(
		mockRecordRepo,
		mockUserTEamRepo,
		mockTeamRepo,
	)

	userId := "user-123"
	teamId := "team-123"
	enemyTeamName := "Team B"
	place := "Tokyo"
	date := time.Date(2023, 9, 19, 0, 0, 0, 0, time.UTC)
	record, _ := entity.NewRecord(
		teamId,
		entity.WithEnemyTeamName(enemyTeamName),
		entity.WithPlace(place),
		entity.WithDate(date),
	)

	t.Run("正常系: レコードが正常に作成される", func(t *testing.T) {
		mockUserTEamRepo.EXPECT().IsMember(userId, teamId).Return(true, nil)
		mockTeamRepo.EXPECT().FindById(teamId).Return(nil, nil)
		mockRecordRepo.EXPECT().Save(gomock.Any()).Return(record, nil)

		createdRecord, err := recordUsecase.CreateRecord(
			userId,
			teamId,
			enemyTeamName,
			place,
			entity.Win,
			date,
		)
		assert.NoError(t, err)
		assert.Equal(t, record, createdRecord)
	})

	t.Run("異常系: dbへの保存に失敗する", func(t *testing.T) {
		mockUserTEamRepo.EXPECT().IsMember(userId, teamId).Return(true, nil)
		mockTeamRepo.EXPECT().FindById(teamId).Return(nil, nil)
		mockRecordRepo.EXPECT().Save(gomock.Any()).Return(nil, errors.New("failed to save record"))

		createdRecord, err := recordUsecase.CreateRecord(
			userId,
			teamId,
			enemyTeamName,
			place,
			entity.Win,
			date,
		)

		assert.Error(t, err)
		assert.Nil(t, createdRecord)
		assert.Equal(t, "failed to save record", err.Error())
	})

	t.Run("異常系: チームが見つからない", func(t *testing.T) {
		mockUserTEamRepo.EXPECT().IsMember(userId, teamId).Return(true, nil)
		mockTeamRepo.EXPECT().FindById(teamId).Return(nil, errors.New("team not found"))

		createdRecord, err := recordUsecase.CreateRecord(
			userId,
			teamId,
			enemyTeamName,
			place,
			entity.Win,
			date,
		)

		assert.Error(t, err)
		assert.Nil(t, createdRecord)
		assert.Equal(t, "team not found", err.Error())
	})

	t.Run("異常系: ユーザーがチームに所属していない", func(t *testing.T) {
		mockUserTEamRepo.EXPECT().IsMember(userId, teamId).Return(false, nil)

		createdRecord, err := recordUsecase.CreateRecord(
			userId,
			teamId,
			enemyTeamName,
			place,
			entity.Win,
			date,
		)

		assert.Error(t, err)
		assert.Nil(t, createdRecord)
		assert.Equal(t, "user is not a member of the team", err.Error())
	})
}
