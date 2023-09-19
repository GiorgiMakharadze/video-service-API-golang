package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/GiorgiMakharadze/video-service-API-golang/controller"
	"github.com/GiorgiMakharadze/video-service-API-golang/middlewares"
	"github.com/GiorgiMakharadze/video-service-API-golang/service"
	gindump "github.com/tpkeeper/gin-dump"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, err := os.Create("gin.log")
	if err != nil {
		fmt.Println(err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Video input is valid"})
		}

	})

	server.Run(":8080")
}
