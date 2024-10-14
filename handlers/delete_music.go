package handlers

import (
	"music_library/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteMusic(c *gin.Context) {
	song, _ := c.GetQuery("song")

	group, _ := c.GetQuery("group")

	music := model.Music{
		Song:  song,
		Group: group,
	}

	responseString, err := music.Delete()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, &model.ResponseError{
			Err: responseString,
		})

		return
	}

	c.IndentedJSON(http.StatusOK, &model.ResponseError{
		Err: responseString,
	})

}
