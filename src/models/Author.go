package models

type Author struct {
	ID          uint   `gorm:"primaryKey"`
	Surname     string `json:"surname" gorm:"not null"`
	Firstname   string `json:"first-name" gorm:"not null"`
	Lastname    string `json:"last-name" gorm:"not null"`
	Nationality string `json:"nationality"`
	Email       string `json:"mail" gorm:"unique;not null"`
	Password    string `json:"password" gorm:"not null"`
	Books       []Book `gorm:"many2many:book_author;"`
}
