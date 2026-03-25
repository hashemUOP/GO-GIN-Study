package model

import "time"

type CreateUser struct {
	// gorm:"primaryKey" tells GORM user_id is the PK
	// gorm:"column:user_id" matches your SERIAL PRIMARY KEY
	UserID    uint      `gorm:"primaryKey;column:user_id" json:"user_id"`
	Username  string    `gorm:"column:username;unique;not null" json:"username"`
	Email     string    `gorm:"column:email;unique;not null" json:"email"`
	Password  string    `gorm:"column:password;not null" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	IsActive  bool      `gorm:"column:is_active;default:true" json:"is_active"`
}

// ADD THIS FUNCTION RIGHT HERE
// This forces GORM to use the "users" table, no matter what the struct is named.
func (CreateUser) TableName() string {
	return "users"
}
