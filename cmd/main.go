package main

import (
	"CurlARC/internal/handler"
	"CurlARC/internal/injector"
	"fmt"

	"github.com/labstack/echo"
)

// func main() {
// 	if err := run(context.Background()); err != nil {
// 		slog.Error("failed to terminated server", "error", err)
// 		os.Exit(1)
// 	}
// }

// func run(ctx context.Context) error {
// 	if err := config.Init(); err != nil {
// 		return err
// 	}

// 	srv := server.NewServer()
// 	return srv.Run(ctx)
// }

func main() {
	fmt.Println("sever start")
	userHandler := injector.InjectUserHandler()
	e := echo.New()
	handler.InitRouting(e, userHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
