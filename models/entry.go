package models

import (
	"time"

	"gorm.io/gorm"
)

type Entry struct {
	ID        uint64         `gorm:"primarykey" json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`

	Dish        string  `json:"dish"`
	Fat         float64 `json:"fat"`
	Ingredients string  `json:"ingredients"`
	Calories    string  `json:"calories"`
}

type entryRepo struct {
	DB *gorm.DB
}

func (e *entryRepo) GetByID(ID uint64) (*Entry, error) {
	return e.GetWithTx(e.DB, &Entry{ID: ID})
}

func (e *entryRepo) GetWithTx(tx *gorm.DB, where *Entry) (*Entry, error) {
	var o Entry
	err := tx.Model(&Entry{}).Where(where).First(&o).Error
	if err != nil {
		return nil, err
	}
	return &o, nil
}

func (e *entryRepo) GetByIngredient(ingredient string) (*Entry, error) {
	return e.GetWithTx(e.DB, &Entry{Ingredients: ingredient})
}

func (e *entryRepo) GetAllEntries() ([]Entry, error) {
	var u []Entry
	err := e.DB.Model(&Entry{}).Scan(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (e *entryRepo) Create(u *Entry) error {
	return e.CreateWithTx(e.DB, u)
}

func (e *entryRepo) CreateWithTx(tx *gorm.DB, u *Entry) error {
	err := tx.Create(u).Error
	if err != nil {
		return err
	}
	return nil
}

func (e *entryRepo) Delete(ID uint64) error {
	err := e.DB.Where(&Entry{ID: ID}).Delete(&Entry{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (e *entryRepo) UpdateIngredient(u *Entry, id uint64, ingredient string) error {
	return e.UpdateWithTx(e.DB, u, id, ingredient)
}

func (e *entryRepo) Update(u *Entry, id uint64) error {
	return e.UpdateWithID(e.DB, u, id)
}

func (e *entryRepo) UpdateWithTx(tx *gorm.DB, u *Entry, id uint64, ingredient string) error {
	err := tx.Model(&Entry{ID: id, Ingredients: ingredient}).Updates(u).Error
	if err != nil {
		return err
	}
	return nil
}

func (e *entryRepo) UpdateWithID(tx *gorm.DB, u *Entry, id uint64) error {
	err := tx.Model(&Entry{ID: id}).Updates(u).Error
	if err != nil {
		return err
	}
	return nil
}
