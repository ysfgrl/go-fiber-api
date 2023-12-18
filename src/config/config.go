package config

import (
	"github.com/spf13/viper"
	"go-fiber-api/src/models"
	"go-fiber-api/src/utils/response"
	"time"
)

var appConf *config = nil
var App *app = nil
var Token *token = nil
var Mongo *mongo = nil
var Redis *redis = nil
var Minio *minio = nil
var Elastic *elastic = nil
var Rabbit *rabbit = nil

type app struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type token struct {
	Secret string        `mapstructure:"secret"`
	Expire time.Duration `mapstructure:"expire"`
}

type mongo struct {
	Url string `mapstructure:"url"`
	Db  string `mapstructure:"db"`
}

type redis struct {
	Url string `mapstructure:"url"`
	Db  string `mapstructure:"db"`
}

type minio struct {
	Url       string `mapstructure:"url"`
	Port      int    `mapstructure:"port"`
	AccessKey string `mapstructure:"accessKey"`
	SecretKey string `mapstructure:"secretKey"`
}

type elastic struct {
	Url string `mapstructure:"url"`
}

type rabbit struct {
	Url           string `mapstructure:"url"`
	Que           string `mapstructure:"que"`
	Expire        int    `mapstructure:"expire"`
	Exchange      string `mapstructure:"exchange"`
	ExchangeType  string `mapstructure:"exchangeType"`
	BiddingKey    string `mapstructure:"biddingKey"`
	PreFetchCount int    `mapstructure:"preFetchCount"`
	ConsumerTag   string `mapstructure:"consumerTag"`
}
type config struct {
	App     app     `mapstructure:"app"`
	Token   token   `mapstructure:"token"`
	Mongo   mongo   `mapstructure:"mongo"`
	Redis   redis   `mapstructure:"redis"`
	Minio   minio   `mapstructure:"minio"`
	Elastic elastic `mapstructure:"elastic"`
	Rabbit  rabbit  `mapstructure:"rabbit"`
}

func InitConf(name string) *models.MyError {
	appConf = &config{}
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName(name)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return response.GetError(err)
	}
	err = viper.Unmarshal(appConf)
	if err != nil {
		return response.GetError(err)
	}
	App = &appConf.App
	Token = &appConf.Token
	Mongo = &appConf.Mongo
	Redis = &appConf.Redis
	Minio = &appConf.Minio
	Elastic = &appConf.Elastic
	Rabbit = &appConf.Rabbit
	return nil
}
