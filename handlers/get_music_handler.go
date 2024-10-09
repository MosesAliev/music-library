package handlers

import (
	"log"
	"music_library/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMusicHandler(c *gin.Context) {
	log.Println("here")
	song, _ := c.GetQuery("song")

	group, _ := c.GetQuery("group")

	music := model.Music{
		Song:  song,
		Group: group,
	}

	music.GetMusic(model.DB) // ищем песню по названию и группе

	c.IndentedJSON(http.StatusOK, music) // сериализуем JSON

}
