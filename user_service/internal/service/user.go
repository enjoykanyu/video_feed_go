package service

import (
	"context"
	repository "user_service/internal/repostry"
	"user_service/internal/store/model"
	"user_service/proto"
)

type UserService struct {
	proto.UnimplementedUserServiceServer // 关键！必须嵌入

	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// func (s *UserService) Register(username, password, email string) (uint, error) {
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return 0, err
// 	}

// 	user := &model.User{
// 		Username: username,
// 		Password: string(hashedPassword),
// 		Email:    email,
// 	}

// 	if err := s.repo.Create(user); err != nil {
// 		return 0, err
// 	}
// 	return user.ID, nil
// }

func (s *UserService) Register(
	ctx context.Context,
	req *proto.UserRequest,
) (*proto.Response, error) {
	// 实现具体注册逻辑
	// 例如：
	user, err := s.repo.Create(&model.User{
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Phone,
	})
	if err != nil {
		return nil, err
	}
	return &proto.Response{UserId: user.Id}, nil

	// 示例返回（替换为实际逻辑）
	// return &proto.Response{UserId: 1}, nil
}
