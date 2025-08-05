package config

type Config struct {
	Address            string `env:"ADDRESS" envDefault:":8083"`
	MigrationSourceURL string `env:"MIGRATION_SOURCE_URL" envDefault:"file://./resources/migrations"`
	WebPath            string `env:"WEB_PATH" envDefault:"web"`
}
