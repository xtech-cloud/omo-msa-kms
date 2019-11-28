package httpAPI

import (
	"fmt"
	"errors"
	"kms/dao"

	"github.com/gin-gonic/gin"
)

type VerifyLicenseRequest struct {
	UID  string `json:"uid" binding:"required"`
	Code string `json:"code" binding:"required"`
}

func HandleVerifyLicense(_context *gin.Context) {
	var req VerifyLicenseRequest
	err := _context.ShouldBindJSON(&req)
	if nil != err {
		renderBadError(_context, err)
		return
	}

	license, err := dao.QueryLicense(req.UID)
	if nil != err {
		renderModuleError(_context, err)
		return
	}

	if license.Code != req.Code {
		renderCustomError(_context, 1, errors.New("code not matched"))
		return
	}

	key, err := dao.QueryKey(license.Number)
	if nil != err {
		renderModuleError(_context, err)
		return
	}

	renderOK(_context, gin.H{
		"status": key.Status,
	})
}

type FetchLicenseRequest struct {
	AppName string `json:"appname" binding:"required"`
	Code string `json:"code" binding:"required"`
}

func HandleFetchLicense(_context *gin.Context) {
	var req FetchLicenseRequest
	err := _context.ShouldBindJSON(&req)
	if nil != err {
		renderBadError(_context, err)
		return
	}

	licenses, err := dao.FindLicense(req.Code)
	if nil != err {
		renderModuleError(_context, err)
		return
	}

	if len(licenses) < 1 {
		renderCustomError(_context, 1, errors.New("not found"))
		return
	}

	license := licenses[0]

	key, err := dao.FindKey(req.AppName, license.Number)
	if nil != err {
		renderModuleError(_context, err)
		return
	}

	if key.Status > 0 {
		renderCustomError(_context, 2, errors.New(fmt.Sprintf("status: %v", key.Status)))
		return
	}

	renderOK(_context, gin.H{
		"license": license.License,
	})
}
