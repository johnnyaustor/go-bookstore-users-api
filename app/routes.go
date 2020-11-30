package app

import (
	"github.com/johnnyaustor/go-bookstore-users-api/app/controllers/ping"
	"github.com/johnnyaustor/go-bookstore-users-api/app/controllers/users"
)

func routes() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.Create)
	router.GET("/users/:id", users.Get)
	router.PUT("/users/:id", users.Update)
	router.PATCH("/users/:id", users.Update)
	router.DELETE("/users/:id", users.Delete)
	router.GET("/internal/users/search", users.Search)
}
