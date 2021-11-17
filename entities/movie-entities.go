package entities

type Model_movie struct {
	Movie_id        int         `json:"movie_id"`
	Movie_date      string      `json:"movie_date"`
	Movie_type      string      `json:"movie_type"`
	Movie_title     string      `json:"movie_title"`
	Movie_descp     string      `json:"movie_descp"`
	Movie_thumbnail string      `json:"movie_thumbnail"`
	Movie_cover     string      `json:"movie_cover"`
	Movie_year      int         `json:"movie_year"`
	Movie_rating    float32     `json:"movie_rating"`
	Movie_imdb      float32     `json:"movie_imdb"`
	Movie_view      int         `json:"movie_view"`
	Movie_genre     interface{} `json:"movie_genre"`
	Movie_status    string      `json:"movie_status"`
	Movie_statuscss string      `json:"movie_statuscss"`
	Movie_create    string      `json:"movie_create"`
	Movie_update    string      `json:"movie_update"`
}
type Model_moviegenre struct {
	Moviegenre_name string `json:"moviegenre_name"`
}
type Model_genre struct {
	Genre_id      int    `json:"genre_id"`
	Genre_name    string `json:"genre_name"`
	Genre_display int    `json:"genre_display"`
	Genre_create  string `json:"genre_create"`
	Genre_update  string `json:"genre_update"`
}

type Controller_movie struct {
	Movie_search string `json:"movie_search"`
}
type Controller_movieupload struct {
	Page      string `json:"page" validate:"required"`
	Sdata     string `json:"sdata" validate:"required"`
	Movie_raw string `json:"movie_raw" validate:"required"`
}
type Controller_movieupdate struct {
	Page       string `json:"page" validate:"required"`
	Sdata      string `json:"sdata" validate:"required"`
	Movie_id   string `json:"movie_id" validate:"required"`
	Movie_tipe string `json:"movie_tipe" validate:"required"`
}
type Controller_genresave struct {
	Page          string `json:"page" validate:"required"`
	Sdata         string `json:"sdata" validate:"required"`
	Genre_id      int    `json:"genre_id"`
	Genre_name    string `json:"genre_name" validate:"required"`
	Genre_display int    `json:"genre_display" validate:"required"`
}
type Controller_genredelete struct {
	Page     string `json:"page" validate:"required"`
	Sdata    string `json:"sdata" validate:"required"`
	Genre_id int    `json:"genre_id" validate:"required"`
}
