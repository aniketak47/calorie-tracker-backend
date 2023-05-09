package main

import (
	"log"

	"github.com/aniketak47/go-react-calorie-tracker/controllers"
	"github.com/aniketak47/go-react-calorie-tracker/models"
	"github.com/aniketak47/go-react-calorie-tracker/pkg/config"
	"github.com/aniketak47/go-react-calorie-tracker/pkg/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("config GetConfig() error: %s", err)
	}

	db, err := database.GetDB(cfg)
	if err != nil {
		log.Fatalf("database GetDB() error: %s", err)
	}

	app := config.App{
		Config: cfg,
		DB:     db,
	}

	err = db.AutoMigrate(models.GetMigrationModel()...)
	if err != nil {
		log.Fatalf("database migration error: %s", err)
	}
	models.AddSeedData(db)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())
	app.Router = router

	ctrl := controllers.BaseController{
		DB:     db,
		Config: app.Config,
	}

	router.POST("/entry/create", ctrl.AddEntry)
	router.GET("/entries", ctrl.GetEntries)
	router.GET("/entry/:id", ctrl.EntryById)
	router.GET("/ingredient/:ingredient", ctrl.GetEntriesByIngredient)

	router.PUT("/entry/update/:id", ctrl.UpdateEntry)
	router.PUT("/ingredient/update/:id", ctrl.UpdateIngredient)
	router.DELETE("/entry/delete/:id", ctrl.DeleteEntry)

	router.Run(":" + cfg.Server.Port)
}
