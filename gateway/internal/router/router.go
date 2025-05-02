package router

import (
	"gateway/internal/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userCtrl *controller.UserController) *gin.Engine {
	r := gin.Default()

	// API版本控制
	v1 := r.Group("/api/v1")
	{
		// 用户模块
		user := v1.Group("/user")
		{
			user.POST("/register", userCtrl.Register)
		}

		// 消息模块
		message := v1.Group("/message")
		{
			message.POST("/send", userCtrl.Register)
		}
	}

	return r
}
