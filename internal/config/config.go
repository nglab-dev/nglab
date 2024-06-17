package config

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/nglab-dev/nglab/internal/utils"
	"github.com/spf13/viper"
)

type Server struct {
	Port    int    `json:"port" mapstructure:"port"`
	Address string `json:"address" mapstructure:"address"`
}

type Config struct {
	Server *Server `json:"server" mapstructure:"server"`
}

var configFilePath = "./configs/config.yaml"

var defaultConfig = &Config{
	Server: &Server{
		Port:    8080,
		Address: "127.0.0.1",
	},
}

func New() *Config {
	conf := defaultConfig

	viper.SetConfigFile(configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		panic(errors.New("failed to load config file: " + err.Error()))
	}

	if err := viper.Unmarshal(conf); err != nil {
		panic(errors.New("failed to unmarshal config file: " + err.Error()))
	}

	return conf
}

func SetConfigFilePath(path string) {
	if !utils.IsFile(path) {
		panic(errors.New("config file not found: " + path))
	}
	configFilePath = path
}

func (s *Server) ListenAddr() string {
	if err := validator.New().Struct(s); err != nil {
		return "0.0.0.0:8080"
	}
	return fmt.Sprintf("%s:%d", s.Address, s.Port)
}
