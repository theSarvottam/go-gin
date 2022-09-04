package service

import (
	"bytes"
	"encoding/json"

	db "github.com/theSarvottam/go-gin/db"
	model "github.com/theSarvottam/go-gin/model"
)

func GetAlbums() []model.Album {
	return db.GetAlbums()
}

func CreateAlbum(ar model.CreateAlbumRequest) model.CreateAlbumResponse {

	var album model.Album
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(ar)
	json.Unmarshal(reqBodyBytes.Bytes(), &album)

	album_inserted := db.SaveAlbum(&album)

	var ares model.CreateAlbumResponse
	resBodyBytes := new(bytes.Buffer)
	json.NewEncoder(resBodyBytes).Encode(album_inserted)
	json.Unmarshal(reqBodyBytes.Bytes(), &ares)

	return ares

}
