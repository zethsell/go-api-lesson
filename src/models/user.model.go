package models

import (
	"api/src/security"
	"errors"
	"github.com/badoux/checkmail"
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

func (user *User) Prepare(action string) error {
	if err := user.validate(action); err != nil {
		return err
	}

	if err := user.format(action); err != nil {
		return err
	}

	return nil
}

func (user *User) validate(action string) error {
	if user.Name == "" {
		return errors.New("o nome é obrigatório")
	}
	if user.Nick == "" {
		return errors.New("o nick é obrigatório")
	}
	if user.Email == "" {
		return errors.New("o email é obrigatório")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("email invalido")
	}

	if action == "store" && user.Password == "" {
		return errors.New("a senha é obrigatória")
	}
	return nil
}

func (user *User) format(action string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)

	if action == "store" {
		hash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hash)

	}

	return nil
}
