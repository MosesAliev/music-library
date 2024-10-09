package model

import (
	"log"
	"time"
)

type Music struct {
	Group       string `gorm:"primaryKey"`
	Song        string `gorm:"primaryKey"`
	ReleaseDate time.Time
	Text        string
	Link        string
}

// Добавляем информацию о песне в БД
func (music Music) AddMusic(DB Dbinstance) {
	res := DB.Db.Create(&music)

	if res.Error != nil {
		log.Println(res.Error)

	}

}

// Поиск песни по названию и группе
func (music *Music) GetMusic(DB Dbinstance) {
	res := DB.Db.First(&music)

	if res.Error != nil {
		log.Println(res.Error)

	}

}
