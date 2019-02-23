package api

import (
	"github.com/gin-gonic/gin"
	kms "github.com/xtech-cloud/omo-mod-kms"
)

type MakeLicenseRequest struct {
	AppKey     string `json:"appkey" binding:"required"`
	AppSecret  string `json:"appsecret" binding:"required"`
	PublicKey  string `json:"publickey" binding:"required"`
	PrivateKey string `json:"privatekey" binding:"required"`
	DeviceCode string `json:"devicecode" binding:"required"`
	Expiry     int    `json:"expiry"`
	Storage    string `json:"storage"`
}

func HandleMakeLicense(_context *gin.Context) {
	var req MakeLicenseRequest
	err := _context.ShouldBindJSON(&req)
	if nil != err {
		renderBadError(_context, err)
		return
	}

	license, err := kms.MakeLicense(
		req.AppKey,
		req.AppSecret,
		req.DeviceCode,
		req.Storage,
		req.Expiry,
		req.PublicKey,
		req.PrivateKey)

	if nil != err {
		renderModuleError(_context, err)
		return
	}

	renderOK(_context, gin.H{
		"license": license,
	})
}

type VerifyLicenseRequest struct {
	License    string `json:"license" binding:"required"`
	AppKey     string `json:"appkey" binding:"required"`
	AppSecret  string `json:"appsecret" binding:"required"`
	DeviceCode string `json:"devicecode" binding:"required"`
}

func HandleVerifyLicense(_context *gin.Context) {
	var req VerifyLicenseRequest
	err := _context.ShouldBindJSON(&req)
	if nil != err {
		renderBadError(_context, err)
		return
	}

	result, err := kms.VerifyLicense(req.License, req.AppKey, req.AppSecret, req.DeviceCode)
	if nil != err {
		renderModuleError(_context, err)
		return
	}

	renderOK(_context, gin.H{
		"result": result,
	})
}
