package gowasp

import "embed"

//go:embed resources/migrations/*
var MigrationsFolder embed.FS

//go:embed openapi.yaml
var OpenAPI []byte
