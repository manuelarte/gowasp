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
	usersHandler := handlers.UsersHandler{UserService: userService}

	blogsHandler := handlers.BlogsHandler{}

	config.RegisterErrorResponseHandlers()
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.LoadHTMLGlob("web/templates/**/*")

	r.GET("/users/signup", usersHandler.SignupPage)
	r.GET("/users/login", usersHandler.LoginPage)

	r.GET("/users/welcome", config.AuthMiddleware(), usersHandler.WelcomePage)

	r.POST("/users/signup", usersHandler.Signup)
	r.POST("/users/login", usersHandler.Login)
	r.DELETE("/users/logout", usersHandler.Logout)

	r.GET("/blogs", config.AuthMiddleware(), blogsHandler.GetBlogFileByName)

	err = r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
