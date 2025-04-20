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

func FindTermsByWord(
	db *gorm.DB,
	word string,
	includeDesc bool,
	includeFr bool,
	includeDe bool,
	page int,
	limit int,
) (int64, []Term, []SourceCount, []DomainCount, error) {
	var terms []Term
	var sources []SourceCount
	var domains []DomainCount
	var termCount int64

	// Default page and limit if not provided
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	// Base query for terms
	query := db.Preload("Domain").
		Preload("Source").
		Where("english LIKE ? OR english LIKE ? OR arabic LIKE ? OR arabic LIKE ?",
			"% "+word+"%", word+"%", "% "+word+"%", word+"%").
		Order("LENGTH(english)").
		Offset(offset).
		Limit(limit)

	// Conditional filters
	if includeFr {
		query = query.Where("french IS NOT NULL AND french != ''")
	}
	if includeDe {
		query = query.Where("german IS NOT NULL AND german != ''")
	}
	if includeDesc {
		query = query.Where("description IS NOT NULL AND description != ''")
	}

	// Fetch paginated terms
	err := query.Find(&terms).Error
	if err != nil {
		return 0, nil, nil, nil, err
	}

	// Fetch total term count
	termCount, err = FoundTermCount(db, word, includeDesc, includeFr, includeDe)
	if err != nil {
		return 0, nil, nil, nil, err
	}

	// Fetch source counts (independent of pagination)
	err = db.Table("terms").
		Select("terms.source_id, COUNT(*) AS cnt, sources.name AS name, sources.name_ar AS name_ar").
		Joins("JOIN sources ON sources.id = terms.source_id").
		Where("terms.english LIKE ? OR terms.english LIKE ? OR terms.arabic LIKE ? OR terms.arabic LIKE ?",
			"% "+word+"%", word+"%", "% "+word+"%", word+"%").
		Group("terms.source_id, sources.name, sources.name_ar").
		Scan(&sources).Error
	if err != nil {
		return 0, nil, nil, nil, err
	}

	// Fetch domain counts (independent of pagination)
	err = db.Table("terms").
		Select("terms.domain_id, COUNT(*) AS cnt, domains.name AS name, domains.name_ar AS name_ar").
		Joins("JOIN domains ON domains.id = terms.domain_id").
		Where("terms.english LIKE ? OR terms.english LIKE ? OR terms.arabic LIKE ? OR terms.arabic LIKE ?",
			"% "+word+"%", word+"%", "% "+word+"%", word+"%").
		Group("terms.domain_id, domains.name, domains.name_ar").
		Scan(&domains).Error
	if err != nil {
		return 0, nil, nil, nil, err
	}

	log.Printf("Found %d terms in the database", len(terms))
	return termCount, terms, sources, domains, nil
}

func FoundTermCount(db *gorm.DB, word string,
	includeDesc bool,
	includeFr bool,
	includeDe bool) (int64, error) {
	var termCount int64
	query := db.Model(&Term{}).
		Where("english LIKE ? OR english LIKE ? OR arabic LIKE ? OR arabic LIKE ?",
			"% "+word+"%", word+"%", "% "+word+"%", word+"%").
		Order("LENGTH(english)")
	// Conditional filters
	if includeFr {
		query = query.Where("french IS NOT NULL AND french != ''")
	}
	if includeDe {
		query = query.Where("german IS NOT NULL AND german != ''")
	}
	if includeDesc {
		query = query.Where("description IS NOT NULL AND description != ''")
	}

	// Fetch paginated
	err := query.Count(&termCount).Error
	if err != nil {
		return 0, err
	}
	return termCount, nil
}

func GetRandomTerms(db *gorm.DB) ([]Term, error) {
	var terms []Term

	for i := 0; i < 15; i++ {
		var term Term

		err := db.
			Raw(`
				SELECT * FROM terms
				WHERE id = (
					SELECT ABS(RANDOM()) % (SELECT MAX(id) FROM terms) + 1
				)
				LIMIT 1;
			`).Scan(&term).Error

		if err != nil {
			return nil, err
		}

		// Only append if something was actually found (some IDs might be missing)
		if term.ID != 0 {
			terms = append(terms, term)
		}
	}

	return terms, nil
}
