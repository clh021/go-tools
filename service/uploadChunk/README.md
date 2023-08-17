# source-analysis-tools-ui

### 介绍
扫描分析工具项目的前端。目前由 golang+vue 负责web的部分，cli 的部分暂定由其中的 golang 负责只是推进优先级降低。

### 软件架构
- web golang+vue
- cli golang(web中的golang)

### 开发环境
- vscode: 1.81.1
- golang: go version go1.19.3 linux/amd64
- node  : v16.18.0
- npm   : 9.7.2
- pnpm  : 8.6.12
- yarn  : 1.22.19
- OS    : kubuntu 23.10

#### 搭建开发环境
- 启动工具服务部分
```bash
cd ./deploy
docker-compose up -d
``` 
- 启动项目服务部分
```bash
go run ./cli/main.go # 后台接口服务
cd web
pnpm run dev # 前端UI
```

#### 安装教程

1.  xxxx
2.  xxxx
3.  xxxx

#### 使用说明

1.  xxxx
2.  xxxx
3.  xxxx
