package persistence

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ei-sugimoto/microtodo/server/member/domain"
	"github.com/ei-sugimoto/microtodo/server/member/repository"
)

var ErrMemberNotFound = fmt.Errorf("member not found")

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

func (m *Member) Login(ctx context.Context, name, password string) (*domain.Member, error) {
	var member domain.Member
	err := m.db.QueryRow("SELECT id, name, password FROM members WHERE name = ? AND password = ?", name, password).Scan(&member.ID, &member.Name, &member.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrMemberNotFound
		}
		return nil, err
	}
	return &member, nil
}
