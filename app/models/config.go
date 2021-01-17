package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	//phpmyadmin xampp
	MySQLDsn = `root:@tcp(localhost)/prestasi?parseTime=True&charset=utf8`
)

func init() {
	var err error
	for {
		DB, err = gorm.Open(mysql.Open(MySQLDsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
		if err == nil {
			break
		}
	}
	//DB.AutoMigrate(new(Mahasiswa))
}
