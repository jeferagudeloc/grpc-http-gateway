package model

type Profile struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Type string
}

type User struct {
	ID         string `gorm:"primaryKey"`
	Name       string `gorm:"not null"`
	LastName   string `gorm:"not null"`
	Email      string `gorm:"not null"`
	Profile    string
	Status     string
	ProfileID  string
	ProfileRef Profile `gorm:"foreignKey:ProfileID"`
}
