package config

type Config struct {
	Address string `env:"ADDRESS" envDefault:":8083"`
	WebPath string `env:"WEB_PATH" envDefault:"web"`
}
