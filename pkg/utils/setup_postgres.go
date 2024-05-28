package utils

import (
	"standard/internal/config"
	"standard/pkg/constants"
	"standard/pkg/driver"
	"time"

	"github.com/jmoiron/sqlx"
)

func SetupPostgresConnection() (*sqlx.DB, error) {
	var dsn string

	switch config.GetConfig().Environment {
	case constants.EnvironmentDevelopment:
		dsn = config.GetConfig().DBPostgreDsn
	case constants.EnvironmentProduction:
		dsn = config.GetConfig().DBPostgreURL
	}

	// Setup sqlx config of postgreSQL
	config := driver.SQLXConfig{
		DriverName:     config.GetConfig().DBPostgreDriver,
		DataSourceName: dsn,
		MaxOpenConns:   100,
		MaxIdleConns:   10,
		MaxLifetime:    15 * time.Minute,
	}

	// Initialize postgreSQL connection with sqlx
	conn, err := config.InitializeSQLXDatabase()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
