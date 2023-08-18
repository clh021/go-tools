package db

import (
	"time"

	"gorm.io/gorm"
)

type Files struct {
	gorm.Model
	Name       string
	Md5sum     string
	UploadOKAt time.Time // 上传成功时间
}

func (d *DBService) AddFiles(name , md5sum string) (*gorm.DB) {
	file := Files{Name: name, Md5sum: md5sum, UploadOKAt: time.Now()}
	result := d.Db.Create(&file)
	return result
}