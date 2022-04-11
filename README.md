# dolphin-go-restful-api
元豚科技DevOPS-自己的Restful API，基于Golang开发，实现创建任务、读取状态、获取结果交互，满足工具跟平台API独立分离的思想和诉求。

## Idea
做这个API服务主要的想法是不想老做重复的工作，这个API跟现在业务需求比较贴近，过往的工作里经常都要做封装，比如，下游有一堆扫描器、安全工具、质量工具，如果需要对这些工具做统一入口整合，这个时候RestFul工具API就很重要了，这个服务主要做的就是这个，上游下发任务，这里接了任务创建之后，底层Python3的脚本去处理任务，处理完把状态更新，前台API直接跟这个服务交互，知道状态之后，直接拉取结果，自行处理，实现入口全部统一。

## 实现功能
1. 创建任务，根据TypeID区分任务的类型，自己可以改，想改成啥改成啥，上游自己记住就行。
2. 获取任务状态，（0:待处理,1:处理中,2:处理完成,3:处理失败），当然，你自己想改成啥改成啥，我是为了我自己方便。
3. ENV获取信息，默认服务部署到K8S里，所以需要支持configmap，这里已经集成了，写个dockerfile就能怼进去。

## 目录清单

```
.
├── LICENSE
├── README.md
├── bin
│   └── run
└── src
    ├── api
    │   ├── apis
    │   │   └── tasks.go
    │   ├── conf
    │   │   └── env.go
    │   ├── database
    │   │   └── mysql.go
    │   ├── models
    │   │   ├── reports.go
    │   │   └── tasks.go
    │   └── router
    │       └── router.go
    ├── go.mod
    ├── go.sum
    └── server.go

```
## 用法

```shell
  go mod tidy
  go build -o bin/run src/server.go
  ./bin/run
```

## 交流和问题反馈

扫码加我，可以随时交流，我们也有开源的产品PaaS，也有SaaS化产品，百度搜『元豚科技』，我们做K8S云原生的服务搭建，全自动化构建这块，搭建CI/CD，还有集群服务，有需要随时联系。

<img src="./wechat.png" width="200px">
