package config

type Config struct {
	Args *Args
}

func NewConfig() *Config {
	return &Config{
		Args: &Args{},
	}
}
