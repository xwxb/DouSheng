// @Author: Ciusyan 2023/2/8
package rpc

import (
	"github.com/Go-To-Byte/DouSheng/user-service/apps/relation"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"os"

	"github.com/Go-To-Byte/DouSheng/dou-kit/client"
	"github.com/Go-To-Byte/DouSheng/dou-kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou-kit/exception"

	"github.com/Go-To-Byte/DouSheng/user-service/apps/user"
)

// 用户中心 rpc 服务的 SDK

const (
	// 由自己服务的对外提供SDK，所以注册的名称是什么，就去发现什么服务
	discoverName = "user-service"
)

// UserServiceClient 用户服务的SDJ
type UserServiceClient struct {
	// 用户模块RPC服务
	userService user.ServiceClient
	// 关系服务RPC服务
	relationService relation.ServiceClient

	l logger.Logger
}

// NewUserServiceClientFromCfg 从配置文件读取注册中心配置
func NewUserServiceClientFromCfg() (*UserServiceClient, error) {
	// 注册中心配置 [从配置文件中读取]
	cfg := conf.C().Consul.Discovers[discoverName]

	// 根据注册中心的配置，获取用户中心的客户端
	clientSet, err := client.NewClientSet(cfg)

	if err != nil {
		return nil,
			exception.WithStatusMsgf("获取服务[%s]失败：%s", cfg.DiscoverName, err.Error())
	}
	return newDefault(clientSet), nil
}

// NewUserServiceClientFromEnv 从环境变量读取注册中心配置
func NewUserServiceClientFromEnv() (*UserServiceClient, error) {
	// 注册中心配置 [从环境变量文件中读取]

	cfg := conf.NewDefaultDiscover()
	cfg.SetAddr(os.Getenv("CONSUL_ADDR"))
	cfg.SetDiscoverName(os.Getenv("CONSUL_DISCOVER_NAME"))

	// 去发现 user-service 服务
	// 根据注册中心的配置，获取用户中心的客户端
	clientSet, err := client.NewClientSet(cfg)

	if err != nil {
		return nil,
			exception.WithStatusMsgf("获取服务[%s]失败：%s", cfg.DiscoverName, err.Error())
	}
	return newDefault(clientSet), nil
}

func newDefault(clientSet *client.ClientSet) *UserServiceClient {
	conn := clientSet.Conn()
	return &UserServiceClient{
		l: zap.L().Named("USER_SERVICE_RPC"),

		// User 模块
		userService: user.NewServiceClient(conn),
		// 关系 模块
		relationService: relation.NewServiceClient(conn),
	}
}

// UserService 将服务端用户模块完整的接口暴露给外界调用，也可以做精细化控制
func (c *UserServiceClient) UserService() user.ServiceClient {
	if c.userService == nil {
		c.l.Errorf("获取用户中心[Token Client]失败")
		return nil
	}

	return c.userService
}

// RelationService 将服务端关系模块完整的接口暴露给外界调用，也可以做精细化控制
func (c *UserServiceClient) RelationService() relation.ServiceClient {
	if c.relationService == nil {
		c.l.Errorf("获取用户中心[Token Client]失败")
		return nil
	}

	return c.relationService
}
