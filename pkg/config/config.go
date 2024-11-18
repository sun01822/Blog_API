package config

import (
	"github.com/kraken-io/kraken-go"
	"github.com/spf13/viper"
	"log"
)

// Config is a struct that holds the configuration of the application
type Config struct {
	DBUser          string `mapstructure:"DBUSER"`
	DBPass          string `mapstructure:"DBPASS"`
	DBIp            string `mapstructure:"DBIP"`
	DBName          string `mapstructure:"DBNAME"`
	Port            string `mapstructure:"PORT"`
	JWTSecret       string `mapstructure:"JWTSECRET"`
	KrakenAPIKey    string `mapstructure:"KRAKENAPIKEY"`
	KrakenAPISecret string `mapstructure:"KRAKENAPISECRET"`
	AppKey          string `mapstructure:"APPKEY"`
}

// Global var to access from any package
var LocalConfig *Config

// InitConfig is a function that initializes the configuration of the application
func InitConfig() *Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	// automatically reads the config vars
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file", err)
	}
	var config *Config
	// converts the read config vars into mapped struct type
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Unable to decode into struct", err)
	}
	return config

}

// Calling the InitConfig function to initialize the configuration
func SetConfig() {
	LocalConfig = InitConfig()
}

func Kraken() (*kraken.Kraken, error) {
	conf := LocalConfig
	return kraken.New(conf.KrakenAPIKey, conf.KrakenAPISecret)
}
