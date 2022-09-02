package pkg

import (
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"kratos-pp-layout/internal/conf"
)

var (
	nacosIp       string       //nacosip地址
	nacosPort     uint64       //nacos端口
	nacosDataId   string       //nacos配置的DataID
	nacosGroup    string       //nacos配置的group分组
)

func NacosInit(c *conf.Data) vo.NacosClientParam {
	nacosIp = c.Nacos.Ip
	nacosPort = c.Nacos.Port
	nacosDataId = c.Nacos.DataId
	nacosGroup = c.Nacos.Group
	//连接注册中心配置
	return vo.NacosClientParam{
		//客户端配置
		ClientConfig: &constant.ClientConfig{
			NamespaceId:         c.Nacos.NamespaceId,
			TimeoutMs:           c.Nacos.TimeoutMs,
			LogDir:              c.Nacos.LogDir,
		},
		//服务端配置
		ServerConfigs: []constant.ServerConfig{
			{IpAddr: nacosIp, Port: nacosPort},
		},
	}
}