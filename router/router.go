package router

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(200, gin.H{
				"msg": "Login Success",
			})
		})
		auth.POST("/register", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(200, gin.H{
				"msg": "Register Success",
			})
		})
	}
	return r
}
