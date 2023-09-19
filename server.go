package main

import (
	"github.com/GiorgiMakharadze/video-service-API-golang/controller"
	"github.com/GiorgiMakharadze/video-service-API-golang/service"
	"github.com/gin-gonic/gin"
)

var(
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)

)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}