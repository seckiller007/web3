package main

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name  string `gorm:"column:name"`
	Age   int    `gorm:"column:age"`
	Grade string `gorm:"column:grade"`
}
