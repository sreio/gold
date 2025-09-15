package config

type Web struct {
	Host  string `yaml:"host"`
	Port  int    `yaml:"port"`
	Model string `yaml:"model"`
	Token string `yaml:"token"`
}
