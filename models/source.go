package models

import (
	"log"

	"gorm.io/gorm"
)

type Source struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"unique;column:name"`
	Description string `gorm:"column:description"`
	NameAr      string `gorm:"unique;column:name_ar"`
	URL         string `gorm:"unique;column:url"`
}

// GetAllSources retrieves all sources from the database
func GetAllSources(db *gorm.DB) ([]Source, error) {
	var sources []Source
	if err := db.Find(&sources).Error; err != nil {
		return nil, err
	}
	log.Printf("Found %d sources in the database", len(sources))
	return sources, nil
}
