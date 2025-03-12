package config

type Config struct {
	MigrationSourceURL string `env:"MIGRATION_SOURCE_URL" envDefault:"file://./resources/migrations"`
	WebPath            string `env:"WEB_PATH" envDefault:"web"`
}
