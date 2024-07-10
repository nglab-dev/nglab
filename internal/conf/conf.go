package conf

import (
	"errors"

	"github.com/spf13/viper"
)

func Load(file string) (v *viper.Viper, err error) {
	v = viper.New()
	v.SetConfigFile(file)
	v.AutomaticEnv()

	if err = v.ReadInConfig(); err != nil {
		return nil, errors.New("config file read failed: " + err.Error())
	}

	// if err = v.Unmarshal(&cfgInstance); err != nil {
	// 	return errors.New("config unmarshal failed: " + err.Error())
	// }

	return v, nil
}
