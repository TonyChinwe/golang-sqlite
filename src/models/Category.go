package models

type Category struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `json:"name" gorm:"not null;unique"`
	Number uint   `json:"number" gorm:"not null;unique"`
}
