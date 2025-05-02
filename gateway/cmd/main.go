package main

import (
	// 控制器所在路径
	"gateway/internal/controller"
	"gateway/internal/grpc_client"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 初始化gRPC客户端
	userClient, err := grpc_client.NewUserClient("localhost:50051") // 替换为实际地址
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	} else {
		log.Println("gRPC client created successfully")
	}
	defer userClient.Close()

	// 2. 创建控制器
	userController := controller.NewUserController(userClient)

	// 3. 配置Gin路由
	router := gin.Default()
	api := router.Group("/api/v1")
	{
		api.POST("/register", userController.Register)
	}

	// 4. 启动HTTP服务器
	router.Run("127.0.0.1:8089") // listen and serve on 0.0.0.0:8080

	// go func() {
	// 	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		log.Fatalf("Server error: %v", err)
	// 	}
	// }()

	// 5. 优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
}
