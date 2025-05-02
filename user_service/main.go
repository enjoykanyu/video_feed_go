package main

import (
	"log"
	"net"
	"user_service/internal/config"
	repository "user_service/internal/repostry"
	"user_service/internal/service"
	"user_service/internal/store/model"
	"user_service/proto"

	"google.golang.org/grpc"
)

// 初始化配置文件
func init() {
	config.InitConfig()
}
func main() {
	// 初始化数据库
	db := model.InitDb()
	defer model.Db.Close()
	// 创建服务实例
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	grpcServer := grpc.NewServer()

	//// 初始化数据库连接后执行迁移
	db.AutoMigrate(&model.User{}, &model.Message{})

	// 注册gRPC服务
	proto.RegisterUserServiceServer(grpcServer, userService)

	// 启动服务
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("gRPC服务启动在 :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
