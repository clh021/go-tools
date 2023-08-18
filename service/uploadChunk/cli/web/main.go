package web

import (
	"log"
	"os"
	"path/filepath"

	"gitee.com/linakesi/source-analysis-tools-ui/cli/config"
	"gitee.com/linakesi/source-analysis-tools-ui/cli/db"
	"gitee.com/linakesi/source-analysis-tools-ui/cli/drone"
	"gitee.com/linakesi/source-analysis-tools-ui/cli/fs"
	"gitee.com/linakesi/source-analysis-tools-ui/cli/minio"
	"github.com/gin-gonic/gin"
	"github.com/spf13/afero"
)

func New(c config.WebConfig, d *drone.DroneService, m *minio.MinioService, s *db.DBService) *webService {
		p, err := filepath.Abs(c.UploadPath)
	if err != nil {
		log.Panic(err)
	}

	if err := os.MkdirAll(c.UploadPath, 0755); err != nil {
		log.Fatal("Failed to create upload directory")
	}

	log.Printf("Abs File System Path: %s\n", p)
	Fs := afero.NewBasePathFs(afero.NewOsFs(), p)
	FM := fs.NewFsManager(Fs)
	w := &webService{
		FM: FM,
		conf: &c,
		drone: d,
		minio: m,
		db: s,
	}
	r := gin.Default()

	r.POST("/api/upload", w.upload)
	r.POST("/api/uploadChunk", w.UploadChunk)
	r.POST("/api/uploadChunkDone", w.UploadChunkDone)
	r.POST("/api/sendToMinIO", w.SendToMinIO)

	// 设置静态文件路由，用于访问上传的文件
	r.Static("/download", c.UploadPath)

	if err := r.Run(c.ServerPort); err != nil {
		log.Fatal("Failed to start server")
	}
	return w
}

type webService struct {
	FM *fs.FsManager
	conf *config.WebConfig
	drone *drone.DroneService
	minio *minio.MinioService
	db    *db.DBService
}
