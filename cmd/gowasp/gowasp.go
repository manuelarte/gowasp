package main

import (
	"context"
	"database/sql"
	"io/fs"
	"log/slog"
	"net/http"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/manuelarte/pagorminator"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/manuelarte/gowasp"
	"github.com/manuelarte/gowasp/internal/api/html"
	"github.com/manuelarte/gowasp/internal/api/rest"
	"github.com/manuelarte/gowasp/internal/config"
	"github.com/manuelarte/gowasp/internal/posts"
	"github.com/manuelarte/gowasp/internal/posts/postcomments"
	"github.com/manuelarte/gowasp/internal/users"
)

//go:generate go tool oapi-codegen -config ../../cfg.yaml ../../openapi.yaml
//go:generate go tool gospecpaths --package rest --output ../../internal/api/rest/paths.gen.go ../../openapi.yaml
func main() {
	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	cfg, err := env.ParseAs[config.Config]()
	if err != nil {
		logger.ErrorContext(ctx, "error parsing the configuration", "error", err)

		return
	}

	db, err := config.MigrateDatabase(gowasp.MigrationsFolder)
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

	r := gin.Default()
	r.Use(
		cors.New(configureConfigCors()),
		sessions.Sessions("mysession", cookie.NewStore([]byte("secret"))),
	)
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/web")
	})
	htmlPosts := html.NewPosts(postService, postCommentService)
	html.RegisterPostsHandlers(r, htmlPosts)
	{
		sfs, _ := fs.Sub(fs.FS(gowasp.SwaggerUI), "static/swagger-ui")
		r.StaticFS("swagger", http.FS(sfs))
	}
	{
		sfs, _ := fs.Sub(fs.FS(gowasp.Web), "web/dist")
		r.StaticFS("web", http.FS(sfs))
	}

	r.GET("/api/docs", func(c *gin.Context) {
		_, _ = c.Writer.Write(gowasp.OpenAPI)
	})

	{
		// Rest API
		restAPI := rest.API{
			UsersHandler:    rest.NewUsers(userService),
			CommentsHandler: rest.NewComments(postCommentService),
			PostsHandler:    rest.NewPosts(postService),
		}
		rest.RegisterHandlers(r, restAPI)
	}

	err = r.Run(cfg.Address)
	if err != nil {
		logger.ErrorContext(ctx, "error running the application", "error", err)

		return
	}
}

func configureConfigCors() cors.Config {
	configCors := cors.DefaultConfig()
	configCors.AllowOrigins = []string{"http://localhost:3000", "http://localhost:8083", "http://localhost:63342"}
	configCors.AllowCredentials = true
	// TODO(manuelarte): I can't make axios to read the Set-Cookie header, so I'm setting it as a header
	configCors.AddExposeHeaders("X-XSRF-TOKEN")
	configCors.AddAllowMethods("GET, POST, PUT, DELETE, OPTIONS")

	return configCors
}
