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

type Database struct {
	Driver   string `json:"driver" mapstructure:"driver"`
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
	User     string `json:"user" mapstructure:"user"`
	Password string `json:"password" mapstructure:"password"`
	Name     string `json:"name" mapstructure:"name"`
	Params   string `json:"params" mapstructure:"params"`
}

type Config struct {
	Server   *Server   `json:"server" mapstructure:"server"`
	Database *Database `json:"database" mapstructure:"database"`
}

var configFilePath = "./configs/config.yaml"

var defaultConfig = Config{
	Server: &Server{
		Port:    8080,
		Address: "127.0.0.1",
	},
	Database: &Database{
		Driver: "sqlite",
		Name:   "./nglab.db",
	},
}

func New() Config {
	conf := defaultConfig

	viper.SetConfigFile(configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		panic(errors.New("failed to load config file: " + err.Error()))
	}

	if err := viper.Unmarshal(&conf); err != nil {
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

func (d *Database) DSN() string {
	if err := validator.New().Struct(d); err != nil {
		return "./db.sqlite3"
	}

	if d.Driver == "sqlite3" {
		return d.Name
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		d.User, d.Password, d.Host, d.Port, d.Name, d.Params,
	)
}