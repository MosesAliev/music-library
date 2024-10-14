package main

import (
	"music_library/db"
	"music_library/router"
)

func init() {
	db.ConnectDB()

	r := router.SetupRouter()

	r.Run()

}

func main() {

}
