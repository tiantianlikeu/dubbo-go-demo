package server_init

import (
	"fmt"
	"os"
	"os/signal"
	"dubbo-go-demo/go-server/impl"
	"syscall"
	"time"
)

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	hessian "github.com/apache/dubbo-go-hessian2"
	"github.com/dubbogo/gost/log/logger"
	"dubbo-go-demo/entity"
)

var (
	survivalTimeout = int(3e9)
)

type Init struct{}

func ServerInit() *Init {
	return &Init{}
}

func (this *Init) RegisterProvider() {

	// ------for hessian2------
	hessian.RegisterPOJO(entity.NewUser())
	// ------  服务注册
	config.SetProviderService(new(impl.UserApiImpl))

	if err := config.Load(); err != nil {
		panic(err)
	}
}

func (this *Init) InitSignal() {
	signals := make(chan os.Signal, 1)
	// It is not possible to block SIGKILL or syscall.SIGSTOP
	signal.Notify(signals, os.Interrupt, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-signals
		logger.Infof("get signal %s", sig.String())
		switch sig {
		case syscall.SIGHUP:
			// reload()
		default:
			time.AfterFunc(time.Duration(survivalTimeout), func() {
				logger.Warnf("app exit now by force...")
				os.Exit(1)
			})

			// The program exits normally or timeout forcibly exits.
			fmt.Println("provider app exit now...")
			return
		}
	}
}
