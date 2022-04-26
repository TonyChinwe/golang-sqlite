package models

import "time"

type Book struct {
	ID            uint     `gorm:"primaryKey"`
	Title         string   `json:"title" gorm:"not null"`
	Author        []Author `gorm:"many2many:book_author;"`
	Subject       Subject
	SubjectID     int `json:"subject-id" gorm:"not null"`
	Category      Category
	CategoryID    int `json:"category-id" gorm:"not null"`
	Genre         Genre
	GenreID       int       `json:"genre-id"`
	Leveid        int       `json:"level-id"`
	Level         Level     `gorm:"foreignKey:Leveid"`
	DatePublished time.Time `json:"published-date" gorm:"not null"`
}
