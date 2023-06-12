package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	MongoUrl       string        `mapstructure:"MONGO_URL"`
	MongoDb        string        `mapstructure:"MONGO_DB"`
	RedisUrl       string        `mapstructure:"REDIS_URL"`
	MinioUrl       string        `mapstructure:"MINIO_URL"`
	MinioPort      string        `mapstructure:"MINIO_PORT"`
	MinioAccessKey string        `mapstructure:"MINIO_ACCESS_KEY"`
	MinioSecretKey string        `mapstructure:"MINIO_SECRET_KEY"`
	Port           string        `mapstructure:"PORT"`
	TokenSecret    string        `mapstructure:"TOKEN_SECRET"`
	TokenExpiresIn time.Duration `mapstructure:"TOKEN_EXPIRED_IN"`
}

func (c *Config) Init(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("dev")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&c)
	return nil
}
