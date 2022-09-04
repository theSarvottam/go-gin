package model

// book album
type Album struct {
	Id     int64
	Name   string
	Author string
}

type CreateAlbumRequest struct {
	Id     int64
	Name   string
	Author string
}

type CreateAlbumResponse struct {
	Id   int64
	Name string
}
