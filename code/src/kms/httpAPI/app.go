package httpAPI

import (
	"kms/dao"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type CreateAppRequest struct {
	AppName string `json:"appname" binding:"required"`
}

func HandleCreateApp(_context *gin.Context) {
	var req CreateAppRequest
	err := _context.ShouldBindJSON(&req)
	if nil != err {
		renderBadError(_context, err)
		return
	}

	err = dao.CreateApp(req.AppName)
	if nil != err {
		renderModuleError(_context, err)
		return
	}

	renderOK(_context, gin.H{})
}

type QueryAppRequest struct {
	AppName string `json:"appname" binding:"required"`
}

func HandleQueryApp(_context *gin.Context) {
	var req QueryAppRequest
	err := _context.ShouldBindJSON(&req)
	if nil != err {
		renderBadError(_context, err)
		return
	}

	application, err := dao.QueryApp(req.AppName)
	if nil != err {
		renderModuleError(_context, err)
		return
	}
	renderOK(_context, gin.H{
		"AppName":    application.AppName,
		"AppKey":     application.AppKey,
		"AppSecret":  application.AppSecret,
		"PublicKey":  application.PublicKey,
		"PrivateKey": application.PrivateKey,
		"Profile":    application.Profile,
		"CreatedAt":  application.GModel.CreatedAt.String(),
	})
}

type ListAppRequest struct {
}

func HandleListApp(_context *gin.Context) {
	var req ListAppRequest
	err := _context.ShouldBindJSON(&req)
	if nil != err {
		renderBadError(_context, err)
		return
	}

	applications, err := dao.ListApp()
	if nil != err {
		renderModuleError(_context, err)
		return
	}

	names := make([]string, len(applications))
	for i := 0; i < len(applications); i++ {
		names[i] = applications[i].AppName
	}
	renderOK(_context, gin.H{
		"Applications": names,
	})
}

type ModifyAppProfileRequest struct {
	AppName string `json:"appname" binding:"required"`
	Profile string `json:"profile"`
}

func HandleModifyAppProfile(_context *gin.Context) {
	var req ModifyAppProfileRequest
	err := _context.ShouldBindJSON(&req)
	if nil != err {
		renderBadError(_context, err)
		return
	}

	err = dao.UpdateAppProfile(req.AppName, req.Profile)
	if nil != err {
		renderModuleError(_context, err)
		return
	}

	renderOK(_context, gin.H{})
}

type ModifyAppSecurityRequest struct {
	AppName string `json:"appname" binding:"required"`
	AppKey string `json:"appkey" binding:"required"`
	AppSecret string `json:"appsecret" binding:"required"`
	PrivateKey string `json:"privatekey" binding:"required"`
	PublicKey string `json:"publickey" binding:"required"`
}

func HandleModifyAppSecurity(_context *gin.Context) {
	var req ModifyAppSecurityRequest
	err := _context.ShouldBindJSON(&req)
	if nil != err {
		renderBadError(_context, err)
		return
	}

	err = dao.UpdateAppSecurity(req.AppName, req.AppKey, req.AppSecret, req.PrivateKey, req.PublicKey)
	if nil != err {
		renderModuleError(_context, err)
		return
	}

	renderOK(_context, gin.H{})
}
