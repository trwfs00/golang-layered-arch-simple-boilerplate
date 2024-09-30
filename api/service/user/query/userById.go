package query

import (
	"boilerplate/api/repository"
	"boilerplate/lib/database/entity"
)

type GetUserByIdService interface {
	Execute(id int) (*entity.User, error)
}

type getUserByIdService struct {
	userRepo repository.UserRepository
}

func NewGetUserByIdService(userRepo repository.UserRepository) GetUserByIdService {
	return &getUserByIdService{userRepo: userRepo}
}

func (s *getUserByIdService) Execute(id int) (*entity.User, error) {
	// TODO: Add validation or if it is required more business logic
	return s.userRepo.GetUserById(id)
}
