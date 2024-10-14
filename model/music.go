package model

import (
	"errors"
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
func (music Music) Add() error {
	res := DB.Db.Create(&music)

	if res.Error != nil {
		log.Println(res.Error)

		return res.Error
	}

	return nil
}

// Поиск песни по названию и группе
func (music *Music) Get() error {
	res := DB.Db.First(&music)

	if res.RowsAffected == 0 {
		log.Println("not found")

		return errors.New("not found")

	}

	return nil
}

// Изменение информации о песне
func (music Music) Update() (string, error) {
	res := DB.Db.Save(&music)

	if res.Error != nil {
		log.Println(res.Error)

		return res.Error.Error(), res.Error
	}

	return "updated", nil
}

// Удаление песни
func (music Music) Delete() (string, error) {
	res := DB.Db.Where("song = ? AND group = ?", music.Song, music.Group).Delete(&music)

	if res.Error != nil {
		log.Println(res.Error)

		return res.Error.Error(), res.Error
	}

	return "Deleted", nil
}
