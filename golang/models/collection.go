package models

import (
	"log"

	"gorm.io/gorm"
)

type Collection struct {
	ID          uint     `gorm:"primaryKey"`
	Name        string   `gorm:"column:name"`
	Description string   `gorm:"column:description"`
	NameAr      string   `gorm:"column:name_ar"`
	URL         string   `gorm:"column:url"`
	Sources     []Source `gorm:"foreignKey:CollectionID"`
}

func GetAllCollections(db *gorm.DB) ([]Collection, error) {
	var collections []Collection
	if err := db.Preload("Sources").Find(&collections).Error; err != nil {
		return nil, err
	}
	log.Printf("Found %d collections in the database", len(collections))
	return collections, nil
}
