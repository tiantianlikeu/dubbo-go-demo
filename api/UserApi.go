package api

import (
	"context"
	"dubbo-go-demo/entity"
)

type UserApi struct {
	GetUser func(ctx context.Context, req int32) (*entity.User, error) `dubbo:"getUser"`
}

// MethodMapper 定义方法名映射，从 Go 的方法名映射到 Java 小写方法名，只有 dubbo 协议服务接口才需要使用
func (s *UserApi) MethodMapper() map[string]string {
	return map[string]string{
		"GetUser": "getUser",
	}
}
