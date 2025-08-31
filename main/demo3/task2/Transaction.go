package main

import "time"

type Transaction struct {
	ID            uint `gorm:"primaryKey"`
	FromAccountId uint
	ToAccountId   uint
	Amount        uint
	CreatedAt     time.Time `json:"created_at"`
	// 定义外键关联
	FromAccount Account `gorm:"foreignKey:FromAccountId"`
	ToAccount   Account `gorm:"foreignKey:ToAccountId"`
}
