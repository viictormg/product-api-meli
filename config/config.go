package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DbHost    string   `mapstructure:"DB_HOST"`
	DbPort    string   `mapstructure:"DB_PORT"`
	DbUser    string   `mapstructure:"DB_USER"`
	DbPass    string   `mapstructure:"DB_PASS"`
	Brokers   []string `mapstructure:"KAFKA_BROKERS"`
	Topic     string   `mapstructure:"KAFKA_TOPIC"`
	SslMode   string   `mapstructure:"SSL_MODE"`
	DbName    string   `mapstructure:"DB_NAME"`
	retry     int      `mapstructure:"KAFKA_RETRY"`
	RedisHost string   `mapstructure:"REDIS_HOST"`
	RedisPort string   `mapstructure:"REDIS_PORT"`
}

type KafkaConfig struct {
	Brokers []string
	Retry   int
	Topic   string
}

type ConfingDB struct {
	DbHost  string
	DbPort  string
	DbUser  string
	DbPass  string
	SslMode string
	DbName  string
}

type RedisConfig struct {
	Host string
	Port string
}

func NewConfig() *Config {
	config := Config{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	fmt.Println(config.DbHost)

	return &config
}

func (c *Config) GeKafkaConfg() *KafkaConfig {
	return &KafkaConfig{
		Brokers: c.Brokers,
		Retry:   c.retry,
		Topic:   c.Topic,
	}
}

func (c *Config) GetDbConfig() ConfingDB {
	return ConfingDB{
		DbHost:  c.DbHost,
		DbPort:  c.DbPort,
		DbUser:  c.DbUser,
		DbPass:  c.DbPass,
		SslMode: c.SslMode,
		DbName:  c.DbName,
	}
}

func (c *Config) GetRedisConfig() *RedisConfig {
	return &RedisConfig{
		Host: c.RedisHost,
		Port: c.RedisPort,
	}
}
