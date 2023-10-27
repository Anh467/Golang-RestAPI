package storage

import "gorm.io/gorm"

type sqlserverStorage struct {
	db *gorm.DB
}

func NewMySQLStorage(db *gorm.DB) *sqlserverStorage {
	return &sqlserverStorage{db: db}
}
