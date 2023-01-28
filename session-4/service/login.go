package service

import (
	"Hacktiv8project/session-4/repository"
	"errors"

	"github.com/labstack/gommon/log"
)

type LoginSvc struct {
	loginRepo *repository.LDAPRepo
}

func NewLoginService(loginRepo *repository.LDAPRepo) *LoginSvc {
	return &LoginSvc{loginRepo: loginRepo}
}

func (s *LoginSvc) Authenticate(username, password string) (*repository.UserLDAPData, error) {
	ok, data, err := s.loginRepo.AuthUsingLDAP(username, password)
	if !ok || err != nil {
		err := errors.New("auth using ldap not ok")
		log.Error("auth using ldap not ok")
		return nil, err
	}

	return data, nil
}
