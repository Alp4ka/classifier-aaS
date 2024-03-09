package main

import (
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/app"
	dbpkg "github.com/Alp4ka/classifier-aaS/pkg/db"
	"github.com/Alp4ka/mlogger"
	"github.com/Alp4ka/mlogger/field"
	"github.com/Alp4ka/mlogger/misc"
	"github.com/jmoiron/sqlx"
	"log"
)

const AppName = "payment"

type environment struct {
	ctx context.Context

	cfg *app.Config
	db  *sqlx.DB

	app *app.App
}

func setup() *environment {
	var env environment
	setupContext(&env)
	setupLogging(&env)
	setupConfig(&env)
	setupDatabase(&env)
	setupApp(&env)
	return &env
}

func setupContext(env *environment) {
	// Env setup.
	env.ctx = context.Background()
}

func setupLogging(env *environment) {
	logger, err := mlogger.NewProduction(
		env.ctx,
		mlogger.Config{
			Level: misc.LevelDebug,
		},
	)
	if err != nil {
		log.Fatal("could not create logger", err)
	}

	mlogger.ReplaceGlobals(logger)
}

func setupConfig(env *environment) {
	cfg, err := app.FromEnv()
	if err != nil {
		mlogger.L().Fatal("failed to load config", field.Error(err))
	}

	// Env setup.
	env.cfg = cfg
}

func setupDatabase(env *environment) {
	// DB.
	db, err := dbpkg.NewPgDB(
		dbpkg.Config{
			DSN:                   env.cfg.PgDSN,
			MaxIdleConns:          env.cfg.PgMaxIdleConnections,
			MaxOpenConns:          env.cfg.PgMaxOpenConnections,
			MaxConnectionAttempts: 10,
		},
	)
	if err != nil {
		mlogger.L().Fatal("Failed to initialize database", field.Error(err))
	}

	// Env setup.
	env.db = db
}

func setupApp(env *environment) {
	// Env setup.
	env.app = app.New(env.cfg)
}
