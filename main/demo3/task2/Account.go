package main

import "time"

type Account struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Balance   int64     `gorm:"not null;default:0"       json:"balance"` // 分
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
