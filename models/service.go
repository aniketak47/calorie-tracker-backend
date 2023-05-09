package models

import "gorm.io/gorm"

func InitEntryRepo(DB *gorm.DB) IEntry {
	return &entryRepo{
		DB: DB,
	}
}
