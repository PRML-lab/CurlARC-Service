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

// func TestAppendEndData(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRecordRepo := mock.NewMockRecordRepository(ctrl)
// 	mockUserTEamRepo := mock.NewMockUserTeamRepository(ctrl)
// 	mockTeamRepo := mock.NewMockTeamRepository(ctrl)

// 	recordUsecase := usecase.NewRecordUsecase(
// 		mockRecordRepo,
// 		mockUserTEamRepo,
// 		mockTeamRepo,
// 	)

// 	recordId := "record-123"
// 	userId := "user-123"
// 	teamId := "team-123"
// 	enemyTeamName := "Team B"
// 	place := "Tokyo"
// 	date := time.Date(2023, 9, 19, 0, 0, 0, 0, time.UTC)

// 	record, _ := entity.NewRecord(
// 		teamId,
// 		entity.WithEnemyTeamName(enemyTeamName),
// 		entity.WithPlace(place),
// 		entity.WithDate(date),
// 	)

// 	endsData := []entity.DataPerEnd{
// 		{
// 			Score: 1,
// 			Shots: []entity.Shot{
// 				{
// 					Type:        "Draw",
// 					SuccessRate: 0.85,
// 					Shooter:     "Lead",
// 					Stones: entity.Stones{
// 						FriendStones: []entity.Coordinate{
// 							{Index: 1, R: 2.1, Theta: 0.1},
// 						},
// 						EnemyStones: []entity.Coordinate{},
// 					},
// 				},
// 				{
// 					Type:        "Guard",
// 					SuccessRate: 0.80,
// 					Shooter:     "Lead",
// 					Stones: entity.Stones{
// 						FriendStones: []entity.Coordinate{
// 							{Index: 1, R: 2.1, Theta: 0.1},
// 						},
// 						EnemyStones: []entity.Coordinate{
// 							{Index: 1, R: 2.5, Theta: 3.14},
// 						},
// 					},
// 				},
// 				{
// 					Type:        "Draw",
// 					SuccessRate: 0.82,
// 					Shooter:     "Second",
// 					Stones: entity.Stones{
// 						FriendStones: []entity.Coordinate{
// 							{Index: 1, R: 2.1, Theta: 0.1},
// 							{Index: 2, R: 1.5, Theta: 0.2},
// 						},
// 						EnemyStones: []entity.Coordinate{
// 							{Index: 1, R: 2.5, Theta: 3.14},
// 						},
// 					},
// 				},
// 				{
// 					Type:        "Takeout",
// 					SuccessRate: 0.78,
// 					Shooter:     "Second",
// 					Stones: entity.Stones{
// 						FriendStones: []entity.Coordinate{
// 							{Index: 1, R: 2.1, Theta: 0.1},
// 							{Index: 2, R: 1.5, Theta: 0.2},
// 						},
// 						EnemyStones: []entity.Coordinate{
// 							{Index: 1, R: 2.5, Theta: 3.14},
// 						},
// 					},
// 				},
// 				{
// 					Type:        "Draw",
// 					SuccessRate: 0.88,
// 					Shooter:     "Third",
// 					Stones: entity.Stones{
// 						FriendStones: []entity.Coordinate{
// 							{Index: 1, R: 2.1, Theta: 0.1},
// 							{Index: 2, R: 1.5, Theta: 0.2},
// 							{Index: 3, R: 0.8, Theta: 0.1},
// 						},
// 						EnemyStones: []entity.Coordinate{
// 							{Index: 1, R: 2.5, Theta: 3.14},
// 						},
// 					},
// 				},
// 				{
// 					Type:        "Hit and Roll",
// 					SuccessRate: 0.75,
// 					Shooter:     "Third",
// 					Stones: entity.Stones{
// 						FriendStones: []entity.Coordinate{
// 							{Index: 1, R: 2.1, Theta: 0.1},
// 							{Index: 2, R: 1.5, Theta: 0.2},
// 							{Index: 3, R: 0.8, Theta: 0.1},
// 						},
// 						EnemyStones: []entity.Coordinate{
// 							{Index: 1, R: 1.2, Theta: 3.0},
// 						},
// 					},
// 				},
// 				{
// 					Type:        "Freeze",
// 					SuccessRate: 0.90,
// 					Shooter:     "Skip",
// 					Stones: entity.Stones{
// 						FriendStones: []entity.Coordinate{
// 							{Index: 1, R: 2.1, Theta: 0.1},
// 							{Index: 2, R: 1.5, Theta: 0.2},
// 							{Index: 3, R: 0.8, Theta: 0.1},
// 							{Index: 4, R: 0.3, Theta: 0.05},
// 						},
// 						EnemyStones: []entity.Coordinate{
// 							{Index: 1, R: 1.2, Theta: 3.0},
// 						},
// 					},
// 				},
// 				{
// 					Type:        "Draw",
// 					SuccessRate: 0.92,
// 					Shooter:     "Skip",
// 					Stones: entity.Stones{
// 						FriendStones: []entity.Coordinate{
// 							{Index: 1, R: 2.1, Theta: 0.1},
// 							{Index: 2, R: 1.5, Theta: 0.2},
// 							{Index: 3, R: 0.8, Theta: 0.1},
// 							{Index: 4, R: 0.3, Theta: 0.05},
// 							{Index: 5, R: 0.5, Theta: 3.14},
// 						},
// 						EnemyStones: []entity.Coordinate{
// 							{Index: 1, R: 1.2, Theta: 3.0},
// 						},
// 					},
// 				},
// 			},
// 		},
// 		// 追加のエンドをここに挿入（省略）
// 	}

