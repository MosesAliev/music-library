package handlers

import (
	"music_library/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateMusic(c *gin.Context) {
	var music model.Music
	c.ShouldBindJSON(&music) // десериализуем JSON

	responseString, err := music.Update()

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
