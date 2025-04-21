package entity

type User struct {
	ID       string `gorm:"primaryKey"`
	Username string
	Password string
	Email    string
}
