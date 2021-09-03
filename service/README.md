# 创建服务
1. 使用 go-zero 工具
```bash
goctl api new greet
```
2. 自己编写 package.Main

3. 查看感兴趣的开源软件，发现入口(cmd)，拷贝文件=>修改包名main为npsModule=>修改main为Main



# 注册服务
增加服务注册逻辑
```golang
func main() {
    //...
	reexec.Register("goZeroGreet", goZeroGreet.Main)
    //...
}
```

# 运行服务
```bash
NANOAPP_CMD=ginExample ./bin # --help
```


# 服务加入记录
1. go-zero 不适合加入，适合作为独立项目开发
2. nps&npc 加入出错
   ```
    # ehang.io/nps/server/proxy
    /go/pkg/mod/ehang.io/nps@v0.26.10/server/proxy/tcp.go:71:3: undefined: beego.InitBeforeHTTPRun
   ```
3. file browser