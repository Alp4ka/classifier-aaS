package main

import (
	"github.com/Alp4ka/mlogger"
	"github.com/Alp4ka/mlogger/field"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	env := setup()

	// Register GRPC API endpoints:
	// app.GRPCServer().RegisterService(&api.MyAPI_ServiceDesc, myapi.New())
	mlogger.L().Debug("Starting app")
	err := env.app.Run(env.ctx)
	mlogger.L().Debug("App finished", field.Error(err))
}
