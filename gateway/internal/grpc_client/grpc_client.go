// shared/pkg/grpc_client/user_client.go
package grpc_client

import (
	"context"
	"user_service/proto" // 确保proto生成的代码路径正确

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	conn   *grpc.ClientConn
	client proto.UserServiceClient // 假设proto生成的接口为UserServiceClient
}

func NewUserClient(serverAddr string) (*UserClient, error) {
	// 建立gRPC连接
	conn, err := grpc.Dial(
		serverAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()), // 禁用TLS
		grpc.WithBlock(), // 阻塞直到连接成功
	)
	if err != nil {
		return nil, err
	}

	return &UserClient{
		conn:   conn,
		client: proto.NewUserServiceClient(conn),
	}, nil
}

// Register 调用gRPC服务注册方法
func (c *UserClient) Register(
	ctx context.Context,
	username, password, phone string,
) (uint32, error) {
	resp, err := c.client.Register(ctx, &proto.UserRequest{
		Username: username,
		Password: password,
		Phone:    phone,
	})
	if err != nil {
		return 0, err
	}
	return resp.UserId, nil
}

// Close 关闭连接
func (c *UserClient) Close() error {
	return c.conn.Close()
}
