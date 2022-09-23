package app

import (
	"github.com/woyow/example-module/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/sirupsen/logrus"

	// migrate tools
	// "github.com/golang-migrate/migrate/v4/database/postgres"
	// "github.com/golang-migrate/migrate/v4/source/file"

	"errors"
	"fmt"
	"time"
)

const (
	_defaultAttempts = 20
	_defaultTimeout  = time.Second
)

type MigrateConfig struct {
	cfgApp *config.App
	cfgPG *config.PG
	log *logrus.Logger
}

func NewMigrateConfig(cfgApp *config.App, cfgPG *config.PG, logger *logrus.Logger) *MigrateConfig {
	return &MigrateConfig{
		cfgApp: cfgApp,
		cfgPG: cfgPG,
		log: logger,
	}
}

func (c *MigrateConfig) Migrate() {

	c.log.Info("config env for migrations: ", c.cfgApp.Env)

	// Load configuration of application
	config.ReadConfig(c.cfgApp.Env)

	if c.cfgPG.PgBouncer.Enable == true {
		c.cfgPG.Host = c.cfgPG.PgBouncer.Host
		c.cfgPG.Port = c.cfgPG.PgBouncer.Port
	}

	pgProto := "postgres://"

	pgURL := pgProto + c.cfgPG.Username + ":" + c.cfgPG.Password + "@" + c.cfgPG.Host + ":" +  c.cfgPG.Port + "/" + c.cfgPG.DBName

	pgURL += fmt.Sprintf("?sslmode=%s", c.cfgPG.SSLMode)

	c.log.Debug("Migrations pg url: ", pgURL)

	var (
		attempts = _defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	for attempts > 0 {
		m, err = migrate.New("file://db/migrations", pgURL)
		if err == nil {
			break
		}

		c.log.Debugf("Migrate: postgres is trying to connect, attempts left: %d", attempts)
		time.Sleep(_defaultTimeout)
		attempts--
	}

	if err != nil {
		c.log.Fatalf("Migrate: postgres connect error: %s", err)
	}

	err = m.Up()
	defer m.Close()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		c.log.Fatalf("Migrate: up error: %s", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		c.log.Debug("Migrate: no change")
		return
	}

	c.log.Debug("Migrate: up success")
}
