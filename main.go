package main

import (
	"fmt"
	"os"
	"test/service/caddyModule"
	"test/service/fileModule"
	"test/service/ginExample"
	"test/service/plantumlModule"

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
	lnksutils.IsFileExist("index.html")
	cmd := os.Getenv("NANOAPP_CMD")
	os.Unsetenv("NANOAPP_CMD")
	if reexec.Init(cmd) {
		return
	}
	//默认 程序 ginExample
	// ginExample.Main()
}
