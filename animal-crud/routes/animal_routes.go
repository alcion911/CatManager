package routes

import (
	"animal-crud/db"
	"animal-crud/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CatImage struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

func GetCatImages(c *gin.Context) {
	resp, err := http.Get("https://api.thecatapi.com/v1/images/search?limit=10")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching cat images"})
		return
	}
	defer resp.Body.Close()

	var catImages []CatImage
	if err := json.NewDecoder(resp.Body).Decode(&catImages); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding cat images"})
		return
	}

	c.JSON(http.StatusOK, catImages)
}

func GetAnimals(c *gin.Context) {
	rows, err := db.DB.Query("SELECT id, name, image FROM animals")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var animals []models.Animal
	for rows.Next() {
		var animal models.Animal
		if err := rows.Scan(&animal.ID, &animal.Name, &animal.Image); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		animals = append(animals, animal)
	}
	c.JSON(http.StatusOK, animals)
}

func CreateAnimal(c *gin.Context) {
	var animal models.Animal
	if err := c.ShouldBindJSON(&animal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.DB.Exec("INSERT INTO animals (name, image) VALUES ($1, $2)", animal.Name, animal.Image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, animal)
}

func UpdateAnimal(c *gin.Context) {
	id := c.Param("id")
	var animal models.Animal
	if err := c.ShouldBindJSON(&animal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.DB.Exec("UPDATE animals SET name = $1, image = $2 WHERE id = $3", animal.Name, animal.Image, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, animal)
}

func DeleteAnimal(c *gin.Context) {
	id := c.Param("id")
	_, err := db.DB.Exec("DELETE FROM animals WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
