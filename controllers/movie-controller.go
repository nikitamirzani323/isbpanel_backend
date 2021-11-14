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

const Fieldmovie_home_redis = "LISTMOVIE_BACKEND_ISBPANEL"
const Fieldgenre_home_redis = "LISTGENRE_BACKEND_ISBPANEL"

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
	perpage_RD, _ := jsonparser.GetInt(jsonredis, "perpage")
	totalrecord_RD, _ := jsonparser.GetInt(jsonredis, "totalrecord")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		movie_id, _ := jsonparser.GetInt(value, "movie_id")
		movie_type, _ := jsonparser.GetString(value, "movie_type")
		movie_title, _ := jsonparser.GetString(value, "movie_title")
		movie_descp, _ := jsonparser.GetString(value, "movie_descp")
		movie_thumbnail, _ := jsonparser.GetString(value, "movie_thumbnail")
		movie_year, _ := jsonparser.GetInt(value, "movie_year")
		movie_rating, _ := jsonparser.GetFloat(value, "movie_rating")
		movie_imdb, _ := jsonparser.GetFloat(value, "movie_imdb")
		movie_view, _ := jsonparser.GetInt(value, "movie_view")
		movie_status, _ := jsonparser.GetString(value, "movie_status")
		movie_statuscss, _ := jsonparser.GetString(value, "movie_statuscss")
		movie_create, _ := jsonparser.GetString(value, "movie_create")
		movie_update, _ := jsonparser.GetString(value, "movie_update")

		var objmoviegenre entities.Model_moviegenre
		var arraobjmoviegenre []entities.Model_moviegenre
		record_moviegenre_RD, _, _, _ := jsonparser.Get(value, "movie_genre")
		jsonparser.ArrayEach(record_moviegenre_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			moviegenre_name, _ := jsonparser.GetString(value, "moviegenre_name")
			objmoviegenre.Moviegenre_name = moviegenre_name
			arraobjmoviegenre = append(arraobjmoviegenre, objmoviegenre)
		})

		obj.Movie_id = int(movie_id)
		obj.Movie_type = movie_type
		obj.Movie_title = movie_title
		obj.Movie_descp = movie_descp
		obj.Movie_thumbnail = movie_thumbnail
		obj.Movie_year = int(movie_year)
		obj.Movie_rating = float32(movie_rating)
		obj.Movie_imdb = float32(movie_imdb)
		obj.Movie_view = int(movie_view)
		obj.Movie_status = movie_status
		obj.Movie_statuscss = movie_statuscss
		obj.Movie_genre = arraobjmoviegenre
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
			"status":      fiber.StatusOK,
			"message":     message_RD,
			"record":      arraobj,
			"perpage":     perpage_RD,
			"totalrecord": totalrecord_RD,
			"time":        time.Since(render_page).String(),
		})
	}
}
func Genrehome(c *fiber.Ctx) error {
	var obj entities.Model_genre
	var arraobj []entities.Model_genre
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldgenre_home_redis)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		genre_id, _ := jsonparser.GetInt(value, "genre_id")
		genre_name, _ := jsonparser.GetString(value, "genre_name")
		genre_display, _ := jsonparser.GetInt(value, "genre_display")
		genre_create, _ := jsonparser.GetString(value, "genre_create")
		genre_update, _ := jsonparser.GetString(value, "genre_update")

		obj.Genre_id = int(genre_id)
		obj.Genre_name = genre_name
		obj.Genre_display = int(genre_display)
		obj.Genre_create = genre_create
		obj.Genre_update = genre_update
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_genre()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldgenre_home_redis, result, 0)
		log.Println("GENRE MYSQL")
		return c.JSON(result)
	} else {
		log.Println("GENRE CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Genresave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_genresave)
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

	result, err := models.Save_genre(
		client_admin,
		client.Genre_name, client.Sdata, client.Genre_id, client.Genre_display)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val_genre := helpers.DeleteRedis(Fieldgenre_home_redis)
	log.Printf("Redis Delete BACKEND MOVIE GENRE : %d", val_genre)
	return c.JSON(result)
}
func Genredelete(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_genredelete)
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

	result, err := models.Delete_genre(client_admin, client.Genre_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val_genre := helpers.DeleteRedis(Fieldgenre_home_redis)
	log.Printf("Redis Delete BACKEND MOVIE GENRE : %d", val_genre)
	return c.JSON(result)
}
