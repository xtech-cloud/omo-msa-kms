package dao

import (
	"errors"
	"kms/model"

	kms "github.com/xtech-cloud/omo-mod-kms"
)

func CreateApp(_appname string) error {
	appkey, appsecret, publicKey, privateKey, err := kms.CreateApp(_appname)
	if nil != err {
		return err
	}

	db, err := model.OpenDB()
	if nil != err {
		return err
	}
	defer model.CloseDB(db)

	application := model.Application{
		AppName:    _appname,
		AppKey:     appkey,
		AppSecret:  appsecret,
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}

	isBlank := db.NewRecord(application)
	if !isBlank {
		return errors.New("application is exists!")
	}
	return db.Create(&application).Error
}

func QueryApp(_appname string) (*model.Application, error) {
	db, err := model.OpenDB()
	if nil != err {
		return nil, err
	}
	defer model.CloseDB(db)

	application := model.Application{}

	err = db.Where("app_name = ?", _appname).First(&application).Error
	if nil != err {
		return nil, err
	}

	return &application, nil
}

func FindApp(_appkey string, _appsecret string) (*model.Application, error) {
	db, err := model.OpenDB()
	if nil != err {
		return nil, err
	}
	defer model.CloseDB(db)

	application := model.Application{}

	err = db.Where("app_key= ? AND app_secret = ?", _appkey, _appsecret).First(&application).Error
	if nil != err {
		return nil, err
	}

	return &application, nil
}

func ListApp() ([]*model.Application, error) {
	db, err := model.OpenDB()
	if nil != err {
		return nil, err
	}
	defer model.CloseDB(db)

	applications := make([]*model.Application, 0)

	err = db.Find(&applications).Error
	if nil != err {
		return nil, err
	}

	return applications, nil
}

func UpdateAppProfile(_appname string, _profile string) error {
	db, err := model.OpenDB()
	if nil != err {
		return err
	}
	defer model.CloseDB(db)

	err = db.Model(&model.Application{}).Where("app_name= ?", _appname).Update("profile", _profile).Error
	return err
}

func UpdateAppSecurity(_appname string, _appKey string, _appSecret string, _privateKey string, _publicKey string) error {
	db, err := model.OpenDB()
	if nil != err {
		return err
	}
	defer model.CloseDB(db)

	err = db.Model(&model.Application{}).Where("app_name= ?", _appname).Updates(model.Application{
		AppKey: _appKey, AppSecret: _appSecret, PublicKey: _publicKey, PrivateKey: _privateKey}).Error
	return err
}
