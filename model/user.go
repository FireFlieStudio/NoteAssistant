package model

type User struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Name   string `gorm:"type:varchar(100);"`
	Email  string `gorm:"type:varchar(100);unique"`
	Secret string `gorm:"type:varchar(20);"`
	Timestamps
	SoftDeletes
}
