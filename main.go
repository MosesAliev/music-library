package main

import (
	"music_library/model"
	"music_library/router"
)

func init() {
	model.DB.ConnectDB()

	r := router.SetupRouter()

	r.Run()

}

func main() {

}
