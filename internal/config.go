package config

import (
	"context"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Constants struct {
	PORT  string
	Mongo struct {
		URL    string
		DBName string
	}
}

type Config struct {
	Constants
	Database *mongo.Database
}

func New() (*Config, error) {
	config := Config{}
	constants, err := initViper()
	config.Constants = constants
	if err != nil {
		return &config, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(config.Constants.Mongo.URL))

	config.Database = client.Database(config.Constants.Mongo.DBName)
	return &config, err
}

func initViper() (Constants, error) {
	viper.SetConfigName("feather.config") // Configuration fileName without the .TOML or .YAML extension
	viper.AddConfigPath(".")              // Search the root directory for the configuration file
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {                       // Handle errors reading the config file
		return Constants{}, err
	}
	viper.WatchConfig() // Watch for changes to the configuration file and recompile
	viper.SetDefault("PORT", "8080")
	if err = viper.ReadInConfig(); err != nil {
		log.Panicf("Error reading config file, %s", err)
	}

	var constants Constants
	err = viper.Unmarshal(&constants)
	return constants, err
}
