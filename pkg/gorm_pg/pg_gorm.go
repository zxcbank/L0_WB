package gormpg

import (
	"database/sql"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/pkg/errors"
	"github.com/uptrace/bun/driver/pgdriver"
	gorm_postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

// PgConfig - конфигурация для соединения с Postgresql
type PgConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	DBName   string `mapstructure:"dbName"`
	SSLMode  bool   `mapstructure:"sslMode"`
	Password string `mapstructure:"password"`
}

// PgGorm - модель базы данных
type PgGorm struct {
	DB     *gorm.DB
	Config *PgConfig
}

func NewPgGorm(config *PgConfig) (*PgGorm, error) {
	err := createDatabaseIfNotExists(config)

	if err != nil {
		panic(err)
		return nil, err
	}

	connectionString := getConnectionString(config, config.DBName)

	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = 10 * time.Second
	maxRetries := 5

	var gormDb *gorm.DB

	err = backoff.Retry(func() error {
		gormDb, err = gorm.Open(gorm_postgres.Open(connectionString), &gorm.Config{})

		if err != nil {
			return errors.Errorf("failed to connect postgres: %v and connection information: %s", err, connectionString)
		}

		return nil
	}, backoff.WithMaxRetries(bo, uint64(maxRetries-1)))

	return &PgGorm{DB: gormDb, Config: config}, err
}

func Migrate(gorm *gorm.DB, types ...interface{}) error {

	for _, t := range types {
		err := gorm.AutoMigrate(t)
		if err != nil {
			return err
		}
	}
	return nil
}

func createDatabaseIfNotExists(config *PgConfig) error {

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		"postgres",
	)

	pgSqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connectionString)))

	var exists int

	selectDbQueryString := fmt.Sprintf("SELECT 1 FROM  pg_catalog.pg_database WHERE datname='%s'", config.DBName)

	rows, err := pgSqlDb.Query(selectDbQueryString)
	if err != nil {
		return err
	}

	if rows.Next() {
		err = rows.Scan(&exists)
		if err != nil {
			return err
		}
	}

	if exists == 1 {
		return nil
	}

	createDbQueryString := fmt.Sprintf("CREATE DATABASE %s", config.DBName)

	_, err = pgSqlDb.Exec(createDbQueryString)
	if err != nil {
		return err
	}

	defer func(pgSqlDb *sql.DB) {
		err := pgSqlDb.Close()
		if err != nil {
			panic(err)
		}
	}(pgSqlDb)

	return nil
}

func getConnectionString(config *PgConfig, dbName string) string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		config.Host,
		config.Port,
		config.User,
		dbName,
		config.Password,
	)
}
