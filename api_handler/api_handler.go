package apihandler

import (
	"findindices/api_handler/handler"
	"findindices/service/logic"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	svc := logic.NewFindIndicesService()
	router.POST("find-pairs", handler.Handler(svc))
}
