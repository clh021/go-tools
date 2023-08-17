package web

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// POST /upload
func (w *webService) upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	// 生成文件名
	filename := filepath.Base(file.Filename)
	fileSavePath := filepath.Join(w.conf.UploadPath, filename)

	// 保存文件到服务器
	if err := c.SaveUploadedFile(file, fileSavePath); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	// Upload the zip file
	// objectName := "golden-oldies.zip"
	// filePath := "/tmp/golden-oldies.zip"
	// contentType := "application/zip"
	// w.minio.Put(fileSavePath, md5sumName)

	c.String(http.StatusOK, "File uploaded successfully.")
}
