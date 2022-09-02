package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	v1 "kratos-pp-layout/api/helloworld/v1"
	"kratos-pp-layout/internal/conf"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewNacosConf, NewDiscovery, NewRegistrar, NewData, NewGreeterRepo, DataBaseInit, NewUserServiceClient)

// Data .
type Data struct {
	// TODO wrapped database client
	Db *gorm.DB

}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		Db: db,
	}, cleanup, nil
}

func DataBaseInit(c *conf.Data) *gorm.DB {
	if db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       c.Database.Source, // DSN data source name
		DefaultStringSize:         256,        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,      // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}); err != nil {
		fmt.Println("数据源连接失败")
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(5)
		sqlDB.SetMaxOpenConns(10)
		sqlDB.SetConnMaxLifetime(time.Hour * 24)
		fmt.Println("链接成功")
		return db
	}
	return nil
}


func NewUserServiceClient(r registry.Discovery, logger log.Logger) v1.GreeterClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///kratos-pp-layout.grpc"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			logging.Client(logger), //日志中间件,
		),
	)
	if err != nil {
		panic(err)
	}
	return v1.NewGreeterClient(conn)
}


