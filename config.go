package main

import (
	"os"

	"github.com/joho/godotenv"
)

func getConfig() (*Config, error) {

	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	config := &Config{
		AppConfig: AppConfig{
			AppPort: os.Getenv("APP_PORT"),
			AppName: os.Getenv("APP_NAME"),
		},
		DBConfig: DBConfig{
			DBHost:     os.Getenv("DB_HOST"),
			DBPort:     os.Getenv("DB_PORT"),
			DBName:     os.Getenv("DB_NAME"),
			DBUsername: os.Getenv("DB_USERNAME"),
			DBPassword: os.Getenv("DB_PASSWORD"),
		},
		RedisConfig: RedisConfig{
			RedisHost:     os.Getenv("REDIS_HOST"),
			RedisPort:     os.Getenv("REDIS_PORT"),
			RedisUsername: os.Getenv("REDIS_USERNAME"),
			RedisPassword: os.Getenv("REDIS_USERNAME"),
		},
	}

	return config, err
}

type Config struct {
	AppConfig   AppConfig   `yaml:"app"`
	DBConfig    DBConfig    `yaml:"db"`
	RedisConfig RedisConfig `yaml:"cache"`
}

type AppConfig struct {
	AppPort string `yaml:"port"`
	AppName string `yaml:"name"`
}

type DBConfig struct {
	DBHost       string `yaml:"host"`
	DBPort       string `yaml:"port"`
	DBName       string `yaml:"name"`
	DBUsername   string `yaml:"username"`
	DBPassword   string `yaml:"password"`
	RunMigration bool   `yaml:"run_migration"`
}

type RedisConfig struct {
	RedisHost     string `yaml:"host"`
	RedisPort     string `yaml:"port"`
	RedisUsername string `yaml:"username"`
	RedisPassword string `yaml:"password"`
}
