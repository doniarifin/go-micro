package service

import (
	"go-micro/model"
)

type AuthService interface {
	Register(user *model.AuthUser) error
	Login(user *model.AuthUser) (*model.AuthUser, error)
}

type authService struct {
	repo model.AuthRepository
}

func NewAuthService(repo model.AuthRepository) AuthService {
	return &authService{repo}
}

func (s *authService) Register(r *model.AuthUser) error {
	return s.repo.Create(r)
}
func (s *authService) Login(r *model.AuthUser) (*model.AuthUser, error) {
	return s.repo.FindByEmail(r)
}
