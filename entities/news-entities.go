package entities

type Model_news struct {
	News_id     int    `json:"news_id"`
	News_title  string `json:"news_title"`
	News_descp  string `json:"news_descp"`
	News_url    string `json:"news_url"`
	News_image  string `json:"news_image"`
	News_create string `json:"news_create"`
	News_update string `json:"news_update"`
}
