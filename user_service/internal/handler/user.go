package handler

import (
	"context"
	"user_service/internal/service"
	"user_service/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	service *service.UserService
	proto.UnimplementedUserServiceServer
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Register(ctx context.Context, req *proto.UserRequest) (*proto.Response, error) {
	res, err := h.service.Register(ctx, req)
	if err != nil {
		if isDuplicateError(err) {
			return nil, status.Error(codes.AlreadyExists, "用户已存在")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.Response{
		Code:    200,
		UserId:  uint32(res.UserId),
		Message: "注册成功",
	}, nil
}

func isDuplicateError(err error) bool {
	// 根据具体数据库错误判断
	return true
}
