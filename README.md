# mqshop 麻雀电商系统
麻雀电商系统 基于go-skeleton + Reactjs + shadcn-UI 开发的面向出海的插件化后台管理系统

# 努力开发中，敬请期待...

技术栈：
- 后端骨架：
[go-skeleton](https://github.com/MQEnergy/mqshop)
- 前端组件：
[shadcn-ui](https://github.com/shadcn-ui/ui) 
- 后台前端项目：
[mqshop-admin](https://github.com/MQEnergy/mqshop-admin) 


## 运行项目
#### 1、后端运行
1）配置configs中yaml文件（默认config.dev.yaml）数据库等信息

2）新建mysql数据库`mqshop`, 执行migrate
```shell
go run cmd/cli/main.go genMigrate
```
生成一个初始化账号密码：admin / admin888

3）运行项目：
```shell
air
# or go run cmd/app/main.go
```

浏览器打开[http://127.0.0.1:9527/ping](http://127.0.0.1:9527/ping) 以下结果表示运行成功：
```json
{
"errcode": 0,
"requestid": "8d79f2a5-1ee2-4c67-9b98-57192b5d37ed",
"message": "请求成功",
"data": "pong"
}
```

具体运行查看[go-skeleton](https://github.com/MQEnergy/go-skeleton)脚手架的使用文档


#### 2、前端（待开发...）
```shell
cd web
pnpm install
pnpm run dev
```

