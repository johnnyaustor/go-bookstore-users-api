package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/johnnyaustor/go-bookstore-users-api/app/domain/users"
	"github.com/johnnyaustor/go-bookstore-users-api/app/services"
	"github.com/johnnyaustor/go-bookstore-users-api/app/utils/errors"
	"net/http"
	"strconv"
)

func getUserId(userIdParam string) (int64, *errors.RestError) {
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		return 0, errors.BadRequest("user id should be number")
	}
	return userId, nil
}
func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		restError := errors.BadRequest(err.Error())
		// todo: return bad request
		c.JSON(restError.Status, restError)
		return
	}

	createUser, restError := services.UsersService.CreateUser(user)
	if restError != nil {
		// todo: handle user creation error
		c.JSON(restError.Status, restError)
		return
	}
	c.JSON(http.StatusCreated, createUser.Marshall(c.GetHeader("X-Public") == "true"))
}

func Get(c *gin.Context) {
	userId, restError := getUserId(c.Param("id"))
	if restError != nil {
		c.JSON(restError.Status, restError)
		return
	}

	user, restError := services.UsersService.GetUser(userId)
	if restError != nil {
		// todo: handle user creation error
		c.JSON(restError.Status, restError)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context) {
	userId, restError := getUserId(c.Param("id"))
	if restError != nil {
		c.JSON(restError.Status, restError)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.BadRequest(err.Error())
		c.JSON(restError.Status, restError)
		return
	}

	isPartial := c.Request.Method == http.MethodPatch

	user.Id = userId

	updateUser, restError := services.UsersService.UpdateUser(isPartial, user)
	if restError != nil {
		c.JSON(restError.Status, restError)
		return
	}
	c.JSON(http.StatusOK, updateUser.Marshall(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	userId, restError := getUserId(c.Param("id"))
	if restError != nil {
		c.JSON(restError.Status, restError)
		return
	}

	if restError := services.UsersService.DeleteUser(userId); restError != nil {
		c.JSON(restError.Status, restError)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status":"deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")

	userList, err := services.UsersService.SearchUsers(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	userList.Marshall(c.GetHeader("X-Public") == "true")
	c.JSON(http.StatusOK, userList)
}