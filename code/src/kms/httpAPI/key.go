package httpAPI

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"kms/dao"
	"kms/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	kms "github.com/xtech-cloud/omo-mod-kms"
)

type GenerateKeyRequest struct {
	AppKey    string `json:"appkey" binding:"required"`
	AppSecret string `json:"appsecret" binding:"required"`
	Count     int    `json:"count"`
	Capacity  int    `json:"capacity"`
	Expiry    int    `json:"expiry"`
	Storage   string `json:"storage"`
	Profile   string `json:"profile"`
}

func HandleGenerateKey(_context *gin.Context) {
	var req GenerateKeyRequest
	err := _context.ShouldBindJSON(&req)
	if nil != err {
		renderBadError(_context, err)
		return
	}

	if req.Count == 0 {
		req.Count = 1
	}

	if req.Capacity == 0 {
		req.Capacity = 1
	}

	application, err := dao.FindApp(req.AppKey, req.AppSecret)
	if nil != err {
		renderModuleError(_context, err)
		return
	}

	numbers := make([]string, 0)

	for i := 0; i < req.Count; i++ {
		number, err := newNumber()
		if nil != err {
			continue
		}
		key := model.Key{
			Number:   number,
			AppName:  application.AppName,
			Capacity: req.Capacity,
			Expiry:   req.Expiry,
			Status:   0,
			Storage:  req.Storage,
			Profile:  req.Profile,
		}
		err = dao.CreateKey(key)
		if nil != err {
			continue
		}
		numbers = append(numbers, key.Number)
	}

	renderOK(_context, gin.H{
		"keys": numbers,
	})
}

type QueryKeyRequest struct {
	Number string `json:"number" binding:"required"`
}

func HandleQueryKey(_context *gin.Context) {
	var req QueryKeyRequest
	err := _context.ShouldBindJSON(&req)
	if nil != err {
		renderBadError(_context, err)
		return
	}

	key, err := dao.QueryKey(req.Number)
	if nil != err {
		renderModuleError(_context, err)
		return
	}

	licenses, err := dao.FilterLicense(req.Number)
	if nil != err {
		renderModuleError(_context, err)
		return
	}

	codes := make([]string, len(licenses))
	for i := 0; i < len(licenses); i++ {
		codes[i] = licenses[i].Code
	}

	renderOK(_context, gin.H{
		"Number":    key.Number,
		"Capacity":  key.Capacity,
		"Expiry":    key.Expiry,
		"Storage":   key.Storage,
		"AppName":   key.AppName,
		"Status":    key.Status,
		"Profile":	 key.Profile,
		"CreatedAt": key.GModel.CreatedAt.String(),
		"Active":    codes,
	})
}

type ListKeyRequest struct {
	AppKey    string `json:"appkey" binding:"required"`
	AppSecret string `json:"appsecret" binding:"required"`
}

func HandleListKey(_context *gin.Context) {
	var req ListKeyRequest
	err := _context.ShouldBindJSON(&req)
	if nil != err {
		renderBadError(_context, err)
		return
	}

	application, err := dao.FindApp(req.AppKey, req.AppSecret)
	if nil != err {
		renderModuleError(_context, err)
		return
	}

	keys, err := dao.ListKey(application.AppName)
	if nil != err {
		renderModuleError(_context, err)
		return
	}

	numbers := make([]string, len(keys))
	for i := 0; i < len(numbers); i++ {
		numbers[i] = keys[i].Number
	}
	renderOK(_context, gin.H{
		"Keys": numbers,
	})
}

type ModifyKeyProfileRequest struct {
	Number  string `json:"number" binding:"required"`
	Profile string `json:"profile"`
}

func HandleModifyKeyProfile(_context *gin.Context) {
	var req ModifyKeyProfileRequest
	err := _context.ShouldBindJSON(&req)
	if nil != err {
		renderBadError(_context, err)
		return
	}

	err = dao.UpdateKeyProfile(req.Number, req.Profile)
	if nil != err {
		renderBadError(_context, err)
		return
	}

	renderOK(_context, gin.H{})
}

type ModifyKeyStatusRequest struct {
	Number string `json:"number" binding:"required"`
	Status int    `json:"status"`
}

func HandleModifyKeyStatus(_context *gin.Context) {
	var req ModifyKeyStatusRequest
	err := _context.ShouldBindJSON(&req)
	if nil != err {
		renderBadError(_context, err)
		return
	}

	err = dao.UpdateKeyStatus(req.Number, req.Status)
	if nil != err {
		renderModuleError(_context, err)
		return
	}

	renderOK(_context, gin.H{})
}

type ActivateKeyRequest struct {
	Number string `json:"number" binding:"required"`
	Code   string `json:"code" binding:"required"`
}

func HandleActivateKey(_context *gin.Context) {
	var req ActivateKeyRequest
	err := _context.ShouldBindJSON(&req)
	if nil != err {
		renderBadError(_context, err)
		return
	}

	key, err := dao.QueryKey(req.Number)
	//record not found
	if gorm.IsRecordNotFoundError(err) {
		renderCustomError(_context, 1, errors.New("invalid key"))
		return
	}
	//other error
	if nil != err {
		renderModuleError(_context, err)
		return
	}

	//not allow activate when status>0
	if key.Status > 0 {
		renderCustomError(_context, 2, errors.New("status > 0"))
		return
	}

	application, err := dao.QueryApp(key.AppName)
	if nil != err {
		renderModuleError(_context, err)
		return
	}

	licenseFile, err := kms.MakeLicense(
		application.AppKey,
		application.AppSecret,
		req.Code,
		key.Storage,
		key.Expiry,
		application.PublicKey,
		application.PrivateKey,
	)
	if nil != err {
		renderModuleError(_context, err)
		return
	}

	uid := toMD5(req.Number + req.Code)

	license, err := dao.QueryLicense(uid)
	if gorm.IsRecordNotFoundError(err) {
		license = &model.License{
			UID:     uid,
			Number:  req.Number,
			Code:    req.Code,
			License: licenseFile,
		}

		count, err := dao.CountKey(key.Number)
		if nil != err {
			renderModuleError(_context, err)
			return
		}

		if count >= key.Capacity {
			renderCustomError(_context, 3, errors.New("not allow at full capacity"))
			return
		}

		err = dao.CreateLicense(*license)
		if nil != err {
			renderModuleError(_context, err)
			return
		}
	} else if nil != err {
		renderModuleError(_context, err)
		return
	}

	err = dao.UpdateLicenseUpdatedAt(uid)
	if nil != err {
		renderModuleError(_context, err)
		return
	}

	renderOK(_context, gin.H{
		"uid":     license.UID,
		"license": license.License,
	})
}

func newNumber() (string, error) {
	id, err := uuid.NewV4()
	if nil != err {
		return "", err
	}
	h := md5.New()
	h.Write(id.Bytes())
	return hex.EncodeToString(h.Sum(nil)), nil
}

func toMD5(_content string) string {
	h := md5.New()
	h.Write([]byte(_content))
	return hex.EncodeToString(h.Sum(nil))
}
