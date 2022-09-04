package service

import (
	db "github.com/theSarvottam/go-gin/db"
	model "github.com/theSarvottam/go-gin/model"
)

func GetAlbums() []model.Album {
	return db.GetAlbums()
}
