package config

import (
	"sync"

	"github.com/nglab-dev/nglab/pkg/env"
)

type ServerConfig struct {
	Host   string `json:"host" mapstructure:"host" env:"SERVER_HOST" yaml:"host" default:"0.0.0.0"`
	Port   int    `json:"port" mapstructure:"port" env:"SERVER_PORT" yaml:"port" default:"8080"`
	Prefix string `json:"prefix" mapstructure:"prefix" env:"SERVER_PREFIX" yaml:"prefix" default:"/api"`
}

type JWTConfig struct {
	Secret     string `json:"secret" mapstructure:"secret" env:"JWT_SECRET" yaml:"secret" default:"secret"`
	ExpireTime int    `json:"expire_time" mapstructure:"expire_time" env:"JWT_EXPIRE_TIME" yaml:"expire_time" default:"36000"`
}

type LogConfig struct {
	Level    string `json:"level" mapstructure:"level" env:"LOG_LEVEL" yaml:"level" default:"debug"`
	Encoding string `json:"encoding" mapstructure:"encoding" env:"LOG_ENCODING" yaml:"encoding" default:"console"`
}

type DBConfig struct {
	Driver string `json:"driver" mapstructure:"driver" env:"DB_DRIVER" yaml:"driver" default:"sqlite"`
	DSN    string `json:"dsn" mapstructure:"dsn" env:"DB_DSN" yaml:"dsn" default:"db.sqlite"`
}

type RedisConfig struct {
	Enabled  bool   `json:"enabled" mapstructure:"enabled" env:"REDIS_ENABLED" yaml:"enabled" default:"false"`
	Addr     string `json:"addr" mapstructure:"addr" env:"REDIS_ADDR" yaml:"addr" default:"localhost:6379"`
	Password string `json:"password" mapstructure:"password" env:"REDIS_PASSWORD" yaml:"password" default:""`
	DB       int    `json:"db" mapstructure:"db" env:"REDIS_DB" yaml:"db" default:"0"`
}

type Config struct {
	Server ServerConfig `json:"server" mapstructure:"server"`
	JWT    JWTConfig    `json:"jwt" mapstructure:"jwt"`
	DB     DBConfig     `json:"db" mapstructure:"db"`
	Log    LogConfig    `json:"log" mapstructure:"log"`
	Redis  RedisConfig  `json:"redis" mapstructure:"redis"`
}

var (
	once     sync.Once // create sync.Once primitive
	instance *Config   // create nil Config struct
)

// Load function to prepare config variables from .env file and return config.
func Load() *Config {
	// Configuring config one time.
	once.Do(func() {
		host := env.GetString("SERVER_HOST", "0.0.0.0")
		port, _ := env.GetInt("SERVER_PORT", 8080)
		prefix := env.GetString("SERVER_PREFIX", "/api")

		jwtSecret := env.GetString("JWT_SECRET", "secret")
		jwtExpire, _ := env.GetInt("JWT_EXPIRE_TIME", 36000)

		logLevel := env.GetString("LOG_LEVEL", "debug")
		logEncoding := env.GetString("LOG_ENCODING", "console")

		dbDriver := env.GetString("DB_DRIVER", "sqlite")
		dbDSN := env.GetString("DB_DSN", "db.sqlite")

		redisEnalbed, _ := env.GetBool("REDIS_ENABLED", false)
		redisAddr := env.GetString("REDIS_ADDR", "localhost:6379")
		redisPassword := env.GetString("REDIS_PASSWORD", "")
		redisDB, _ := env.GetInt("REDIS_DB", 0)

		instance = &Config{
			Server: ServerConfig{
				Host:   host,
				Port:   port,
				Prefix: prefix,
			},
			JWT: JWTConfig{
				Secret:     jwtSecret,
				ExpireTime: jwtExpire,
			},
			DB: DBConfig{
				Driver: dbDriver,
				DSN:    dbDSN,
			},
			Log: LogConfig{
				Level:    logLevel,
				Encoding: logEncoding,
			},
			Redis: RedisConfig{
				Enabled:  redisEnalbed,
				Addr:     redisAddr,
				Password: redisPassword,
				DB:       redisDB,
			},
		}
	})

	// Return configured config instance.
	return instance
}
