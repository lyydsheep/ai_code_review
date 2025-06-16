package config

import (
	"bytes"
	"embed"
	"github.com/spf13/viper"
	"os"
)

//go:embed *.yaml
var configs embed.FS

func InitConfig() {
	env := os.Getenv("env")
	if env == "" {
		env = "dev"
	}
	vp := viper.New()
	configStream, err := configs.ReadFile("application." + env + ".yaml")
	if err != nil {
		panic(err)
	}
	vp.SetConfigType("yaml")
	if err = vp.ReadConfig(bytes.NewReader(configStream)); err != nil {
		panic(err)
	}
	if err = vp.UnmarshalKey("app", &App); err != nil {
		panic(err)
	}
	if err = vp.UnmarshalKey("database", &DB); err != nil {
		panic(err)
	}
	if err = vp.UnmarshalKey("redis", &Redis); err != nil {
		panic(err)
	}
	if err = vp.UnmarshalKey("kafka", &Kafka); err != nil {
		panic(err)
	}
	if err = vp.UnmarshalKey("email", &Email); err != nil {
		panic(err)
	}
}
