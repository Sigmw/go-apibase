package models

import (
	"errors"
	"sanctum/security"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID         uint64    `json:"id,omitempty"`
	Name       string    `json:"nome,omitempty"`
	Nick       string    `json:"nick,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"senha,omitempty"`
	Created    time.Time `json:"CriadoEm,omitempty"`
	Desativado bool
}

func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.format(step); err != nil {
		return err
	}

	return nil
}

func (user *User) validate(etapa string) error {
	if user.Name == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}

	if user.Nick == "" {
		return errors.New("o nick é obrigatório e não pode estar em branco")
	}

	if user.Email == "" {
		return errors.New("o e-mail é obrigatório e não pode estar em branco")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("o e-mail inserido é inválido")
	}

	if etapa == "cadastro" && user.Password == "" {
		return errors.New("a senha é obrigatória e não pode estar em branco")
	}

	return nil
}

func (user *User) format(etapa string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if etapa == "cadastro" {
		senhaComHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(senhaComHash)
	}

	return nil
}
