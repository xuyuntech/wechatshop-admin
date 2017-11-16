A Shop for TongRenTang E-Business
================

后端部分
---
先开启本地数据库
```
$ make db    #docker 方式运行 mysql 数据库, 数据存在 db_data 里
```

cmd/main.go 是入口, 运行如下命令即可开启服务，cmd 目录下(configy.yml 是配置文件)
```
$ go run main.go 
```

pkg 包说明
====
* `api` 业务相关的 api 接口
* `config` 项目启动时的 yaml 配置文件类
* `manager` 操作类，比如 数据库、MQ 等等
* `middleware` 中间件，比如 auth 等等
* `models` 数据库表模型，以及 engine 创建实例，为了便于开发使用 `xorm` ORM 框架
* `modules` 一些工具类
  * `httplib` Request 工具
* `routes` api 路由设置
* `utils` 工具类

api/api.go 说明
====
这个文件是主要的业务逻辑配置文件，基于 `qor` 配置，

具体配置方法，可以看 `qor` 的 `qor-example` 项目，或者 `qor` 的文档

api/*.go 文件表示的是每一个模块的 api 业务处理函数文件, 相当于 `qor-example` 里的 `controllers` 文件夹