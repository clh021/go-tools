package web

import (
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (w *webService) SendToMinIO(c *gin.Context) {
	fileMD5 := c.PostForm("md5")
	fileName := c.PostForm("file_name")
	fileType := c.PostForm("type")

	log.Println("minIO:", filepath.Join(w.conf.UploadPath,fileMD5), fileMD5, fileType)
	info, err := w.minio.Put(filepath.Join(w.conf.UploadPath,fileMD5), fileMD5, fileType)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded `%s`(%s) of size %d\n", fileName, fileMD5, info.Size)
	c.JSON(200, gin.H{"code": 1, "message": "update to min io"})
}
