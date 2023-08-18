package web

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	FILE_EXIST   = 1
	CHUNK_FINISH = 2
)

func (w *webService) UploadChunk(c *gin.Context) {
	chunkName := fmt.Sprintf("%s.%s.%s.%s.part", c.PostForm("md5"), c.PostForm("chunk_index"), c.PostForm("chunks"), c.PostForm("chunk_md5"))

	if w.FM.IsFileExist(c.PostForm("md5")) {
		c.JSON(200, gin.H{"code": FILE_EXIST})
		return
	}

	if w.FM.IsFileExist(chunkName) {
		c.JSON(200, gin.H{"code": CHUNK_FINISH})
		return
	}

	saveErr := w.FM.SaveFile(chunkName, func(w io.Writer) error {
		data, _ := c.FormFile("file")

		file, err := data.Open()
		if err != nil {
			return err
		}
		defer file.Close()

		hash := md5.New()

		r := io.TeeReader(file, hash)

		if _, err = io.Copy(w, r); err != nil {
			return err
		}
		if hex.EncodeToString(hash.Sum(nil)) != c.PostForm("chunk_md5") {
			fmt.Println(hex.EncodeToString(hash.Sum(nil)), c.PostForm("chunk_md5"))
			return fmt.Errorf("md5 not match")
		}
		return nil
	})
	if saveErr != nil {
		RetError(c, saveErr)
		return
	}

	c.JSON(200, gin.H{"code": CHUNK_FINISH})
}

type uploadParams struct {
}

func (w *webService) UploadChunkDone(c *gin.Context) {
	fileMD5 := c.PostForm("md5")
	fileName := c.PostForm("file_name")
	fileType := c.PostForm("type")
	fileSize := c.PostForm("size")

	var req uploadParams

	// NOTE: Beacuse the request Header, don't use BindJSON
	err := c.Bind(&req)

	if err != nil {
		RetError(c, err)
		return
	}

	if !w.FM.IsFileExist(fileMD5) {
		count, err := strconv.ParseUint(c.PostForm("count"), 10, 64)
		if err != nil {
			RetError(c, err)
			return
		}

		fileSize, err := strconv.ParseUint(c.PostForm("size"), 10, 64)
		if err != nil {
			RetError(c, err)
			return
		}

		files, err := w.FM.GetTree("/")
		if err != nil {
			RetError(c, err)
			return
		}
		matchedFiles := []string{}
		var totalSize uint64
		for _, file := range files {
			if strings.HasPrefix(file.Name(), fileMD5+".") {
				matchedFiles = append(matchedFiles, file.Name())
				totalSize += uint64(file.Size())
			}
		}
		if len(matchedFiles) != int(count) || fileSize != totalSize {
			RetError(c, fmt.Errorf("file chunks not match"))
			return
		}

		sort.Slice(matchedFiles, func(i, j int) bool {
			iIndex, _ := strconv.ParseUint(strings.Split(matchedFiles[i], ".")[1], 10, 32)
			jIndex, _ := strconv.ParseUint(strings.Split(matchedFiles[j], ".")[1], 10, 32)
			return iIndex < jIndex
		})

		createErr := w.FM.Create(fileMD5, func(iow io.Writer) error {
			for _, fileName := range matchedFiles {
				file, err := w.FM.Open(fileName)
				if err != nil {
					return err
				}

				_, err = io.Copy(iow, file)
				if err != nil {
					file.Close()
					return err
				}
			}
			return nil
		})

		if createErr != nil {
			RetError(c, createErr)
			return
		}

		go func() {
			for _, file := range matchedFiles {
				w.FM.RemoveFile(file)
			}
		}()

	}

	w.db.AddFiles(fileName, fileMD5)
	// log.Println("minIO:", filepath.Join(w.conf.UploadPath,fileMD5), fileMD5, fileType)
	// w.minio.PrintEndpointURL()
	// w.minio.Put(filepath.Join(w.conf.UploadPath,fileMD5), fileMD5, fileType)
	c.JSON(http.StatusOK, gin.H{
		"data":     "",
		"fileName": fileName,
		"fileMD5": fileMD5,
		"fileType": fileType,
		"fileSize": fileSize,
		"err":      "",
	})
}
