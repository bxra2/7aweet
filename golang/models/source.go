package models

import (
	"log"

	"gorm.io/gorm"
)

type Source struct {
	ID           uint       `gorm:"primaryKey"`
	Name         string     `gorm:"column:name"`
	Description  string     `gorm:"column:description"`
	NameAr       string     `gorm:"column:name_ar"`
	URL          string     `gorm:"column:url"`
	CollectionID uint       `gorm:"column:collection_id;default:0"`
	Collection   Collection `gorm:"foreignKey:CollectionID;column:collection_id"`
}

type SourceCount struct {
	SourceID uint
	Cnt      int
	Name     string
	NameAr   string
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
