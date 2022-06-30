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

func main() {
	fmt.Printf("Build: %s\n", build)
	reexec.Register("ginExample", ginExample.Main)
	reexec.Register("caddyModule", caddyModule.Main)
	reexec.Register("fileModule", fileModule.Main)
	reexec.Register("plantumlModule", plantumlModule.Main)
	reexec.Register("goAdmin", goAdmin.Main)
	reexec.Register("websocketGin", websocketGin.Main)
	reexec.Register("websocketGo", websocketGo.Main)
	reexec.Register("grpcGoClient", grpcGoClient.Main)
	reexec.Register("grpcGoServer", grpcGoServer.Main)
	reexec.Register("grpcWS", grpcWS.Main)
	reexec.Register("grpcWSClient", grpcWSClient.Main)
	reexec.Register("grpcWeb", grpcWeb.Main)
	reexec.Register("grpcVue", grpcVue.Main)
	reexec.Register("raccoon", raccoon.Main)
	reexec.Register("dmidecode", dmidecode.Main)
	lnksutils.IsFileExist("index.html")
	cmd := os.Getenv("APPINTO")
	os.Unsetenv("APPINTO")
	if reexec.Init(cmd) {
		return
	}
	//默认 程序 ginExample
	// ginExample.Main()
}
