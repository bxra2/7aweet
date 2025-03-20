package models

import (
	"log"

	"gorm.io/gorm"
)

type Term struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	English     string `gorm:"not null"`
	French      string `gorm:"column:french"`
	German      string `gorm:"column:german"`
	Arabic      string `gorm:"not null"`
	Description string `gorm:"column:description"`
	URL         string `gorm:"unique;column:url"`

	// Foreign Key Fields
	SourceID uint `gorm:"not null"`
	DomainID uint `gorm:"not null"`

	// Associations
	Source Source `gorm:"foreignKey:SourceID;column:source_id"`
	Domain Domain `gorm:"foreignKey:DomainID;column:domain_id"`
}

func FindTermsByWord(db *gorm.DB, word string) ([]Term, error) {
	var terms []Term
	// Use LIKE with wildcards to search for the word in the Arabic or English fields
	err := db.Where("arabic LIKE ? OR english LIKE ?",
		"%"+word+"%", "%"+word+"%").
		Find(&terms).Error
	if err != nil {
		return nil, err
	}
	log.Printf("Found %d sources in the database", len(terms))

	return terms, nil
}
