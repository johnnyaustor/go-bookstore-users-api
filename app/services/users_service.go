package services

import (
	"github.com/johnnyaustor/go-bookstore-users-api/app/domain/users"
	"github.com/johnnyaustor/go-bookstore-users-api/app/utils"
	"github.com/johnnyaustor/go-bookstore-users-api/app/utils/errors"
)

/**
* Menerapkan konsep implement interface
* class UsersService implement UsersServiceInterface
 */

var (
	UsersService usersServiceInterface = &usersService{}
)
type usersService struct { }
type usersServiceInterface interface {
	GetUser(id int64) (*users.User, *errors.RestError)
	CreateUser(user users.User) (*users.User, *errors.RestError)
	UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError)
	DeleteUser(userId int64) *errors.RestError
	SearchUsers(status string) (users.Users, *errors.RestError)
}

func (s *usersService) GetUser(id int64) (*users.User, *errors.RestError) {
	result := &users.User{Id: id}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.Password = utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError) {
	currentUser := &users.User{Id: user.Id}
	if restError := currentUser.Get(); restError != nil {
		return nil, restError
	}


	if isPartial {
		if user.FirstName != "" {
			currentUser.FirstName = user.FirstName
		}
		if user.LastName != "" {
			currentUser.LastName = user.LastName
		}
		if user.Email != "" {
			if restError := user.Validate(); restError != nil {
				return nil, restError
			}
			currentUser.Email = user.Email
		}
	} else {
		if restError := user.Validate(); restError != nil {
			return nil, restError
		}
		currentUser.FirstName = user.FirstName
		currentUser.LastName = user.LastName
		currentUser.Email = user.Email
	}

	if restError := currentUser.Update(); restError != nil {
		return nil, restError
	}
	return currentUser, nil
}

func (s *usersService) DeleteUser(userId int64) *errors.RestError {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (s *usersService) SearchUsers(status string) (users.Users, *errors.RestError) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
