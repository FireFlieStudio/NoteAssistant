package model

import "gorm.io/gorm"

type User struct {
	ID     uint   `gorm:"primaryKey;auto_increment"`
	Name   string `gorm:"type:varchar(100);"`
	Email  string `gorm:"type:varchar(100);unique_index"`
	Secret string `gorm:"type:varchar(20);"`
	gorm.Model
}
