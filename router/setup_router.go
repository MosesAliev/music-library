package router

import (
	"music_library/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	var r = gin.Default()

	r.POST("/add_music", handlers.AddMusic)

	r.GET("/info", handlers.GetMusic)

	r.PATCH("/update_music", handlers.UpdateMusic)

	r.DELETE("/delete_music", handlers.DeleteMusic)

	return r
}
