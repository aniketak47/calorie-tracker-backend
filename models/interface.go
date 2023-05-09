package models

import "gorm.io/gorm"

type IEntry interface {
	GetByID(ID uint64) (*Entry, error)
	GetByIngredient(ingredient string) (*Entry, error)
	GetWithTx(tx *gorm.DB, where *Entry) (*Entry, error)
	GetAllEntries() ([]Entry, error)
	Create(u *Entry) error
	CreateWithTx(tx *gorm.DB, u *Entry) error
	Delete(ID uint64) error
	UpdateIngredient(u *Entry, id uint64, ingredient string) error
	Update(u *Entry, id uint64) error
	UpdateWithTx(tx *gorm.DB, u *Entry, id uint64, ingredient string) error
	UpdateWithID(tx *gorm.DB, u *Entry, id uint64) error
}
