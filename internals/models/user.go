package models

import (
	"time"
)

type User struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username" gorm:"uniqueIndex"` // Username (un)
	Password  string    `json:"password"`                    // Password (confirm)
	Email     string    `json:"email" gorm:"uniqueIndex"`    // Email (un)
	Gender    string    `json:"gender"`
	UserType  string    `json:"user_type"` // Type (admin/user/IngridientAdmin)
	Address   Address   `json:"address" gorm:"constraint:OnDelete:CASCADE;"`
	Phone     string    `json:"phone"`
	URL       string    `json:"url"` // URL (fb / insta)
	Response
}

type Address struct {
	Id      uint   `json:"id" gorm:"primaryKey"`
	UserId  string `json:"user_id"`
	Town    string `json:"town"`
	Country string `json:"country"`
	Road    string `json:"road"`
	Number  uint   `json:"number"`
}

type LoginResp struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
