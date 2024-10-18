package models

import (
	"time"

	// For working with PostGIS types
	"gorm.io/gorm"
)

type GeoJSONFeature struct {
	Type       string                 `json:"type"` // Should be "Feature"
	Geometry   Geometry               `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
}

// Geometry represents the geometry of a GeoJSON Feature
type Geometry struct {
	Type        string      `json:"type"`        // e.g., "Point", "LineString", "Polygon"
	Coordinates interface{} `json:"coordinates"` // Can be an array of coordinates
}

// GeoData model to store geospatial data
type GeoData struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"` // Foreign key to the User table
	Name      string `gorm:"not null"`
	Geometry  string `gorm:"type:geometry"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CreateGeoData saves a new GeoData entry in the database
func CreateGeoData(db *gorm.DB, geoData *GeoData) error {
	return db.Create(geoData).Error
}

// GetGeoData retrieves all geospatial data for a specific user
func GetGeoData(db *gorm.DB, userID uint) ([]GeoData, error) {
	var geoData []GeoData
	err := db.Where("user_id = ?", userID).Find(&geoData).Error
	return geoData, err
}

// UpdateGeoData updates an existing geospatial data entry
func UpdateGeoData(db *gorm.DB, geoData *GeoData) error {
	return db.Save(geoData).Error
}
