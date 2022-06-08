package hcloud

type Config struct {
	Api   string
	Token string
}

func NewConfig() *Config {
	return &Config{}
}

func LoadDefaultConfig() *Config {
	return &Config{Api: "https://app.helmut.cloud"}
}
