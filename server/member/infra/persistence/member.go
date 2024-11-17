package persistence

import (
	"context"
	"database/sql"

	"github.com/ei-sugimoto/microtodo/server/member/domain"
	"github.com/ei-sugimoto/microtodo/server/member/repository"
)

type Member struct {
	db *sql.DB
}

func NewMember(db *sql.DB) repository.Member {
	return &Member{db: db}
}

func (m *Member) Create(ctx context.Context, member *domain.Member) error {
	_, err := m.db.Exec("INSERT INTO members (name, password) VALUES (?, ?)", member.Name, member.Password)
	if err != nil {
		return err
	}
	return nil
}
