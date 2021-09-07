package gdb

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Entry struct {
	Config MySQLConf
	*gorm.DB
	SQLDB *sql.DB
}

func NewEntry(config MySQLConf) (*Entry, error) {
	db, err := gorm.Open(mysql.Open(config.String()), config.GORMConfig)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if config.MaxLiftTime != nil {
		sqlDB.SetConnMaxLifetime(*config.MaxLiftTime)
	}
	if config.MaxOpenConns != nil {
		sqlDB.SetMaxOpenConns(*config.MaxOpenConns)
	}
	if config.MaxIdleConns != nil {
		sqlDB.SetMaxIdleConns(*config.MaxIdleConns)
	}
	return &Entry{
		config,
		db,
		sqlDB,
	}, nil
}
