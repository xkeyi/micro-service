package main

import (
	"fmt"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"log"
	"ms-user-service/basic"
	"ms-user-service/basic/config"
	"ms-user-service/handler"
	"ms-user-service/model"
	s "ms-user-service/proto/user"
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	// 新建服务
	service := micro.NewService(
		micro.Name("ms-service-user"),
		micro.Version("latest"),
		micro.Registry(micReg),
		)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) error {
			// 初始化模型层
			model.Init()
			// 初始化handler
			handler.Init()
			return nil
		}),
		)

	// 注册服务
	s.RegisterUserHandler(service.Server(), new(handler.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
