package repository

import (
	"log"
	"user_service/internal/store/model"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *model.User) (*model.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	log.Printf("Created user ID: %d", user.Id) // 插入后打印 ID
	return user, nil
}
