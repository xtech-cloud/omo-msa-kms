package main

import (
	"kms/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/kms/app/create", api.HandleCreateApp)
	router.POST("/kms/license/make", api.HandleMakeLicense)
	router.POST("/kms/license/verify", api.HandleVerifyLicense)
	router.Run(":80")
}
