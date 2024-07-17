package conf

import "github.com/spf13/viper"

type Config struct{}

var DefaultConfigFile = "config.yaml"

func Load(configFiles ...string) (config *Config, err error) {
	if len(configFiles) == 0 {
		configFiles = []string{DefaultConfigFile}
	}

	v := viper.New()
	for _, file := range configFiles {
		v.AddConfigPath(file)
	}

	err = v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = v.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func MustLoad(configFiles ...string) *Config {
	config, err := Load(configFiles...)
	if err != nil {
		panic(err)
	}
	return config
}
