package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log/slog"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/manuelarte/pagorminator"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/manuelarte/gowasp/internal/api/html"
	"github.com/manuelarte/gowasp/internal/api/rest"
	"github.com/manuelarte/gowasp/internal/config"
	"github.com/manuelarte/gowasp/internal/posts"
	"github.com/manuelarte/gowasp/internal/posts/postcomments"
	"github.com/manuelarte/gowasp/internal/users"
)

//go:generate go tool oapi-codegen -config ../../cfg.yaml ../../openapi.yaml
func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	cfg, err := env.ParseAs[config.Config]()
	if err != nil {
		logger.Error("error parsing the configuration", "error", err)

		return
	}

	db, err := config.MigrateDatabase(cfg.MigrationSourceURL)
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

	_ = gormDB.Use(pagorminator.PaGorminator{})
	userService := users.NewService(users.NewRepository(gormDB))
	postService := posts.NewService(posts.NewRepository(gormDB))
	postCommentService := postcomments.NewService(postcomments.NewRepository(gormDB))

	config.RegisterErrorResponseHandlers()
	r := gin.Default()
	configCors := cors.DefaultConfig()
	configCors.AllowOrigins = []string{"http://localhost:8083", "http://localhost:63342"}
	configCors.AllowCredentials = true
	r.Use(cors.New(configCors))
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.SetFuncMap(template.FuncMap{
		"unsafe": renderUnsafe,
	})
	r.Static("/css", fmt.Sprintf("%s/css", cfg.WebPath))
	r.LoadHTMLGlob(fmt.Sprintf("%s%s", cfg.WebPath, "/templates/**/*"))

	{
		htmlUsers := html.NewUsers(postService)
		htmlPosts := html.NewPosts(postService, postCommentService)
		html.RegisterUsersHandlers(r, htmlUsers)
		html.RegisterPostsHandlers(r, htmlPosts)
		html.RegisterDebugHandlers(r)
	}

	{
		// Rest API
		restAPI := rest.API{
			Users:    rest.NewUsers(userService),
			Comments: rest.NewComments(postCommentService),
			Posts:    rest.NewPosts(postService),
		}
		rest.RegisterHandlers(r, restAPI)
	}

	err = r.Run(cfg.Address)
	if err != nil {
		logger.Error("error running the application", "error", err)

		return
	}
}

func renderUnsafe(s string) template.HTML {
	//#nosec G203
	return template.HTML(s)
}
