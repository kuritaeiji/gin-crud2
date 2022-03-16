package service

import (
	"gin-crud2/db"
	"gin-crud2/dto"
	"gin-crud2/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService interface {
	FindAll() []model.User
	Find(id string) (model.User, error)
	Create(ctx *gin.Context) (model.User, error)
	// Update(id string, ctx *gin.Context) (model.User, error)
	// Destroy(id string) error
}

type userService struct {
	db *gorm.DB
}

func NewUserService() UserService {
	return &userService{db.GetDB()}
}

func (s *userService) FindAll() []model.User {
	var users []model.User
	s.db.Find(&users)
	return users
}

func (s *userService) Find(id string) (model.User, error) {
	return s.findUser(id)
}

func (s *userService) Create(ctx *gin.Context) (model.User, error) {
	var userProxy dto.User
	if err := ctx.ShouldBindJSON(&userProxy); err != nil {
		return model.User{}, err
	}
	user := s.unmarshalUser(userProxy)
	s.db.Create(&user)
	return user, nil
}

func (s *userService) findUser(id string) (model.User, error) {
	var user model.User
	if err := s.db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (s *userService) unmarshalUser(userProxy dto.User) model.User {
	var user model.User
	user.FirstName = userProxy.FirstName
	user.LastName = userProxy.LastName
	return user
}
