package database

import (
	"fmt"
	"jadwaldokter/config"
	rDokter "jadwaldokter/features/dokter/repository"
	rJadwal "jadwaldokter/features/jadwal/repository"
	rPoli "jadwaldokter/features/poli/repository"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(c *config.AppConfig) *gorm.DB {
	str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUser,
		c.DBPwd,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)

	db, err := gorm.Open(mysql.Open(str), &gorm.Config{})
	if err != nil {
		log.Error("db config error :", err.Error())
		return nil
	}

	migrateDB(db)
	return db
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&rPoli.Poli{})
	db.AutoMigrate(&rDokter.Dokter{})
	db.AutoMigrate(&rJadwal.Jadwal{})
}
