package controller

import (
	"github.com/NomanSalhab/go_gin_course/entity"
	"github.com/NomanSalhab/go_gin_course/service"
	"github.com/NomanSalhab/go_gin_course/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error /*entity.Video*/
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error /*entity.Video*/ {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	err = validate.Struct(video)
	if err != nil {
		return err
	}

	c.service.Save(video)
	return nil /*video*/
}
