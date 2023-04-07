package restful

import (
	"github.com/gin-gonic/gin"
)

func MiddleAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//// 请求没有携带token
		//token := ctx.GetHeader("token")
		//if token == "" {
		//	response.Error(ctx, vo.ErrorNew(401, "用户未登录", ""))
		//	ctx.Abort()
		//	return
		//}
		//
		//// 缓存中没有token
		//cacheToken, ok := global.Cache.Get("token:" + ctx.RemoteIP())
		//if !ok {
		//	response.Error(ctx, vo.ErrorNew(402, "登录已过期,请重新登录", ""))
		//	ctx.Abort()
		//	return
		//}
		//
		//// token不一致
		//if token != cacheToken.(string) {
		//	response.Error(ctx, vo.ErrorNew(403, "token失效,请重新登录", ""))
		//	ctx.Abort()
		//	return
		//}
		//
		//global.Cache.Set("token:"+ctx.RemoteIP(), token, time.Second*time.Duration(setting.SERVER_TOKEN_EXPIRATION_TIME))
	}
}
