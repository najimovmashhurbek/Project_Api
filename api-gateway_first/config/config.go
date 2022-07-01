package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	UserServiceHost string
	UserServicePort int

	// context timeout in seconds
	CtxTimeout int
	RedisHost string
	RedisPort int
	LogLevel string
	HTTPPort string

	SigninKey string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8090"))
	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "127.0.0.1"))
	c.UserServicePort = cast.ToInt(getOrReturnDefault("USER_SERVICE_PORT", 9000))
	c.SigninKey = cast.ToString(getOrReturnDefault("signin_Key", "ecytvcmjykingtppsouzdehlwsigjzhewjyjbwuwszsmbzyfazzfqrqccbhlqlnbcqtfnhjiltqtxwyxqjogjlutcyfgicuibvdcjesjxyrxrrugcltaznwpfwhwrbfwecxzmnzqohjffjdznokhxtqnmfgqpjrs"))
	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))
	c.RedisHost = cast.ToString(getOrReturnDefault("REDIS_HOST", "localhost"))
	c.RedisPort = cast.ToInt(getOrReturnDefault("REDIS_PORT", 6379))
	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
