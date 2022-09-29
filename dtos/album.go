package dtos

type AlbumDto struct {
	Title  string  `form:"title"`
	Artist string  `form:"artist"`
	Price  float64 `form:"price"`
}
