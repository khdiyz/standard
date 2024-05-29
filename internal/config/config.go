package config

import (
	"os"
	"standard/pkg/logger"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

var (
	instance *Config
	once     sync.Once
)

type Config struct {
	Port        int
	Environment string
	Debug       bool

	DBPostgreDriver string
	DBPostgreDsn    string
	DBPostgreURL    string

	JWTSecret  string
	JWTExpired int
	JWTIssuer  string

	HashKey string
}

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{
			Port:        cast.ToInt(getOrReturnDefault("PORT", "8000")),
			Environment: cast.ToString(getOrReturnDefault("ENVIRONMENT", "")),
			Debug:       cast.ToBool(getOrReturnDefault("DEBUG", "")),

			DBPostgreDriver: cast.ToString(getOrReturnDefault("DB_POSTGRE_DRIVER", "")),
			DBPostgreDsn:    cast.ToString(getOrReturnDefault("DB_POSTGRE_DSN", "")),
			DBPostgreURL:    cast.ToString(getOrReturnDefault("DB_POSTGRE_URL", "")),

			JWTSecret:  cast.ToString(getOrReturnDefault("JWT_SECRET", "")),
			JWTExpired: cast.ToInt(getOrReturnDefault("JWT_EXPIRED", "")),
			JWTIssuer:  cast.ToString(getOrReturnDefault("JWT_ISSUER", "")),

			HashKey: cast.ToString(getOrReturnDefault("HASH_KEY", "")),
		}
	})
	return instance
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	err := godotenv.Load("internal/config/.env")
	if err != nil {
		logger.GetLogger().Error(err)
	}
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
