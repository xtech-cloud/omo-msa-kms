package main

import (
	"kms/httpAPI"
	"kms/model"
	"net"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	model.AutoMigrateDatabase()

	httpAddrArg := os.Getenv("KMS_HTTP_ADDR")
	if "" == httpAddrArg {
		httpAddrArg = ":80"
	}

	grpcAddrArg := os.Getenv("KMS_GRPC_ADDR")
	if "" == grpcAddrArg {
		grpcAddrArg = ":10080"
	}
	// --------------------
	// Http API
	// --------------------
	router := gin.Default()
	router.POST("/kms/app/create", httpAPI.HandleCreateApp)
	router.POST("/kms/app/query", httpAPI.HandleQueryApp)
	router.POST("/kms/app/list", httpAPI.HandleListApp)
	router.POST("/kms/app/modify/profile", httpAPI.HandleModifyAppProfile)
	router.POST("/kms/key/generate", httpAPI.HandleGenerateKey)
	router.POST("/kms/key/query", httpAPI.HandleQueryKey)
	router.POST("/kms/key/list", httpAPI.HandleListKey)
	router.POST("/kms/key/modify/profile", httpAPI.HandleModifyKeyProfile)
	router.POST("/kms/key/modify/status", httpAPI.HandleModifyKeyStatus)
	router.POST("/kms/key/activate", httpAPI.HandleActivateKey)
	router.POST("/kms/license/verify", httpAPI.HandleVerifyLicense)

	go router.Run(httpAddrArg)

	// --------------------
	// GRPC API
	// --------------------
	listener, err := net.Listen("tcp", grpcAddrArg)
	if nil != err {
		panic(err)
	}
	server := grpc.NewServer()
	// TODO register server
	server.Serve(listener)
}
