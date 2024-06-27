package config

type Auth struct {
	Token string `envconfig:"TOKEN" env-required:"true"`
}
