package config

import (
	"github.com/spf13/viper"
	"os"
	"time"
)

var AppConf *config = nil

func init() {
	env := os.Getenv("ENV")
	AppConf = &config{}
	viper.SetConfigFile(env)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&AppConf.App)
	err = viper.Unmarshal(&AppConf.Token)
	err = viper.Unmarshal(&AppConf.Mongo)
	err = viper.Unmarshal(&AppConf.Redis)
	err = viper.Unmarshal(&AppConf.Minio)
	err = viper.Unmarshal(&AppConf.Elastic)
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}

type config struct {
	App struct {
		Port  int    `mapstructure:"APP_PORT"`
		Host  string `mapstructure:"APP_HOST"`
		Debug bool   `mapstructure:"APP_DEBUG"`
	}
	Token struct {
		PublicKey  string        `mapstructure:"TOKEN_PUBLIC_KEY"`
		PrivateKey string        `mapstructure:"TOKEN_PRIVATE_KEY"`
		Method     string        `mapstructure:"TOKEN_METHOD"`
		Type       string        `mapstructure:"TOKEN_TYPE"`
		Expire     time.Duration `mapstructure:"TOKEN_EXPIRE"`
	}
	Mongo struct {
		Url      string `mapstructure:"MONGO_URL"`
		Port     int    `mapstructure:"MONGO_PORT"`
		Db       string `mapstructure:"MONGO_DB"`
		UserName string `mapstructure:"MONGO_USERNAME"`
		Password string `mapstructure:"MONGO_PASSWORD"`
	}
	Redis struct {
		Url  string `mapstructure:"REDIS_URL"`
		Port int    `mapstructure:"REDIS_PORT"`
		Db   string `mapstructure:"REDIS_DB"`
	}
	Minio struct {
		Url       string `mapstructure:"MINIO_URL"`
		Port      int    `mapstructure:"MONIO_PORT"`
		AccessKey string `mapstructure:"MINIO_ACCESS_KEY"`
		SecretKey string `mapstructure:"MINIO_SECRET_KEY"`
	}
	Elastic struct {
		Url  string `mapstructure:"ELASTIC_URL"`
		Port int    `mapstructure:"ELASTIC_PORT"`
	}
}
