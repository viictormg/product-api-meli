package infra

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func NewHTTPServer(lc fx.Lifecycle) *echo.Echo {
	srv := &http.Server{Addr: ":8080"}
	echoServer := echo.New() 

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := echoServer.StartServer(srv); err != nil {
					fmt.Println("Failed to start HTTP server", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return echoServer
}
