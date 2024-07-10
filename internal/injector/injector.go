package injector

import (
	"CurlARC/internal/domain/repository"
	"CurlARC/internal/handler"
	"CurlARC/internal/infra"
	"CurlARC/internal/usecase"
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

func InjectDB() infra.SqlHandler {
	sqlhandler := infra.NewSqlHandler()
	return *sqlhandler
}

// UserRepository (interface) に実装である SqlHandler を渡し生成する

func InjectUserRepository() repository.UserRepository {
	sqlHandler := InjectDB()
	return infra.NewUserRepository(sqlHandler)
}

func InjectFirebaseAuthClient() *auth.Client {
	opt := option.WithCredentialsFile("service_account_file.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	return authClient
}

func InjectUserUsecase() usecase.UserUsecase {
	userRepo := InjectUserRepository()
	authClient := InjectFirebaseAuthClient()
	return usecase.NewUserUsecase(userRepo, authClient)
}

func InjectUserHandler() handler.UserHandler {
	userUsecase := InjectUserUsecase()
	return handler.NewUserHandler(userUsecase)
}
