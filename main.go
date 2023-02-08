package main

import (
	"flag"
	"fmt"
	"os"
	"test/service/bindataWeb"
	"test/service/caddyModule"
	"test/service/chan1"
	"test/service/dmidecode"
	"test/service/fileModule"
	"test/service/ginExample"
	"test/service/goAdmin"
	"test/service/grpcGo/grpcGoClient"
	"test/service/grpcGo/grpcGoServer"
	"test/service/grpcGo2"
	grpcGo2Client "test/service/grpcGo2/client"
	"test/service/grpcVue"
	"test/service/grpcWS"
	grpcWSClient "test/service/grpcWS/client"
	"test/service/grpcWeb"
	"test/service/plantumlModule"
	"test/service/raccoon"
	"test/service/upload"
	"test/service/uploadAdvanced"
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
		"bindataWeb":     bindataWeb.Main,
		"ginExample":     ginExample.Main,
		"caddyModule":    caddyModule.Main,
		"fileModule":     fileModule.Main,
		"plantumlModule": plantumlModule.Main,
		"goAdmin":        goAdmin.Main,
		"websocketGin":   websocketGin.Main,
		"websocketGo":    websocketGo.Main,
		"grpcGoClient":   grpcGoClient.Main,
		"grpcGoServer":   grpcGoServer.Main,
		"grpcGo2":        grpcGo2.Main,
		"grpcGo2Client":  grpcGo2Client.Main,
		"grpcWS":         grpcWS.Main,
		"grpcWSClient":   grpcWSClient.Main,
		"grpcWeb":        grpcWeb.Main,
		"grpcVue":        grpcVue.Main,
		"upload":         upload.Main,
		"uploadAdvanced": uploadAdvanced.Main,
		"raccoon":        raccoon.Main,
		"dmidecode":      dmidecode.Main,
		"chan1":          chan1.Main,
	}
}

func parseFlags() ([]string, string) {
	var output string
	var args []string
	var notargs []string
	var in_flags bool = false
	for i := 0; i < len(os.Args); i++ {
		if os.Args[i][0] == '-' {
			in_flags = true
		}
		if i == 0 || in_flags {
			notargs = append(notargs, os.Args[i])
		} else {
			args = append(args, os.Args[i])
		}
	}
	os.Args = notargs
	flag.StringVar(&output, "o", "", "Writes output to the file specified")
	flag.Parse()
	return args, output
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
	if len(cmd) > 0 {
		if reexec.Init(cmd) {
			return
		}
	}
	args, output := parseFlags()
	fmt.Println("args ", args)
	fmt.Println("Flag -o : ", output)
}
