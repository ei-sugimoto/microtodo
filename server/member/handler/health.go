package handler

import (
	"context"

	"connectrpc.com/connect"
	healthv1 "github.com/ei-sugimoto/microtodo/server/member/gen/health/v1"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(ctx context.Context, req *connect.Request[healthv1.HealthRequest]) (*connect.Response[healthv1.HealthResponse], error) {
	resp := connect.NewResponse(&healthv1.HealthResponse{
		Message: "healthy",
	})

	return resp, nil
}
