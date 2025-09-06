package database

import (
	"demo4/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func MustOpen(cfg config.Config) *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)

	if cfg.MySQLDSN != "" {
		log.Println("[DB] using MySQL:", cfg.MySQLDSN)
		db, err = gorm.Open(mysql.Open(cfg.MySQLDSN), &gorm.Config{})
	} else {
		log.Println("[DB] using SQLite:", cfg.SQLitePath)
		db, err = gorm.Open(sqlite.Open(cfg.SQLitePath), &gorm.Config{})
	}
	if err != nil {
		log.Fatal("open db:", err)
	}
	return db
}
