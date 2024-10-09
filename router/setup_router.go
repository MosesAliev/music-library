package router

import (
	"music_library/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	var r = gin.Default()

	r.POST("/addmusic", handlers.AddMusicHandlers)

	r.GET("/info", handlers.GetMusicHandler)

	return r
}
