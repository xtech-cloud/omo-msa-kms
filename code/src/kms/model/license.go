package model

import "github.com/jinzhu/gorm"

type License struct {
	GModel  gorm.Model `gorm:"embedded"`
	UID     string     `gorm:"column:uid;type:char(32);unique;not null"`
	Number  string     `gorm:"column:number;type:char(32);not null"`
	Code    string     `gorm:"column:code;type:varchar(64);not null"`
	License string     `gorm:"column:license;type:TEXT;not null"`
}

func (License) TableName() string {
	return "kms_license"
}
