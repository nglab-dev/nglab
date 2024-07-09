package conf

import (
	"errors"

	"github.com/spf13/viper"
)

type LogConfig struct {
	Level string `json:"level" yaml:"level" env:"LOG_LEVEL" default:"info"`
}

type DBConfig struct {
	Driver string `json:"driver" yaml:"driver" env:"DB_DRIVER" default:"sqlite"`
	DSN    string `json:"dsn" yaml:"dsn" env:"DB_DSN" default:"./db.sqlite"`
}

type ServerConfig struct {
	Port    string `json:"port" yaml:"port" env:"HTTP_PORT" default:"8080"`
	Address string `json:"address" yaml:"address" env:"HTTP_ADDRESS" default:"0.0.0.0"`
}

type Config struct {
	Log    LogConfig    `json:"log" yaml:"log"`
	DB     DBConfig     `json:"db" yaml:"db"`
	Server ServerConfig `json:"server" yaml:"server"`
}

var cfgInstance *Config

func Get() *Config {
	return cfgInstance
}

func Load(file string) (err error) {
	v := viper.New()
	v.SetConfigFile(file)
	v.AutomaticEnv()

	if err = v.ReadInConfig(); err != nil {
		return errors.New("config file read failed: " + err.Error())
	}

	if err = v.Unmarshal(&cfgInstance); err != nil {
		return errors.New("config unmarshal failed: " + err.Error())
	}

	return nil
}
