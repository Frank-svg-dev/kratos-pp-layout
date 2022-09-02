package data

import (
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"kratos-pp-layout/internal/conf"
)

func NewNacosConf(c  *conf.Data) vo.NacosClientParam {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(c.Nacos.Ip, c.Nacos.Port),
	}
	cc := &constant.ClientConfig{
		NamespaceId:         c.Nacos.NamespaceId,
		TimeoutMs:           c.Nacos.TimeoutMs,
		LogDir:              c.Nacos.LogDir,
	}

	return vo.NacosClientParam{
		ClientConfig:  cc,
		ServerConfigs: sc,
	}
}

// NewDiscovery nacos服务发现注入
func NewDiscovery(param vo.NacosClientParam) registry.Discovery {
	client, err := clients.NewNamingClient(param)
	if err != nil {
		panic(err)
	}
	return nacos.New(client)
}

// NewRegistrar 服务注册业务注入
func NewRegistrar(param vo.NacosClientParam) registry.Registrar {
	client, err := clients.NewNamingClient(param)
	if err != nil {
		panic(err)
	}
	return nacos.New(client)
}
