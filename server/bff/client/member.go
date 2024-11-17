package client

import (
	"crypto/tls"
	"net"
	"net/http"

	"connectrpc.com/connect"
	memberv1 "github.com/ei-sugimoto/microtodo/server/member/gen/member/v1"
	"github.com/ei-sugimoto/microtodo/server/member/gen/member/v1/memberv1connect"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/http2"
)

type MemberClient struct{}

func NewMemberClient() *MemberClient {
	return &MemberClient{}
}

type CreateRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

var httpClient *http.Client

func init() {
	httpClient = &http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}
}

func (c *MemberClient) Create(ctx echo.Context) error {

	CreateReq := CreateRequest{}
	if err := ctx.Bind(&CreateReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	client := memberv1connect.NewMemberServiceClient(
		httpClient,
		"http://member:5556",
		connect.WithGRPCWeb(),
	)
	res, err := client.Create(ctx.Request().Context(), connect.NewRequest(&memberv1.CreateRequest{
		Name:     CreateReq.Name,
		Password: CreateReq.Password,
	}))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, res)
}

func (c *MemberClient) Login(ctx echo.Context) error {
	CreateReq := CreateRequest{}
	if err := ctx.Bind(&CreateReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	client := memberv1connect.NewMemberServiceClient(
		httpClient,
		"http://member:5556",
		connect.WithGRPCWeb(),
	)
	res, err := client.Login(ctx.Request().Context(), connect.NewRequest(&memberv1.LoginRequest{
		Name:     CreateReq.Name,
		Password: CreateReq.Password,
	}))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, res.Msg)
}
