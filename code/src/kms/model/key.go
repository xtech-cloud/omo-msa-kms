package model

import "github.com/jinzhu/gorm"

type Key struct {
	GModel   gorm.Model `gorm:"embedded"`
	AppName  string     `gorm:"column:app_name;type:varchar(32);not null"`
	Number   string     `gorm:"column:number;type:char(32);unique;not null"`
	Capacity int        `gorm:"column:capacity;not null;default:1"`
	Expiry   int        `gorm:"column:expiry;not null;default:0"`
	Status   int        `gorm:"column:status;not null;default:0"`
	Storage  string     `gorm:"column:storage"`
	Profile  string     `gorm:"column:profile;type:TEXT"`
}

func (Key) TableName() string {
	return "key"
}
