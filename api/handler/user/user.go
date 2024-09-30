package user

import (
	"boilerplate/api/service/user/command"
	"boilerplate/api/service/user/query"
)

type UserHandler struct {
	getUserByIdService query.GetUserByIdService
	createUserService  command.CreateUserService
}

// declare a new user handler
func NewUserHandler(getUserByIdService query.GetUserByIdService, createUserService command.CreateUserService) *UserHandler {
	return &UserHandler{
		getUserByIdService: getUserByIdService,
		createUserService:  createUserService,
	}
}
