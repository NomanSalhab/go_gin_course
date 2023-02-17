package main

import (
	"io"
	"net/http"
	"os"

	"github.com/NomanSalhab/go_gin_course/controller"
	"github.com/NomanSalhab/go_gin_course/middlewares"
	"github.com/NomanSalhab/go_gin_course/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	server := gin.New()

	// server.Use(gin.Recovery(), gin.Logger())

	server.Use(gin.Recovery(), middlewares.Logger(),
		middlewares.BasicAuth(),
		gindump.Dump(),
	)

	// server.Use(gin.Logger())

	// server.GET("/test", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "OK!!",
	// 	})
	// })

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := VideoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is Valid!"})
		}
	})

	server.Run(":8016")
}
