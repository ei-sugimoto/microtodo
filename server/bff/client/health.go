package client

import "github.com/labstack/echo/v4"

type HealthClient struct {
}

func NewHealthClient() *HealthClient {
	return &HealthClient{}
}

func (c *HealthClient) Health(ctx echo.Context) error {
	return nil
}
