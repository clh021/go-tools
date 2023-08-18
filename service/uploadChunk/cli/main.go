package main

import (
	"fmt"
	"log"

	"gitee.com/linakesi/source-analysis-tools-ui/cli/config"
	"gitee.com/linakesi/source-analysis-tools-ui/cli/db"
	"gitee.com/linakesi/source-analysis-tools-ui/cli/drone"
	"gitee.com/linakesi/source-analysis-tools-ui/cli/minio"
	"gitee.com/linakesi/source-analysis-tools-ui/cli/web"
)

// 支持启动时显示构建日期和构建版本
// 需要通过命令 ` go build -ldflags "-X main.build=`git rev-parse HEAD`" ` 打包
var build = "not set"

func main() {
	fmt.Printf("Build: %s\n", build)
	c := config.Loading()
	d := drone.New(c.Drone)
	s := db.New("gorm.db")
	// sqlDB, err := s.Db.DB()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer sqlDB.Close()
	m := minio.New(c.Minio)
	// m.InitClient()
	log.Println("-------------------------")
	// m.PrintEndpointURL()
	log.Println("-------------------------")
	web.New(c.Web, d, m, s)
}
