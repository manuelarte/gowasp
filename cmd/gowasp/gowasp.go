package main

import (
	"database/sql"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.LoadHTMLGlob("web/templates/**/*")

	r.GET("/users/signup", userHandler.SignupPage)
	r.GET("/users/login", userHandler.LoginPage)

	r.GET("/users/welcome", config.AuthMiddleware(), userHandler.WelcomePage)

	r.POST("/users/signup", userHandler.Signup)
	r.POST("/users/login", userHandler.Login)
	r.DELETE("/users/logout", userHandler.Logout)

	err = r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
