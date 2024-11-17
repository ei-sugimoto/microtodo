package domain

import (
	"encoding/base64"
	"fmt"
)

type Member struct {
	ID       int
	Name     string
	Password string
}

var ErrNameRequired = fmt.Errorf("name is required")
var ErrPasswordRequired = fmt.Errorf("password is required")

func NewMember(name, password string) (*Member, error) {
	if err := validate(name, password); err != nil {
		return nil, err
	}
	passwordHash := passwordHash(password)
	return &Member{
		Name:     name,
		Password: passwordHash,
	}, nil
}

func validate(name, password string) error {
	if name == "" {
		return ErrNameRequired
	}
	if password == "" {
		return ErrPasswordRequired
	}
	return nil
}

func passwordHash(password string) string {
	return base64.StdEncoding.EncodeToString([]byte(password))
}
