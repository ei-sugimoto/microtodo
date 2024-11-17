package repository

import (
	"context"

	"github.com/ei-sugimoto/microtodo/server/member/domain"
)

type Member interface {
	Create(ctx context.Context, m *domain.Member) error
	Login(ctx context.Context, name, password string) (*domain.Member, error)
}
