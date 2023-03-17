package impl

import (
	"context"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/dubbogo/gost/log/logger"
	"dubbo-go-demo/api"
	_ "dubbo-go-demo/api"
	"dubbo-go-demo/entity"
	"strconv"
	"time"
)

type UserApiImpl struct {
	api.UserApi
}

func NewUserApiImpl() *UserApiImpl {
	return &UserApiImpl{}
}

func (u *UserApiImpl) GetUser(ctx context.Context, req int32) (*entity.User, error) {
	var err error
	logger.Infof("req:%#v", req)
	user := entity.NewUser()
	user.Id = strconv.Itoa(int(req))
	user.Name = "laurence"
	user.Age = 22
	user.Time = time.Now()
	return user, err
}
