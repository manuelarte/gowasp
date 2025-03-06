package config

type Config struct {
	MigrationSourceURL string `env:"MIGRATION_SOURCE_URL" envDefault:"file://./resources/migrations"`
	TemplatesPath      string `env:"TEMPLATES_PATH" envDefault:"web/templates/**/*"`
}
