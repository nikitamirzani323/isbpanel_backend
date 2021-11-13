package entities

type Model_movie struct {
	Movie_id     int     `json:"movie_id"`
	Movie_type   string  `json:"movie_type"`
	Movie_title  string  `json:"movie_title"`
	Movie_descp  string  `json:"movie_descp"`
	Movie_year   int     `json:"movie_year"`
	Movie_rating float32 `json:"movie_rating"`
	Movie_imdb   float32 `json:"movie_imdb"`
	Movie_view   int     `json:"movie_view"`
	Movie_status int     `json:"movie_status"`
	Movie_create string  `json:"movie_create"`
	Movie_update string  `json:"movie_update"`
}

type Controller_movie struct {
	Movie_search string `json:"movie_search"`
}
