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


### 看看我
```text
如果你觉得有用请给个小星星，如果你有建议请提个issue
```

#### 项目结构说明
```text
.
├── README.md
├── api // 类似于java中对外提供的interface模块
│   └── UserApi.go
├── entity // 类似于java中对外提供的model模块
│   └── User.go
├── go-client
│   ├── cmd
│   │   └── app.go // 系统启动入口，有点像springboot的application
│   ├── conf
│   │   └── dubbogo.yaml // 客户端配置文件
│   └── server_init
│       └── dubbo_init.go // 服务注册
├── go-server
│   ├── cmd
│   │   └── app.go // 系统启动入口
│   ├── conf
│   │   └── dubbogo.yaml // 服务端配置文件
│   ├── impl // 对外提供的服务实现类
│   │   └── UserApiImpl.go 
│   ├── server_init
│   │   └── dubbo_init.go // 服务初始化
│   └── service // 这里写业务逻辑模块
│       └── UserService.go 
├── go.mod
└── go.sum

```