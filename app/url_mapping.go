package app

import (
	"github.com/khaihoan0805/bookstore_users-api/controllers/users"

	"github.com/khaihoan0805/bookstore_users-api/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.Get)
	router.POST("/users", users.Create)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
}
