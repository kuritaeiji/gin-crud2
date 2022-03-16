package controller

import (
	"gin-crud2/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserController interface {
	Index(ctx *gin.Context)
	Show(ctx *gin.Context)
	Create(ctx *gin.Context)
}

type userController struct {
	service service.UserService
}

func NewUserController() UserController {
	return &userController{service.NewUserService()}
}

func (c *userController) Index(ctx *gin.Context) {
	users := c.service.FindAll()
	ctx.JSON(200, users)
}

func (c *userController) Show(ctx *gin.Context) {
	user, err := c.service.Find(ctx.Params.ByName("id"))
	if err != nil && err == gorm.ErrRecordNotFound {
		ctx.AbortWithStatus(400)
	}
	ctx.JSON(200, user)
}

func (c *userController) Create(ctx *gin.Context) {
	user, err := c.service.Create(ctx)
	if err != nil {
		if e, ok := err.(validator.ValidationErrors); ok {
			ctx.JSON(400, gin.H{
				"message": e.Error(),
			})
			return
		}
		panic(err)
	}
	ctx.JSON(200, user)
}
