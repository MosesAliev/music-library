package model

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening gorm database", err)
	}

	return gormDB, mock
}

func TestGetMusic(t *testing.T) {
	_, mock := NewMockDB()

	rows := sqlmock.NewRows([]string{"group", "song", "release_date", "text", "link"}).
		AddRow("A", "B", "09-06-2003", "C", "D")

	mock.ExpectQuery("SELECT * FROM musics").WillReturnRows(rows)

	music := Music{}

	music.Get()

}
