package config

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/nglab-dev/nglab/pkg/utils"
	"github.com/spf13/viper"
)

type App struct {
	Env string `json:"env" mapstructure:"env"`
}

type Server struct {
	Port    int    `json:"port" mapstructure:"port"`
	Address string `json:"address" mapstructure:"address"`
}

type Database struct {
	Dialect  string `json:"dialect" mapstructure:"dialect"`
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
	User     string `json:"user" mapstructure:"user"`
	Password string `json:"password" mapstructure:"password"`
	Name     string `json:"name" mapstructure:"name"`
	Params   string `json:"params" mapstructure:"params"`
}

type Auth struct {
	Enabled       bool     `json:"enabled" mapstructure:"enabled"`
	JWTSecret     string   `json:"jwt_secret" mapstructure:"jwt_secret"`
	JWTExpireTime int      `json:"jwt_expire_time" mapstructure:"jwt_expire_time"`
	IgnorePaths   []string `json:"ignore_paths" mapstructure:"ignore_paths"`
}

type Redis struct {
	Host     string `mapstructure:"host" yaml:"host"`
	Port     int    `mapstructure:"port" yaml:"port"`
	Password string `mapstructure:"password" yaml:"password"`
	DB       int    `mapstructure:"db" yaml:"db"`
	PoolSize int    `mapstructure:"pool_size" yaml:"pool_size"`
}

type Config struct {
	App      *App      `json:"app" mapstructure:"app"`
	Server   *Server   `json:"server" mapstructure:"server"`
	Database *Database `json:"database" mapstructure:"database"`
	Auth     *Auth     `json:"auth" mapstructure:"auth"`
	Redis    *Redis    `json:"redis" mapstructure:"redis"`
}

var configFilePath = "./configs/config.yaml"

var defaultConfig = Config{
	App: &App{
		Env: "dev",
	},
	Server: &Server{
		Port:    8080,
		Address: "0.0.0.0",
	},
	Database: &Database{
		Dialect: "sqlite",
		Name:    "./data/nglab.db",
	},
	Auth: &Auth{
		Enabled:       true,
		JWTSecret:     "nglab",
		JWTExpireTime: 3600,
		IgnorePaths:   []string{"/api/login", "/api/register"},
	},
	Redis: &Redis{
		Host:     "127.0.0.1",
		Port:     6379,
		Password: "",
		DB:       0,
		PoolSize: 10,
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
		return "./data/nglab.db"
	}

	if d.Dialect == "sqlite" {
		return d.Name
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		d.User, d.Password, d.Host, d.Port, d.Name, d.Params,
	)
}

func (a *App) IsDev() bool {
	return a.Env == "dev"
}

func (a *App) IsProd() bool {
	return a.Env == "prod"
}

func (a *Redis) Addr() string {
	if err := validator.New().Struct(a); err != nil {
		return defaultConfig.Redis.Addr()
	}

	return fmt.Sprintf("%s:%d", a.Host, a.Port)
}
