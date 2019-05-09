package httpAPI

import (
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
