package handler

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/ei-sugimoto/microtodo/server/member/domain"
	memberv1 "github.com/ei-sugimoto/microtodo/server/member/gen/member/v1"
	"github.com/ei-sugimoto/microtodo/server/member/infra"
	"github.com/ei-sugimoto/microtodo/server/member/infra/persistence"
	"github.com/ei-sugimoto/microtodo/server/member/usecase"
)

type MemberHandler struct{}

func NewMemberHandler() *MemberHandler {
	return &MemberHandler{}
}

func (h *MemberHandler) Create(ctx context.Context, req *connect.Request[memberv1.CreateRequest]) (*connect.Response[memberv1.CreateResponse], error) {
	db := infra.NewDB()
	defer db.Close()

	memberRepository := persistence.NewMember(db)
	memberUsecase := usecase.NewMember(memberRepository)
	err := memberUsecase.Create(ctx, req.Msg.Name, req.Msg.Password)
	if err != nil {
		if errors.Is(err, domain.ErrNameRequired) || errors.Is(err, domain.ErrPasswordRequired) {
			return nil, connect.NewError(connect.CodeInvalidArgument, err)
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	resp := connect.NewResponse(&memberv1.CreateResponse{})
	return resp, nil
}
