package database

import (
	"fmt"
	"log"

	"github.com/aniketak47/go-react-calorie-tracker/models"
	"github.com/aniketak47/go-react-calorie-tracker/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db  *gorm.DB
	err error
)

func GetDB(cfg config.Config) (*gorm.DB, error) {
	mainDSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.Username, cfg.Database.Password, cfg.Database.Name, cfg.Database.Port, cfg.Database.SSLMode,
	)

	db, err = gorm.Open(postgres.Open(mainDSN), &gorm.Config{Logger: logger.Default.LogMode(4)})
	if err != nil {
		log.Fatalf("DB connection error: %s", err)
		return nil, err
	}

	err = db.AutoMigrate(models.GetMigrationModel()...)
	if err != nil {
		log.Fatalf("database migration error : %s", err)
		return nil, err
	}

	return db, nil
}
