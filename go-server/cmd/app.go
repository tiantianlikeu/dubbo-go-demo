package main

import (
	"dubbo-go-demo/go-server/server_init"
)

// export DUBBO_GO_CONFIG_PATH=$PATH_TO_APP/conf/dubbogo.yaml
func main() {
	server_init.ServerInit().RegisterProvider()
	server_init.ServerInit().InitSignal()
}
