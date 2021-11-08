package controllers

import (
	"log"
	"time"

	"github.com/buger/jsonparser"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/isbpanel_backend/entities"
	"github.com/nikitamirzani323/isbpanel_backend/helpers"
	"github.com/nikitamirzani323/isbpanel_backend/models"
)

const Fieldnews_home_redis = "LISTNEWS_BACKEND_ISBPANEL"

func Newshome(c *fiber.Ctx) error {
	var obj entities.Model_news
	var arraobj []entities.Model_news
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldnews_home_redis)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		news_id, _ := jsonparser.GetInt(value, "news_id")
		news_title, _ := jsonparser.GetString(value, "news_title")
		news_descp, _ := jsonparser.GetString(value, "news_descp")
		news_url, _ := jsonparser.GetString(value, "news_url")
		news_image, _ := jsonparser.GetString(value, "news_image")
		news_create, _ := jsonparser.GetString(value, "news_create")
		news_update, _ := jsonparser.GetString(value, "news_update")

		obj.News_id = int(news_id)
		obj.News_title = news_title
		obj.News_descp = news_descp
		obj.News_url = news_url
		obj.News_image = news_image
		obj.News_create = news_create
		obj.News_update = news_update
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_newsHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldnews_home_redis, result, 0)
		log.Println("NEWS MYSQL")
		return c.JSON(result)
	} else {
		log.Println("NEWS CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
