package controllers

import (
	"log"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/isbpanel_backend/entities"
	"github.com/nikitamirzani323/isbpanel_backend/helpers"
	"github.com/nikitamirzani323/isbpanel_backend/models"
)

const Fieldmovie_home_redis = "LISTMOVIE_BACKEND_ISBPANEL"

func Moviehome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_movie)
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
	if client.Movie_search != "" {
		val_tafsirmimpi := helpers.DeleteRedis(Fieldmovie_home_redis + "_" + client.Movie_search)
		log.Printf("Redis Delete BACKEND MOVIE : %d", val_tafsirmimpi)
	}
	var obj entities.Model_movie
	var arraobj []entities.Model_movie
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldmovie_home_redis + "_" + client.Movie_search)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		movie_id, _ := jsonparser.GetInt(value, "movie_id")
		movie_type, _ := jsonparser.GetString(value, "movie_type")
		movie_title, _ := jsonparser.GetString(value, "movie_title")
		movie_descp, _ := jsonparser.GetString(value, "movie_descp")
		movie_year, _ := jsonparser.GetInt(value, "movie_year")
		movie_rating, _ := jsonparser.GetFloat(value, "movie_rating")
		movie_imdb, _ := jsonparser.GetFloat(value, "movie_imdb")
		movie_view, _ := jsonparser.GetInt(value, "movie_view")
		movie_status, _ := jsonparser.GetInt(value, "movie_status")
		movie_create, _ := jsonparser.GetString(value, "movie_create")
		movie_update, _ := jsonparser.GetString(value, "movie_update")

		obj.Movie_id = int(movie_id)
		obj.Movie_type = movie_type
		obj.Movie_title = movie_title
		obj.Movie_descp = movie_descp
		obj.Movie_year = int(movie_year)
		obj.Movie_rating = float32(movie_rating)
		obj.Movie_imdb = float32(movie_imdb)
		obj.Movie_view = int(movie_view)
		obj.Movie_status = int(movie_status)
		obj.Movie_create = movie_create
		obj.Movie_update = movie_update
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_movieHome(client.Movie_search)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldmovie_home_redis+"_"+client.Movie_search, result, 5*time.Minute)
		log.Println("MOVIE MYSQL")
		return c.JSON(result)
	} else {
		log.Println("MOVIE CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
