package db

import (
	"time"

	"gorm.io/gorm"
)

type Analysis struct {
	gorm.Model
	Md5sum     string
	BeginAt time.Time // 上传成功时间
}