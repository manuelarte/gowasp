package main

import (
	"database/sql"
	"html/template"

	"github.com/manuelarte/gowasp/internal/config"
	"github.com/manuelarte/gowasp/internal/handlers"
	"github.com/manuelarte/gowasp/internal/repositories"
	"github.com/manuelarte/gowasp/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/manuelarte/pagorminator"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func renderUnsafe(s string) template.HTML {
	//#nosec G203
	return template.HTML(s)
}

func main() {
	db, err := config.MigrateDatabase()
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	gormDB, err := gorm.Open(sqlite.New(sqlite.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	_ = gormDB.Use(pagorminator.PaGormMinator{})
	userService := services.UserServiceImpl{Repository: repositories.UserRepositoryDB{DB: gormDB}}
	postService := services.PostServiceImpl{Repository: repositories.PostRepositoryDB{DB: gormDB}}
	postCommentService := services.PostCommentServiceImpl{Repository: repositories.PostCommentRepositoryDB{DB: gormDB}}

	usersHandler := handlers.UsersHandler{UserService: userService, PostService: postService}
	postsHandler := handlers.PostsHandler{PostService: postService, PostCommentService: postCommentService}
	postCommentHandler := handlers.PostCommentsHandler{PostCommentService: postCommentService}

	config.RegisterErrorResponseHandlers()
	r := gin.Default()
	configCors := cors.DefaultConfig()
	configCors.AllowOrigins = []string{"http://localhost:8080", "http://localhost:63342"}
	configCors.AllowCredentials = true
	r.Use(cors.New(configCors))
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.SetFuncMap(template.FuncMap{
		"unsafe": renderUnsafe,
	})
	r.Static("/css", "web/css")
	r.LoadHTMLGlob("web/templates/**/*")

	r.GET("/users/signup", usersHandler.SignupPage)
	r.GET("/users/login", usersHandler.LoginPage)

	r.GET("/users/welcome", config.AuthMiddleware(), usersHandler.WelcomePage)
	r.GET("/static/posts", config.AuthMiddleware(), postsHandler.GetStaticPostFileByName)
	r.GET("/posts/:id/view", config.AuthMiddleware(), postsHandler.ViewPostPage)

	r.GET("/debug", handlers.GetTemplateByName)

	r.POST("/users/signup", usersHandler.Signup)
	r.POST("/users/login", usersHandler.Login)
	r.DELETE("/users/logout", usersHandler.Logout)

	r.GET("/posts", postsHandler.GetAll)
	r.GET("/posts/:id/comments", postCommentHandler.GetPostComments)
	r.POST("/posts/:id/comments", config.AuthMiddleware(), postCommentHandler.CreatePostComment)

	err = r.Run()
	if err != nil {
		panic(err)
	}
}
