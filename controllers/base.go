package controllers

import (
	"github.com/aniketak47/go-react-calorie-tracker/pkg/config"
	"gorm.io/gorm"
)

type BaseController struct {
	DB     *gorm.DB
	Config config.Config
}
