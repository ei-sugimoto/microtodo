package usecase

import (
	"context"

	"github.com/ei-sugimoto/microtodo/server/member/domain"
	"github.com/ei-sugimoto/microtodo/server/member/repository"
)

type Member struct {
	MemberRepository repository.Member
}

func NewMember(memberRepository repository.Member) *Member {
	return &Member{MemberRepository: memberRepository}
}

func (m *Member) Create(ctx context.Context, name, password string) error {
	member, err := domain.NewMember(name, password)
	if err != nil {
		return err
	}
	return m.MemberRepository.Create(ctx, member)
}
