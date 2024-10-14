package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"music_library/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddMusic(c *gin.Context) {
	var music model.Music
	c.ShouldBindJSON(&music) // десериализуем JSON

	err := music.Add() // добавляем песню в БД

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, &model.ResponseError{
			Err: err.Error(),
		})

		return
	}

	url := fmt.Sprintf("http://localhost:8080/info?song=%s&group=%s", music.Song, music.Group)

	req, err := http.NewRequest(http.MethodGet, url, nil) // запрос в апи для получения песни

	if err != nil {
		log.Println(err)

	}

	res, _ := http.DefaultClient.Do(req)

	resBody, _ := io.ReadAll(res.Body)

	var foundMusic model.Music
	json.Unmarshal(resBody, &foundMusic)

	c.IndentedJSON(http.StatusOK, foundMusic)

}
