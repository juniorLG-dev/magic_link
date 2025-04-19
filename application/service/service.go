package service

import (
	"magic_link/port/output"
	"magic_link/adapter/output/model"

	"github.com/google/uuid"
)

type service struct {
	cache output.PortCache
	smtp  output.PortSMTP
}

func NewService(
	cache output.PortCache,
	smtp  output.PortSMTP,
) *service {
	return &service{
		cache: cache,
		smtp: smtp,
	}
}

func (s *service) SendEmail(email string) error {
	code := uuid.New().String()
	user := model.UserCode{
		Email: email,
		Code: code,
	}

	if err := s.cache.Set(user); err != nil {
		return err
	}

	if err := s.smtp.SendEmail(email, code); err != nil {
		return err
	}

	userData := model.User{
		Email: email,
	}

	model.UserData[code] = userData

	return nil
}

func (s *service) VerifyCode(code string) error {
	_, err := s.cache.Get(code)	
	if err != nil {
		return err
	}

	userData := model.UserData[code]
	userData.Checked = true

	return nil
}
