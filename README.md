# dubbo-go-demo


#### 提供go-server 支持java-client和go-client 同时调用

##### 注册中心使用zk，记得启动zk;如下，使用docker跑一个zk
```shell
docker run -d --name=zk -p 2181:2181  zookeeper:3.4.9 
```

##### 环境初始化
```shell
cd  dubbo-go-demo
go mod tidy
```

##### 运行go-server
```shell
cd dubbo-go-demo/go-server
# export DUBBO_GO_CONFIG_PATH=$PATH_TO_APP/conf/dubbogo.yaml
export DUBBO_GO_CONFIG_PATH= ~/dubbo-go-demo/go-server/conf/dubbogo.yaml

go run cmd/app.go

```

##### 运行go-client
```shell
cd dubbo-go-demo/go-client
# export DUBBO_GO_CONFIG_PATH=$PATH_TO_APP/conf/dubbogo.yaml
export DUBBO_GO_CONFIG_PATH= ~/dubbo-go-demo/go-client/conf/dubbogo.yaml

go run cmd/app.go
```

##### 运行java客户端
下载java客户端代码
```shell
git clone https://github.com/tiantianlikeu/dubbo-java-client.git
```
