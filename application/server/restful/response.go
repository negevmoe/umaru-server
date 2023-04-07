package restful

import (
	"github.com/gin-gonic/gin"
)

var response = responseT{}

type responseT struct {
}

func (r responseT) Success(ctx *gin.Context, data any) {
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}

func (r responseT) Error(ctx *gin.Context, err error) {
	ctx.JSON(200, err)
}
