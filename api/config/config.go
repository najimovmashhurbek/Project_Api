package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production
	PostgresHost      string
	PostgresPort      int
	PostgresUser      string
	PostgresPassword  string
	PostgressDatabase string
	UserServiceHost string
	UserServicePort int
	CasbinConfigPath string

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
	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "dbpost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5433))
	c.PostgressDatabase = cast.ToString(getOrReturnDefault("POSGTRES_DB", "casbindb"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASWORD", "1234"))
	
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8090"))
	
	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "user_service"))
	c.UserServicePort = cast.ToInt(getOrReturnDefault("USER_SERVICE_PORT", 9000))
	
	c.SigninKey = cast.ToString(getOrReturnDefault("signin_Key", "ecytvcmjykingtppsouzdehlwsigjzhewjyjbwuwszsmbzyfazzfqrqccbhlqlnbcqtfnhjiltqtxwyxqjogjlutcyfgicuibvdcjesjxyrxrrugcltaznwpfwhwrbfwecxzmnzqohjffjdznokhxtqnmfgqpjrs"))
	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))

	c.RedisHost = cast.ToString(getOrReturnDefault("REDIS_HOST", "redisdb"))
	c.RedisPort = cast.ToInt(getOrReturnDefault("REDIS_PORT", 6378))
	
	c.CasbinConfigPath=cast.ToString(getOrReturnDefault("CASBIN_CONFIG_PATH","./config/rbac_model.conf"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
