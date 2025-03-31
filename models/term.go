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
	URL         string `gorm:"column:url"`

	// Foreign Key Fields
	SourceID uint `gorm:"not null"`
	DomainID uint `gorm:"not null"`

	// Associations
	Source Source `gorm:"foreignKey:SourceID;column:source_id"`
	Domain Domain `gorm:"foreignKey:DomainID;column:domain_id"`
}

func FindTermsByWord(db *gorm.DB, word string, includeDesc bool, includeFr bool, includeDe bool) ([]Term, error) {
	var terms []Term
	// Preload the Domain association to load Domain details with the Term
	query := db.Preload("Domain").
		Preload("Source").
		Where("english LIKE ? OR english LIKE ? OR arabic LIKE ? OR arabic LIKE ?",
			"% "+word+"%", word+"%", "% "+word+"%", word+"%").
		Order("LENGTH(english)")

	if includeFr {
		query = query.Where("french IS NOT NULL AND french != ''")
	}
	if includeDe {
		query = query.Where("german IS NOT NULL AND german != ''")
	}
	if includeDesc {
		query = query.Where("description IS NOT NULL AND description != ''")
	}
	err := query.Find(&terms).Error
	if err != nil {
		return nil, err
	}

	log.Printf("Found %d terms in the database", len(terms))
	return terms, nil
}
