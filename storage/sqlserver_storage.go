package storage

import "gorm.io/gorm"

type sqlserverStore struct {
	db *gorm.DB
}

func NewSQLServerStorage(db *gorm.DB) *sqlserverStore {
	return &sqlserverStore{db: db}
}
