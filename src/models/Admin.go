package models

type Admin struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"mail" gorm:"not null; unique"`
	Password string `json:"password" gorm:"not null"`
}
