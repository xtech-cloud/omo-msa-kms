package dao

import (
	"errors"
	"kms/model"
)

func CreateKey(_key model.Key) error {
	db, err := model.OpenDB()
	if nil != err {
		return err
	}
	defer model.CloseDB(db)

	isBlank := db.NewRecord(_key)
	if !isBlank {
		return errors.New("key is exists!")
	}
	return db.Create(&_key).Error
}

func QueryKey(_number string) (*model.Key, error) {
	db, err := model.OpenDB()
	if nil != err {
		return nil, err
	}
	defer model.CloseDB(db)

	key := model.Key{}

	err = db.Where("number = ?", _number).First(&key).Error
	return &key, err
}

func ListKey(_appname string) ([]*model.Key, error) {
	db, err := model.OpenDB()
	if nil != err {
		return nil, err
	}
	defer model.CloseDB(db)

	keys := make([]*model.Key, 0)

	err = db.Where("app_name = ?", _appname).Find(&keys).Error
	if nil != err {
		return nil, err
	}

	return keys, nil
}

func UpdateKeyProfile(_number string, _profile string) error {
	db, err := model.OpenDB()
	if nil != err {
		return err
	}
	defer model.CloseDB(db)

	err = db.Model(&model.Key{}).Where("number = ?", _number).Update("profile", _profile).Error
	return err
}

func UpdateKeyStatus(_number string, _status int) error {
	db, err := model.OpenDB()
	if nil != err {
		return err
	}
	defer model.CloseDB(db)

	err = db.Model(&model.Key{}).Where("number = ?", _number).Update("status", _status).Error
	return err
}
