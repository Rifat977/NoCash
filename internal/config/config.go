package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	Env          string
	TemplatesDir string
	StaticDir    string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Port:         os.Getenv("PORT"),
		Env:          os.Getenv("ENV"),
		TemplatesDir: os.Getenv("TEMPLATES_DIR"),
		StaticDir:    os.Getenv("STATIC_DIR"),
	}
}

func (c *Config) GetPort() string {
	return ":" + c.Port
}

func (c *Config) GetEnv() string {
	return c.Env
}

func (c *Config) GetTemplatesDir() string {
	return c.TemplatesDir
}

func (c *Config) GetStaticDir() string {
	return c.StaticDir
}
