package handlers

import (
	"music_library/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMusic(c *gin.Context) {
	song, _ := c.GetQuery("song")

	group, _ := c.GetQuery("group")

	music := model.Music{
		Song:  song,
		Group: group,
	}

	err := music.Get() // ищем песню по названию и группе

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, &model.ResponseError{
			Err: err.Error(),
		})

		return
	}

	c.IndentedJSON(http.StatusOK, music) // сериализуем JSON

}