// 	invalidEndsData := []entity.DataPerEnd{
// 		{
// 			Score: 2,
// 			Shots: []entity.Shot{
// 				{
// 					Type:        "",
// 					SuccessRate: 0,
// 					Shooter:     "",
// 					Stones: entity.Stones{
// 						FriendStones: []entity.Coordinate{
// 							{Index: 0, R: 0, Theta: 0},
// 							{Index: 0, R: 0, Theta: 0},
// 						},
// 						EnemyStones: []entity.Coordinate{
// 							{Index: 0, R: 0, Theta: 0},
// 							{Index: 0, R: 0, Theta: 0},
// 						},
// 					},
// 				},
// 			},
// 		},
// 		{
// 			Score: 2,
// 			Shots: []entity.Shot{
// 				{
// 					Type:        "",
// 					SuccessRate: 0,
// 					Shooter:     "",
// 					Stones: entity.Stones{
// 						FriendStones: []entity.Coordinate{
// 							{Index: 0, R: 0, Theta: 0},
// 							{Index: 0, R: 0, Theta: 0},
// 						},
// 						EnemyStones: []entity.Coordinate{
// 							{Index: 0, R: 0, Theta: 0},
// 							{Index: 0, R: 0, Theta: 0},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	// t.Run("正常系: endsDataが正常に追加される", func(t *testing.T) {

// 	// 	mockRecordRepo.EXPECT().FindByRecordId(recordId).Return(record, nil)
// 	// 	mockUserTEamRepo.EXPECT().IsMember(userId, teamId).Return(true, nil)
// 	// 	mockRecordRepo.EXPECT().Update(gomock.Any()).Return(record, nil)

// 	// 	updatedRecord, err := recordUsecase.AppendEndData(recordId, userId, endsData)
// 	// 	assert.NoError(t, err)
// 	// 	assert.NotNil(t, updatedRecord)
// 	// 	assert.Equal(t, 1, len(updatedRecord.GetEndsData()))
// 	// })

// 	t.Run("異常系: レコードが見つからない", func(t *testing.T) {
// 		mockRecordRepo.EXPECT().FindByRecordId(recordId).Return(nil, errors.New("record not found"))

// 		updatedRecord, err := recordUsecase.AppendEndData(recordId, userId, endsData)
// 		assert.Error(t, err)
// 		assert.Nil(t, updatedRecord)
// 		assert.Equal(t, "record not found", err.Error())
// 	})

// 	t.Run("異常系: ユーザーがチームに所属していない", func(t *testing.T) {
// 		mockRecordRepo.EXPECT().FindByRecordId(recordId).Return(record, nil)
// 		mockUserTEamRepo.EXPECT().IsMember(userId, teamId).Return(false, nil)

// 		updatedRecord, err := recordUsecase.AppendEndData(recordId, userId, endsData)
// 		assert.Error(t, err)
// 		assert.Nil(t, updatedRecord)
// 		assert.Equal(t, "appender is not a member of the team", err.Error())
// 	})

// 	t.Run("異常系: endsDataの検証に失敗する", func(t *testing.T) {
// 		mockRecordRepo.EXPECT().FindByRecordId(recordId).Return(record, nil)
// 		mockUserTEamRepo.EXPECT().IsMember(userId, teamId).Return(true, nil)

// 		updatedRecord, err := recordUsecase.AppendEndData(recordId, userId, invalidEndsData)
// 		assert.Error(t, err)
// 		assert.Nil(t, updatedRecord)
// 	})
// }
