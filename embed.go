package gowasp

import "embed"

//go:embed resources/migrations/*
var MigrationsFolder embed.FS
