package main

import (
	"dubbo-go-demo/go-client/server_init"
)

// export DUBBO_GO_CONFIG_PATH=$PATH_TO_APP/conf/dubbogo.yaml
func main() {
	server_init.ServerInit().RegisterConsumer()
	server_init.ServerInit().InitSignal()
}
