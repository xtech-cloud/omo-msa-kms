package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func renderBadError(_context *gin.Context, _err error) {
	_context.JSON(http.StatusBadRequest, gin.H{
		"code":    http.StatusBadRequest,
		"message": _err.Error(),
		"data":    "",
	})
}

func renderModuleError(_context *gin.Context, _err error) {
	_context.JSON(http.StatusOK, gin.H{
		"code":    -1,
		"message": _err.Error(),
		"data":    "",
	})
}

func renderOK(_context *gin.Context, _data gin.H) {
	_context.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "",
		"data":    _data,
	})
}
