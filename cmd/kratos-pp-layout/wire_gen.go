// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-pp-layout/internal/biz"
	"kratos-pp-layout/internal/conf"
	"kratos-pp-layout/internal/data"
	"kratos-pp-layout/internal/server"
	"kratos-pp-layout/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.DataBaseInit(confData)
	dataData, cleanup, err := data.NewData(confData, logger, db)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	greeterService := service.NewGreeterService(greeterUsecase)
	grpcServer := server.NewGRPCServer(confServer, greeterService, logger)
	httpServer := server.NewHTTPServer(confServer, greeterService, logger)
	nacosClientParam := data.NewNacosConf(confData)
	registrar := data.NewRegistrar(nacosClientParam)
	app := newApp(logger, grpcServer, httpServer, registrar, confData)
	return app, func() {
		cleanup()
	}, nil
}
