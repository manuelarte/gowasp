package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	gormDB, err := gorm.Open(sqlite.New(sqlite.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	userService := services.UserServiceImpl{Repository: repositories.UserRepositoryDB{DB: gormDB}}
	userHandler := handlers.UserHandler{UserService: userService}

	config.RegisterErrorResponseHandlers()
	r := gin.Default()
	r.POST("/users", userHandler.Create)

	err = r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
