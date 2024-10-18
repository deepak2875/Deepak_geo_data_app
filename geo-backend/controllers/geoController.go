package controllers

import (
	"encoding/json"
	"geo-backend/database"
	"geo-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaveShapes(c *gin.Context) {
	var shapes []models.GeoJSONFeature
	if err := c.ShouldBindJSON(&shapes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Insert shapes into the database
	for _, shape := range shapes {
		// Convert shape to JSON or whatever format you store it
		shapeData, err := json.Marshal(shape) // Ensure to import "encoding/json"
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal shape"})
			return
		}

		// Store the JSON in your database
		_, err = database.DB.Exec("INSERT INTO shapes (data) VALUES ($1)", shapeData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save shape"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Shapes saved successfully!"})
}

// Get all saved shapes
func GetShapes(c *gin.Context) {
	rows, err := database.DB.Query("SELECT data FROM shapes")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch shapes"})
		return
	}
	defer rows.Close()

	var shapes []models.GeoJSONFeature
	for rows.Next() {
		var shapeData []byte
		if err := rows.Scan(&shapeData); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan shape"})
			return
		}
		var shape models.GeoJSONFeature
		if err := json.Unmarshal(shapeData, &shape); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal shape"})
			return
		}
		shapes = append(shapes, shape)
	}

	c.JSON(http.StatusOK, shapes)
}

// Handle file uploads (GeoJSON/KML)
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}

	if err := c.SaveUploadedFile(file, "./uploads/"+file.Filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully!"})
}
