package command

import (
	"boilerplate/api/repository"
	"boilerplate/lib/database/entity"
)

type CreateUserService interface {
	Execute(name string, phone *string) error
}

type createUserService struct {
	userRepo repository.UserRepository
}

func NewCreateUserService(repo repository.UserRepository) CreateUserService {
	return &createUserService{userRepo: repo}
}

func (s *createUserService) Execute(name string, phone *string) error {
	user := &entity.User{
		Name:  name,
		Phone: phone,
	}
	// TODO: Add validation or if it is required more business logic
	return s.userRepo.CreateUser(user)
}
