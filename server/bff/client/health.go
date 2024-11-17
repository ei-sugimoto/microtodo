package client

import (
	"fmt"
	"log/slog"
	"net/http"

	"connectrpc.com/connect"
	healthv1 "github.com/ei-sugimoto/microtodo/server/member/gen/health/v1"
	"github.com/ei-sugimoto/microtodo/server/member/gen/health/v1/healthv1connect"
	"github.com/labstack/echo/v4"
)

type HealthClient struct {
}

func NewHealthClient() *HealthClient {
	return &HealthClient{}
}

func (c *HealthClient) Health(ctx echo.Context) error {
	client := healthv1connect.NewHealthServiceClient(
		http.DefaultClient,
		"http://member:5556",
	)

	res, err := client.Health(ctx.Request().Context(), connect.NewRequest(&healthv1.HealthRequest{}))
	if err != nil {
		slog.Error(fmt.Sprintf("failed to call HealthService: %v", err))
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to call HealthService")
	}

	return ctx.JSON(http.StatusOK, res.Msg.Message)

}
