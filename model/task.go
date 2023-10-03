package model

import "time"

type Task struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"createdat"`
	UpdateAt  time.Time `json:"updateat"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint      `json:"userid" gorm:"not null"`
}

type TaskResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"createdat"`
	UpdateAt  time.Time `json:"updateat"`
}
