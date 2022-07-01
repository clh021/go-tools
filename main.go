package main

import (
	"fmt"
	"os"
	"test/service/caddyModule"
	"test/service/dmidecode"
	"test/service/fileModule"
	"test/service/ginExample"
	"test/service/goAdmin"
	"test/service/grpcGo/grpcGoClient"
	"test/service/grpcGo/grpcGoServer"
	"test/service/grpcVue"
	"test/service/grpcWS"
	grpcWSClient "test/service/grpcWS/client"
	"test/service/grpcWeb"
	"test/service/plantumlModule"
	"test/service/raccoon"
	"test/service/websocketGin"
	"test/service/websocketGo"

	"github.com/linakesi/lnksutils"
	"github.com/linakesi/lnksutils/reexec"
)

// 支持启动时显示构建日期和构建版本
// 需要通过命令 ` go build -ldflags "-X main.build=`git rev-parse HEAD`" ` 打包
var build = "not set"

func regHandles() map[string]func() {
	return map[string]func(){
		"ginExample":     ginExample.Main,
		"caddyModule":    caddyModule.Main,
		"fileModule":     fileModule.Main,
		"plantumlModule": plantumlModule.Main,
		"goAdmin":        goAdmin.Main,
		"websocketGin":   websocketGin.Main,
		"websocketGo":    websocketGo.Main,
		"grpcGoClient":   grpcGoClient.Main,
		"grpcGoServer":   grpcGoServer.Main,
		"grpcWS":         grpcWS.Main,
		"grpcWSClient":   grpcWSClient.Main,
		"grpcWeb":        grpcWeb.Main,
		"grpcVue":        grpcVue.Main,
		"raccoon":        raccoon.Main,
		"dmidecode":      dmidecode.Main,
	}
}

func main() {
	fmt.Printf("Build: %s\n", build)
	handles := regHandles()
	for name, funcName := range handles {
		reexec.Register(name, funcName)
	}
	lnksutils.IsFileExist("index.html")
	cmd := os.Getenv("APPINTO")
	os.Unsetenv("APPINTO")
	if reexec.Init(cmd) {
		return
	}
	//默认 程序 ginExample
	// ginExample.Main()
}
