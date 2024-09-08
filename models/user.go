package models

import "time"

type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username" binding:"required"`
	Password   string    `json:"password" binding:"required"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}
