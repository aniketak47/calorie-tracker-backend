package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aniketak47/go-react-calorie-tracker/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UpdateIngredientEntry struct {
	Ingredients string `json:"ingredients" binding:"required"`
}

type UpdateEntry struct {
	Dish        string  `json:"dish"`
	Fat         float64 `json:"fat"`
	Ingredients string  `json:"ingredients"`
	Calories    string  `json:"calories"`
}

func (base *BaseController) AddEntry(c *gin.Context) {
	var (
		entry     = models.Entry{}
		entryRepo = models.InitEntryRepo(base.DB)
	)
	err := c.BindJSON(&entry)
	if err != nil {
		fmt.Printf("Binding Error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"description": "Bad format",
		})
		return
	}

	err = entryRepo.Create(&entry)
	if err != nil {
		fmt.Printf("Create User Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"description": "Server error",
		})
		return
	}
	c.JSON(http.StatusCreated, entry)

}

func (base *BaseController) GetEntries(c *gin.Context) {
	var (
		entryRepo = models.InitEntryRepo(base.DB)
	)

	entries, err := entryRepo.GetAllEntries()
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Printf("unable to find entries, something went wrong : %s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"description": "unable to find entries, something went wrong",
		})
		return
	}
	if err == gorm.ErrRecordNotFound {
		fmt.Printf("no record found : %s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"description": "no record found",
		})
		return
	}

	c.JSON(http.StatusOK, entries)
}

func (base *BaseController) EntryById(c *gin.Context) {
	var (
		entryRepo = models.InitEntryRepo(base.DB)
	)
	entryID := c.Param("id")

	if entryID == "" {
		fmt.Print("id is empty")
		c.JSON(http.StatusBadRequest, gin.H{
			"description": "Bad Format",
		})
		return
	}

	entryIDu, err := strconv.ParseUint(entryID, 10, 64)
	if err != nil {
		fmt.Printf("Unable to parse user_id: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"description": "Bad Format",
		})
		return
	}
	id := uint64(entryIDu)
	entry, err := entryRepo.GetByID(id)
	if err != nil || entry == nil {
		c.JSON(http.StatusOK, gin.H{
			"description": "no record found",
		})
		return
	}
	c.JSON(http.StatusOK, entry)
}

func (base *BaseController) GetEntriesByIngredient(c *gin.Context) {
	var (
		entryRepo = models.InitEntryRepo(base.DB)
	)
	ingredient := c.Param("ingredient")

	if ingredient == "" {
		fmt.Print("ingredient is empty")
		c.JSON(http.StatusBadRequest, gin.H{
			"description": "Bad Format",
		})
		return
	}

	entry, err := entryRepo.GetByIngredient(ingredient)
	if err != nil || entry == nil {
		c.JSON(http.StatusOK, gin.H{
			"description": "no record found",
		})
		return
	}
	c.JSON(http.StatusOK, entry)
}

func (base *BaseController) UpdateEntry(c *gin.Context) {
	var (
		request   = UpdateEntry{}
		entryRepo = models.InitEntryRepo(base.DB)
	)

	entryID := c.Param("id")
	if entryID == "" {
		fmt.Print("id is empty")
		c.JSON(http.StatusBadRequest, gin.H{
			"description": "Bad Format",
		})
		return
	}

	entryIDu, err := strconv.ParseUint(entryID, 10, 64)
	if err != nil {
		fmt.Printf("Unable to parse user_id: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"description": "Bad Format",
		})
		return
	}

	err = c.BindJSON(&request)
	if err != nil {
		fmt.Printf("Binding Error: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"description": "Bad Request",
		})
		return
	}

	id := uint64(entryIDu)

	entry, err := entryRepo.GetByID(id)
	if err != nil {
		fmt.Printf("Unable to get entry : %s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"description": "unable to get entry, something went wrong",
		})
		return
	}

	if request.Calories != "" {
		entry.Calories = request.Calories
	}
	if request.Dish != "" {
		entry.Dish = request.Dish
	}
	if request.Ingredients != "" {
		entry.Ingredients = request.Ingredients
	}
	if request.Fat != 0 {
		entry.Fat = request.Fat
	}

	err = entryRepo.Update(entry, id)
	if err != nil {
		fmt.Printf("Update Entry Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"description": "unable to update entry, something went wrong",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "updated successfully :)",
	})
}

func (base *BaseController) UpdateIngredient(c *gin.Context) {
	var (
		request   = UpdateIngredientEntry{}
		entryRepo = models.InitEntryRepo(base.DB)
	)

	entryID := c.Param("id")
	if entryID == "" {
		fmt.Print("id is empty")
		c.JSON(http.StatusBadRequest, gin.H{
			"description": "Bad Format",
		})
		return
	}

	err := c.BindJSON(&request)
	if err != nil {
		fmt.Printf("Binding Error: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"description": "Bad Request",
		})
		return
	}

	entryIDu, err := strconv.ParseUint(entryID, 10, 64)
	if err != nil {
		fmt.Printf("Unable to parse user_id: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"description": "Bad Format",
		})
		return
	}
	id := uint64(entryIDu)

	entry, err := entryRepo.GetByID(id)
	if err != nil {
		fmt.Printf("Unable to get entry : %s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"description": "unable to get entry, something went wrong",
		})
		return
	}

	entry.Ingredients = request.Ingredients

	err = entryRepo.UpdateIngredient(entry, id, request.Ingredients)
	if err != nil {
		fmt.Printf("Update Entry Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"description": "unable to update entry, something went wrong",
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message": "updated successfully :)",
	})

}

func (base *BaseController) DeleteEntry(c *gin.Context) {
	var (
		entryRepo = models.InitEntryRepo(base.DB)
	)
	entryID := c.Params.ByName("id")
	if entryID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "id is missing, please provide id",
		})
		return
	}

	id, err := strconv.ParseUint(entryID, 10, 64)
	if err != nil {
		fmt.Printf("unable to parse id %s", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "unable to parse id",
		})
		return
	}

	err = entryRepo.Delete(id)
	if err != nil {
		fmt.Printf("unable to delete %s", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "unable to delete",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully deleted :)",
	})

}
