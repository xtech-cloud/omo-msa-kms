package dao

import (
	"errors"
	"kms/model"
	"time"
)

func CreateLicense(_license model.License) error {
	db, err := model.OpenDB()
	if nil != err {
		return err
	}
	defer model.CloseDB(db)

	isBlank := db.NewRecord(_license)
	if !isBlank {
		return errors.New("license is exists!")
	}
	return db.Create(&_license).Error
}

func QueryLicense(_uid string) (*model.License, error) {
	db, err := model.OpenDB()
	if nil != err {
		return nil, err
	}
	defer model.CloseDB(db)

	license := model.License{}

	err = db.Where("uid = ?", _uid).First(&license).Error
	return &license, err
}

func FilterLicense(_number string) ([]*model.License, error) {
	db, err := model.OpenDB()
	if nil != err {
		return nil, err
	}
	defer model.CloseDB(db)

	licenses := make([]*model.License, 0)

	err = db.Where("number = ?", _number).Find(&licenses).Error
	return licenses, err
}

func CountKey(_number string) (int, error) {
	db, err := model.OpenDB()
	if nil != err {
		return 0, err
	}
	defer model.CloseDB(db)

	count := 0

	err = db.Model(&model.License{}).Where("number = ?", _number).Count(&count).Error
	return count, err
}

func UpdateLicenseUpdatedAt(_uid string) error {
	db, err := model.OpenDB()
	if nil != err {
		return err
	}
	defer model.CloseDB(db)

	err = db.Model(&model.License{}).Where("uid = ?", _uid).Update("updated_at", time.Now()).Error
	return err
}
