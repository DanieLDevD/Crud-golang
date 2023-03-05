package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	ApiPort string `mapstructure:"PORT"`
	Host string `mapstructure:"HOST"`
	User string `mapstructure:"USER"`
	Pass string `mapstructure:"PASS"`
	Database string `mapstructure:"NAME"`
}

func Load() (c Config, err error) {
viper.AddConfigPath("./envs")
viper.SetConfigName("dev")
viper.SetConfigFile("env")
viper.AutomaticEnv()

err = viper.ReadInConfig()

if err != nil {
	return
}

err = viper.Unmarshal(&c)

return
}