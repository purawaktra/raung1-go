package functions

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/purawaktra/raung1-go/utils"
)

func SetBaseHeader(ctx *gin.Context, requestId string, responseTime string, code string, message string) {
	ctx.Header("host", utils.AppCode)
	ctx.Header("request-id", requestId)
	ctx.Header("response-id", uuid.New().String())
	ctx.Header("response-time", responseTime)
	ctx.Header("response-code", code)
	ctx.Header("response-message", message)
}
