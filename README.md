# dousheng

#### 项目运行指令
```azure
make
go mod tidy
make go_build
./services.sh start
```

`./services.sh start` 启动服务

`./services.sh restart` 重启服务

`./services.sh stop` 重启服务

#### 项目结构

```
├── cmd                 # api服务和多个rpc服务
│   ├── api                 # api HTTP 服务
│   │   └──biz          
│   │      └──mw            # JWT 中间件
│   ├── comment             # comment RPC 服务
│   ├── favorite            # favorite RPC 服务
│   ├── feed                # feed RPC 服务
│   ├── message             # message RPC 服务
│   ├── mykafka         
│   │   └──consumers          
│   │      ├──consumer.go   # kafka消费者
│   │      └──cronjob.go    # 定时任务
│   ├── publish             # publish RPC 服务
│   ├── relation            # relation RPC 服务
│   └── user                # user RPC 服务
├── dal                     # dal层
│   ├──db                   # MYSQL
│   └──redis                # Redis
├── idl                     # 接口定义
├── kitex_gen               # kitex自动生成
├── pack                    # 数据结构转换
├── pkg                     # dal层
│   ├──configs          
│   │      ├──redis         # redis docker配置
│   │      └──sql           # sql表、表结构定义
│   ├── consts              # 常量定义
│   └── errno               # error定义
├── docker-compose.yaml     # docker 服务部署
├── go.mod
├── go.sum
├── Makefile
├── README.md
└── services.sh             # 项目运行脚本
```
