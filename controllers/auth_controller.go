package controllers

import (
	"fast_demo/global"
	"fast_demo/models"
	"fast_demo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register 处理用户注册请求
func Register(ctx *gin.Context) {
	var user models.User

	// ShouldBindJSON — 读取请求的 JSON 请求体，按字段名匹配自动填充到 user 结构体
	// JSON 多字段忽略、少字段留零值、类型不匹配则报错
	if err := ctx.ShouldBindJSON(&user); err != nil {
		// 格式不对（比如没传 JSON 或字段类型错误），返回 400 错误
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 对明文密码进行 bcrypt 加密，防止数据库泄露时密码裸奔
	hashedPwd, err := utils.HashPassword(user.Password)
	if err != nil {
		// 加密失败（系统问题），返回 500
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// 用加密后的哈希值替换明文密码，后续存入数据库
	user.Password = hashedPwd

	// 生成 JWT 令牌，后续每次请求客户端带上它来表明身份
	token, err := utils.GenerateJWT(user.Username)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// AutoMigrate — 检查数据库，表不存在则自动创建，缺字段则新增列，不会删表删列
	// 注意：放在每个请求里效率低，通常只在项目启动时执行一次
	if err := global.Db.AutoMigrate(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 把用户数据写入数据库
	// .Create() 返回 *gorm.DB，需要加 .Error 才能取出真正的错误
	if err := global.Db.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 注册成功，返回 JWT 给客户端
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
