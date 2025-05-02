package controller

import (
	"gateway/internal/grpc_client"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userClient *grpc_client.UserClient
}

func NewUserController(client *grpc_client.UserClient) *UserController {
	return &UserController{userClient: client}
}

func (c *UserController) Register(ctx *gin.Context) {
	type Request struct {
		Username string `json:"username" binding:"required,min=4,max=20"`
		Password string `json:"password" binding:"required,min=6,max=20"`
		Phone    string `json:"phone" binding:"required"`
	}

	var req Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用gRPC服务
	resp, err := c.userClient.Register(ctx, req.Username, req.Password, req.Phone)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user_id": resp,
	})
}
