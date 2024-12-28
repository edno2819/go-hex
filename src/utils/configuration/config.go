package configuration

import (
	"os"
	"strconv"
)

type DBStruct struct {
	Host     string // Database host address
	Port     string // Database port number
	User     string // Database username
	Password string // Database password
	Name     string // Database name
	SSLMode  string // SSL mode (disable, require, etc.)
	TypeDb   string // Database type
}

type AppStruct struct {
	Name        string // Application name
	Environment string // Environment (development, production, staging)
	Debug       bool   // Debug mode
	Version     string // Application version
}

type RedisStruct struct {
	Host     string // Redis host address
	Port     int    // Redis port number
	Password string // Redis password
	DB       int    // Redis database index
}

type TokenStruct struct {
	SecretKey     string // Secret key for signing tokens
	ExpirationSec int    // Token expiration time in seconds
	Issuer        string // Token issuer
}

type HttpStruct struct {
	Host         string // HTTP server host
	Port         int    // HTTP server port
	ReadTimeout  int    // Read timeout in seconds
	WriteTimeout int    // Write timeout in seconds
}

type Config struct {
	DB    DBStruct    // Database configuration
	App   AppStruct   // Application configuration
	Redis RedisStruct // Redis configuration
	Token TokenStruct // Token configuration
	HTTP  HttpStruct  // HTTP server configuration
}

func New() *Config {
	return &Config{}
}

// Load loads the configuration from environment variables
func (c *Config) Load() {
	// Load database configuration
	c.DB.Host = os.Getenv("DB_HOST")
	c.DB.Port = os.Getenv("DB_PORT")
	c.DB.User = os.Getenv("DB_USER")
	c.DB.Password = os.Getenv("DB_PASSWORD")
	c.DB.Name = os.Getenv("DB_NAME")
	c.DB.SSLMode = os.Getenv("DB_SSL_MODE")
	c.DB.TypeDb = os.Getenv("DB_TYPE")

	// Load application configuration
	c.App.Name = os.Getenv("APP_NAME")
	c.App.Environment = os.Getenv("APP_ENVIRONMENT")
	c.App.Debug, _ = strconv.ParseBool(os.Getenv("APP_DEBUG"))
	c.App.Version = os.Getenv("APP_VERSION")

	// Load Redis configuration
	c.Redis.Host = os.Getenv("REDIS_HOST")
	c.Redis.Port, _ = strconv.Atoi(os.Getenv("REDIS_PORT"))
	c.Redis.Password = os.Getenv("REDIS_PASSWORD")
	c.Redis.DB, _ = strconv.Atoi(os.Getenv("REDIS_DB"))

	// Load token configuration
	c.Token.SecretKey = os.Getenv("TOKEN_SECRET_KEY")
	c.Token.ExpirationSec, _ = strconv.Atoi(os.Getenv("TOKEN_EXPIRATION_SEC"))
	c.Token.Issuer = os.Getenv("TOKEN_ISSUER")

	// Load HTTP configuration
	c.HTTP.Host = os.Getenv("HTTP_HOST")
	c.HTTP.Port, _ = strconv.Atoi(os.Getenv("HTTP_PORT"))
	c.HTTP.ReadTimeout, _ = strconv.Atoi(os.Getenv("HTTP_READ_TIMEOUT"))
	c.HTTP.WriteTimeout, _ = strconv.Atoi(os.Getenv("HTTP_WRITE_TIMEOUT"))
}
