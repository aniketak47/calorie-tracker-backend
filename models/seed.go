package models

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func AddSeedData(db *gorm.DB) {
	e := Entry{
		ID:          1,
		Dish:        "Paneer",
		Fat:         2,
		Ingredients: "Milk",
		Calories:    "250",
	}
	err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoNothing: true,
	}).Create(&e).Error
	if err != nil {
		fmt.Printf("error is : %s", err)
		return
	}
	db.Exec(`SELECT setval('entries_id_seq', (SELECT MAX(id) from "entries"));`)

}
