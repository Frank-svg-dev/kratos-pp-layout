// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"kratos-pp-layout/internal/biz"
	"kratos-pp-layout/internal/conf"
	"kratos-pp-layout/internal/data"
	"kratos-pp-layout/internal/server"
	"kratos-pp-layout/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
