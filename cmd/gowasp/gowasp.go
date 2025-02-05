package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"gowasp/internal/config"
	"gowasp/internal/handlers"
	"gowasp/internal/repositories"
	"gowasp/internal/services"
	"log"
)

func main() {
	db, err := config.MigrateDatabase()
	if err != nil {
		log.Fatal(err)
	}

	userService := services.UserServiceImpl{Repository: repositories.UserRepositoryDB{DB: db}}
	userHandler := handlers.UserHandler{UserService: userService}

	config.RegisterErrorResponseHandlers()
	r := gin.Default()
	r.POST("/users", userHandler.Create)

	err = r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
