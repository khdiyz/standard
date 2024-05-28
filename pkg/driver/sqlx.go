package driver

import (
	"fmt"
	"time"

	"standard/pkg/constants"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var fieldDatabase = logrus.Fields{
	constants.LoggerCategory: constants.LoggerCategoryDatabase,
}

// SQLXConfig holds the configuration for the database instance
type SQLXConfig struct {
	DriverName     string
	DataSourceName string
	MaxOpenConns   int
	MaxIdleConns   int
	MaxLifetime    time.Duration
}

// InitializeSQLXDatabase returns a new DBInstance
func (config *SQLXConfig) InitializeSQLXDatabase() (*sqlx.DB, error) {
	db, err := sqlx.Open(config.DriverName, config.DataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// set maximum number of open connections to database
	logrus.WithFields(fieldDatabase).Info(fmt.Sprintf("Setting maximum number of open connections to %d", config.MaxOpenConns))
	db.SetMaxOpenConns(config.MaxOpenConns)

	// set maximum number of idle connections in the pool
	logrus.WithFields(fieldDatabase).Info(fmt.Sprintf("Setting maximum number of idle connections to %d", config.MaxIdleConns))
	db.SetMaxIdleConns(config.MaxIdleConns)

	// set maximum time to wait for new connection
	logrus.WithFields(fieldDatabase).Info(fmt.Sprintf("Setting maximum lifetime for a connection to %s", config.MaxLifetime))
	db.SetConnMaxLifetime(config.MaxLifetime)

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	return db, nil
}
