package config

type App struct {
	Port string `env:"PORT"`
	Host string `env:"HOST"`
	Db   struct {
		Host       string `env:"DB_HOST"`
		Port       string `env:"DB_PORT"`
		Name       string `env:"DB_NAME"`
		Pw         string `env:"DB_PASSWORD"`
		Usr        string `env:"DB_USERNAME"`
		Collection string `env:"DB_COLLECTION"`
	}
	I18n struct {
		LocaleDir string   `env:"I18N_LOCALE_DIR"`
		Fallback  string   `env:"I18N_FALLBACK"`
		Locales   []string `env:"I18N_LOCALES"`
	}
	JWT struct {
		Secret   string `env:"JWT_SECRET"`
		Domain   string `env:"JWT_DOMAIN"`
		Secure   bool   `env:"JWT_SECURE"`
		HTTPOnly bool   `env:"JWT_HTTP_ONLY"`
	}
}