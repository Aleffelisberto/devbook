package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (u *User) Prepare() error {
	if err := u.validate(); err != nil {
		return err
	}

	u.format()

	return nil
}

func (u *User) validate() error {
	if u.Name == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}

	if u.Nick == "" {
		return errors.New("o nick é obrigatório e não pode estar em branco")
	}

	if u.Email == "" {
		return errors.New("o email é obrigatório e não pode estar em branco")
	}

	if u.Password == "" {
		return errors.New("a senha é obrigatório e não pode estar em branco")
	}

	return nil
}

func (u *User) format() {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
}
