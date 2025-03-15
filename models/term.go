package models

type Term struct {
	ID          uint   `gorm:"primaryKey"`
	Arabic      string `gorm:"not null"`
	English     string `gorm:"not null"`
	French      string
	German      string
	Description string
}
