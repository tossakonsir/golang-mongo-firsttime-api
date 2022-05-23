package services

import "example/models"

type UserService interface {
	CreateUser(*models.User) error
	Login(*models.User) error
}
