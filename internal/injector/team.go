package injector

import (
	"CurlARC/internal/domain/repository"
	"CurlARC/internal/handler"
	"CurlARC/internal/infra"
	"CurlARC/internal/usecase"
)

func InjectTeamRepository() repository.TeamRepository {
	sqlHandler := InjectDB()
	return infra.NewTeamRepository(sqlHandler)
}

func InjectUserTeamRepository() repository.UserTeamRepository {
	sqlHandler := InjectDB()
	return infra.NewUserTeamRepository(sqlHandler)
}

func InjectTeamUsecase() usecase.TeamUsecase {
	teamRepo := InjectTeamRepository()
	userRepo := InjectUserRepository()
	userTeamRepo := InjectUserTeamRepository()
	return usecase.NewTeamUsecase(teamRepo, userRepo, userTeamRepo)
}

func InjectTeamHandler() handler.TeamHandler {
	teamUsecase := InjectTeamUsecase()
	return handler.NewTeamHandler(teamUsecase)
}
