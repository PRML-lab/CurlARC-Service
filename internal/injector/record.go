package injector

import (
	"CurlARC/internal/domain/repository"
	"CurlARC/internal/handler"
	"CurlARC/internal/usecase"

	"CurlARC/internal/infra"
)

func InjectRecordRepository() repository.RecordRepository {
	sqlHandler := InjectDB()
	return infra.NewRecordRepository(sqlHandler)
}

func InjectRecordUsecase() usecase.RecordUsecase {
	recordRepo := InjectRecordRepository()
	teamRepo := InjectTeamRepository()
	return usecase.NewRecordUsecase(recordRepo, teamRepo)
}

func InjectRecordHandler() handler.RecordHandler {
	recordUsecase := InjectRecordUsecase()
	return handler.NewRecordHandler(recordUsecase)
}
