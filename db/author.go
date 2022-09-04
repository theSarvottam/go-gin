package db

import model "github.com/theSarvottam/go-gin/model"

var albums = []model.Album{
	{
		Id:     1,
		Name:   "My Book",
		Author: "Dummy Author",
	},
}

func GetAlbums() []model.Album {
	return albums
}
