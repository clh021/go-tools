package db

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBService struct {
	Db *gorm.DB
}

func New(filename string) *DBService {
  d, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	dm := &DBService{
		Db: d,
	}
	dm.migrate()
	return dm
}

func (d *DBService) migrate() {
  d.Db.AutoMigrate(&Files{})
  d.Db.AutoMigrate(&Analysis{})
}

func (d *DBService) RemoveDB() {
	os.Remove("./foo.db")
}
