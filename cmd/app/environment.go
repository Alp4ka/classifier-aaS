package main

import (
	"context"
	classifier_aaS "github.com/Alp4ka/classifier-aaS"
	"github.com/Alp4ka/classifier-aaS/internal/app"
	"github.com/Alp4ka/classifier-aaS/internal/config"
	dbpkg "github.com/Alp4ka/classifier-aaS/pkg/db"
	"github.com/Alp4ka/mlogger"
	"github.com/Alp4ka/mlogger/field"
	"github.com/Alp4ka/mlogger/misc"
	"github.com/jmoiron/sqlx"
	"log"
)

type environment struct {
	ctx        context.Context
	cancelFunc context.CancelFunc

	cfg *config.Config
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
	env.ctx, env.cancelFunc = context.WithCancel(context.Background())
}

func setupLogging(env *environment) {
	env.ctx = field.WithContextFields(env.ctx, field.String("appName", classifier_aaS.AppName))
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
	cfg, err := config.FromEnv()
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
	env.app = app.New(
		app.Config{
			DB:            env.db,
			OpenAIAPIKey:  env.cfg.OpenAIAPIKey,
			HTTPPort:      env.cfg.HTTPPort,
			HTTPRateLimit: env.cfg.RateLimit,
		},
	)
}
