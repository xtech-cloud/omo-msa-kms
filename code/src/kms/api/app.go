package api

import (
	"github.com/gin-gonic/gin"
	kms "github.com/xtech-cloud/omo-mod-kms"
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

	appkey, appsecret, publicKey, privateKey, err := kms.CreateApp(req.AppName)
	if nil != err {
		renderModuleError(_context, err)
		return
	}

	renderOK(_context, gin.H{
		"appkey":     appkey,
		"appsecret":  appsecret,
		"publickey":  publicKey,
		"privatekey": privateKey,
	})
}
