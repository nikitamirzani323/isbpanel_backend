package controllers

import (
	"log"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
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
func Newssave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_newssave)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_admin, _ := helpers.Parsing_Decry(temp_decp, "==")

	result, err := models.Save_news(
		client_admin,
		client.News_title, client.News_descp, client.News_url, client.News_image)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val_news := helpers.DeleteRedis(Fieldnews_home_redis)
	log.Printf("Redis Delete BACKEND NEWS : %d", val_news)
	return c.JSON(result)
}
func Newsdelete(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_newsdelete)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_admin, _ := helpers.Parsing_Decry(temp_decp, "==")

	result, err := models.Delete_news(client_admin, client.News_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val_news := helpers.DeleteRedis(Fieldnews_home_redis)
	log.Printf("Redis Delete BACKEND NEWS : %d", val_news)
	return c.JSON(result)
}
