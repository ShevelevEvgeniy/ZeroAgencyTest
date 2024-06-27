package config

type HTTPServer struct {
	Port    string `envconfig:"HTTP_SERVER_PORT" env-default:"8080"`
	Prefork bool   `envconfig:"PREFORK" env-default:"true"`
}
