package gowasp

import "embed"

//go:embed resources/migrations/*
var MigrationsFolder embed.FS

//go:embed openapi.yaml
var OpenAPI []byte

//go:embed static/swagger-ui/*
var SwaggerUI embed.FS

//go:embed frontend/dist/*
var Frontend embed.FS
