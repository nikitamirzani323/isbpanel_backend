package controllers

import (
	"log"
	"strconv"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nikitamirzani323/isbpanel_backend/entities"
	"github.com/nikitamirzani323/isbpanel_backend/helpers"
	"github.com/nikitamirzani323/isbpanel_backend/models"
)

const Fieldmovie_home_redis = "LISTMOVIE_BACKEND_ISBPANEL"
const Fieldgenre_home_redis = "LISTGENRE_BACKEND_ISBPANEL"
const Fieldmovieseries_home_redis = "LISTMOVIESERIES_BACKEND_ISBPANEL"
const Fieldmovieseriesseason_home_redis = "LISTMOVIESEASON_BACKEND_ISBPANEL"

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
	log.Println(client.Movie_page)
	if client.Movie_search != "" {
		val_movie := helpers.DeleteRedis(Fieldmovie_home_redis + "_" + strconv.Itoa(client.Movie_page) + "_" + client.Movie_search)
		log.Printf("Redis Delete BACKEND MOVIE : %d", val_movie)
	}
	var obj entities.Model_movie
	var arraobj []entities.Model_movie
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldmovie_home_redis + "_" + strconv.Itoa(client.Movie_page) + "_" + client.Movie_search)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	perpage_RD, _ := jsonparser.GetInt(jsonredis, "perpage")
	totalrecord_RD, _ := jsonparser.GetInt(jsonredis, "totalrecord")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		movie_id, _ := jsonparser.GetInt(value, "movie_id")
		movie_date, _ := jsonparser.GetString(value, "movie_date")
		movie_type, _ := jsonparser.GetString(value, "movie_type")
		movie_title, _ := jsonparser.GetString(value, "movie_title")
		movie_label, _ := jsonparser.GetString(value, "movie_label")
		movie_slug, _ := jsonparser.GetString(value, "movie_slug")
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
			moviegenre_id, _ := jsonparser.GetInt(value, "moviegenre_id")
			moviegenre_name, _ := jsonparser.GetString(value, "moviegenre_name")
			objmoviegenre.Moviegenre_id = int(moviegenre_id)
			objmoviegenre.Moviegenre_name = moviegenre_name
			arraobjmoviegenre = append(arraobjmoviegenre, objmoviegenre)
		})

		var objmoviesource entities.Model_moviesource
		var arraobjmoviesource []entities.Model_moviesource
		record_moviesource_RD, _, _, _ := jsonparser.Get(value, "movie_source")
		jsonparser.ArrayEach(record_moviesource_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			moviesource_id, _ := jsonparser.GetInt(value, "moviesource_id")
			moviesource_stream, _ := jsonparser.GetString(value, "moviesource_stream")
			moviesource_url, _ := jsonparser.GetString(value, "moviesource_url")
			objmoviesource.Moviesource_id = int(moviesource_id)
			objmoviesource.Moviesource_stream = moviesource_stream
			objmoviesource.Moviesource_url = moviesource_url
			arraobjmoviesource = append(arraobjmoviesource, objmoviesource)
		})

		obj.Movie_id = int(movie_id)
		obj.Movie_date = movie_date
		obj.Movie_type = movie_type
		obj.Movie_title = movie_title
		obj.Movie_label = movie_label
		obj.Movie_slug = movie_slug
		obj.Movie_descp = movie_descp
		obj.Movie_thumbnail = movie_thumbnail
		obj.Movie_year = int(movie_year)
		obj.Movie_rating = float32(movie_rating)
		obj.Movie_imdb = float32(movie_imdb)
		obj.Movie_view = int(movie_view)
		obj.Movie_status = movie_status
		obj.Movie_statuscss = movie_statuscss
		obj.Movie_genre = arraobjmoviegenre
		obj.Movie_source = arraobjmoviesource
		obj.Movie_create = movie_create
		obj.Movie_update = movie_update
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_movieHome(client.Movie_search, client.Movie_page)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldmovie_home_redis+"_"+strconv.Itoa(client.Movie_page)+"_"+client.Movie_search, result, 5*time.Minute)
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
func Moviesave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_moviesave)
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

	result, err := models.Save_movie(
		client_admin,
		client.Movie_name, client.Movie_label, client.Movie_slug, client.Movie_tipe, client.Movie_descp, client.Movie_urlmovie,
		string(client.Movie_gender), string(client.Movie_source),
		client.Sdata, client.Movie_id, client.Movie_year, client.Movie_status, client.Movie_imdb)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val_movie := helpers.DeleteRedis(Fieldmovie_home_redis + "_1_")
	log.Printf("Redis Delete BACKEND MOVIE : %d", val_movie)
	return c.JSON(result)
}
func Moviedelete(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_moviedelete)
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

	result, err := models.Delete_movie(client_admin, client.Movie_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val_movie := helpers.DeleteRedis(Fieldmovie_home_redis + "_1_")
	log.Printf("Redis Delete BACKEND MOVIE : %d", val_movie)
	return c.JSON(result)
}
func Moviehomeseries(c *fiber.Ctx) error {
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
	log.Println(client.Movie_page)
	if client.Movie_search != "" {
		val_movie := helpers.DeleteRedis(Fieldmovieseries_home_redis + "_" + strconv.Itoa(client.Movie_page) + "_" + client.Movie_search)
		log.Printf("Redis Delete BACKEND MOVIE SERIES : %d", val_movie)
	}
	var obj entities.Model_movieseries
	var arraobj []entities.Model_movieseries
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldmovieseries_home_redis + "_" + strconv.Itoa(client.Movie_page) + "_" + client.Movie_search)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	perpage_RD, _ := jsonparser.GetInt(jsonredis, "perpage")
	totalrecord_RD, _ := jsonparser.GetInt(jsonredis, "totalrecord")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		movie_id, _ := jsonparser.GetInt(value, "movie_id")
		movie_date, _ := jsonparser.GetString(value, "movie_date")
		movie_type, _ := jsonparser.GetString(value, "movie_type")
		movie_title, _ := jsonparser.GetString(value, "movie_title")
		movie_label, _ := jsonparser.GetString(value, "movie_label")
		movie_slug, _ := jsonparser.GetString(value, "movie_slug")
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
			moviegenre_id, _ := jsonparser.GetInt(value, "moviegenre_id")
			moviegenre_name, _ := jsonparser.GetString(value, "moviegenre_name")
			objmoviegenre.Moviegenre_id = int(moviegenre_id)
			objmoviegenre.Moviegenre_name = moviegenre_name
			arraobjmoviegenre = append(arraobjmoviegenre, objmoviegenre)
		})

		var objmovieseason entities.Model_movieseason
		var arraobjmovieseason []entities.Model_movieseason
		record_movieseason_RD, _, _, _ := jsonparser.Get(value, "movie_season")
		jsonparser.ArrayEach(record_movieseason_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			movieseason_id, _ := jsonparser.GetInt(value, "movieseason_id")
			movieseason_display, _ := jsonparser.GetInt(value, "movieseason_display")
			movieseason_name, _ := jsonparser.GetString(value, "movieseason_name")
			objmovieseason.Movieseason_id = int(movieseason_id)
			objmovieseason.Movieseason_display = int(movieseason_display)
			objmovieseason.Movieseason_name = movieseason_name
			arraobjmovieseason = append(arraobjmovieseason, objmovieseason)
		})

		obj.Movie_id = int(movie_id)
		obj.Movie_date = movie_date
		obj.Movie_type = movie_type
		obj.Movie_title = movie_title
		obj.Movie_label = movie_label
		obj.Movie_slug = movie_slug
		obj.Movie_descp = movie_descp
		obj.Movie_thumbnail = movie_thumbnail
		obj.Movie_year = int(movie_year)
		obj.Movie_rating = float32(movie_rating)
		obj.Movie_imdb = float32(movie_imdb)
		obj.Movie_view = int(movie_view)
		obj.Movie_status = movie_status
		obj.Movie_statuscss = movie_statuscss
		obj.Movie_genre = arraobjmoviegenre
		obj.Movie_season = arraobjmovieseason
		obj.Movie_create = movie_create
		obj.Movie_update = movie_update
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_movieseriesHome(client.Movie_search, client.Movie_page)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldmovieseries_home_redis+"_"+strconv.Itoa(client.Movie_page)+"_"+client.Movie_search, result, 5*time.Minute)
		log.Println("MOVIE SERIES MYSQL")
		return c.JSON(result)
	} else {
		log.Println("MOVIE SERIES CACHE")
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
func Movieseriessave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_movieseriessave)
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

	result, err := models.Save_movieseries(
		client_admin,
		client.Movie_name, client.Movie_label, client.Movie_slug, client.Movie_tipe, client.Movie_descp, client.Movie_urlmovie,
		string(client.Movie_gender),
		client.Sdata, client.Movie_id, client.Movie_year, client.Movie_status, client.Movie_imdb)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val_movie := helpers.DeleteRedis(Fieldmovie_home_redis + "_1_")
	log.Printf("Redis Delete BACKEND MOVIE : %d", val_movie)
	return c.JSON(result)
}
func Seasonhome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_movieseason)
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

	var obj entities.Model_movieseason
	var arraobj []entities.Model_movieseason
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldmovieseriesseason_home_redis + "_" + strconv.Itoa(client.Movie_id))
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		movieseason_id, _ := jsonparser.GetInt(value, "movieseason_id")
		movieseason_name, _ := jsonparser.GetString(value, "movieseason_name")
		movieseason_display, _ := jsonparser.GetInt(value, "movieseason_display")

		obj.Movieseason_id = int(movieseason_id)
		obj.Movieseason_name = movieseason_name
		obj.Movieseason_display = int(movieseason_display)
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_season(client.Movie_id)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldmovieseriesseason_home_redis+"_"+strconv.Itoa(client.Movie_id), result, 5*time.Minute)
		log.Println("SEASON MYSQL")
		return c.JSON(result)
	} else {
		log.Println("SEASON CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
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

type responseuploadcloudflare struct {
	Status bool        `json:"success"`
	Record interface{} `json:"result"`
}

func Movieuploadcloud(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_cloudflaremovieupload)
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

	axios := resty.New()
	resp, err := axios.R().
		SetResult(responseuploadcloudflare{}).
		SetError(responseuploadcloudflare{}).
		SetAuthToken("8x02SSARJt_A5B77KnL2oW74qwDPFKA_9DORcf1-").
		SetHeader("Accept", "*/*").
		SetHeader("Content-Type", "image/jpeg,image/png").
		SetFiles(map[string]string{
			"file": client.Movie_raw,
		}).
		SetFormData(map[string]string{
			"requireSignedURLs": `true`,
		}).
		SetContentLength(true).
		Post("https://api.cloudflare.com/client/v4/accounts/dc5ba4b3b061907a5e1f8cdf1ae1ec96/images/v1")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responseuploadcloudflare)
	return c.JSON(fiber.Map{
		"status": result.Status,
		"record": result.Record,
	})
}
func Movieupdatecloud(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_cloudflaremovieupdate)
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
	flag_lock := false
	if client.Movie_tipe == "LOCK" {
		flag_lock = true
	}
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responseuploadcloudflare{}).
		SetError(responseuploadcloudflare{}).
		SetAuthToken("8x02SSARJt_A5B77KnL2oW74qwDPFKA_9DORcf1-").
		SetBody(map[string]interface{}{
			"id":                client.Movie_id,
			"requireSignedURLs": flag_lock,
		}).
		SetContentLength(true).
		Patch("https://api.cloudflare.com/client/v4/accounts/dc5ba4b3b061907a5e1f8cdf1ae1ec96/images/v1/" + client.Movie_id)
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responseuploadcloudflare)
	return c.JSON(fiber.Map{
		"status": result.Status,
		"record": result.Record,
	})
}
func Moviedeletecloud(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_cloudflaremoviedelete)
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
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responseuploadcloudflare{}).
		SetError(responseuploadcloudflare{}).
		SetAuthToken("8x02SSARJt_A5B77KnL2oW74qwDPFKA_9DORcf1-").
		SetContentLength(true).
		Delete("https://api.cloudflare.com/client/v4/accounts/dc5ba4b3b061907a5e1f8cdf1ae1ec96/images/v1/" + client.Movie_id)
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responseuploadcloudflare)
	return c.JSON(fiber.Map{
		"status": result.Status,
		"record": result.Record,
	})
}
func Moviecloud(c *fiber.Ctx) error {
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responseuploadcloudflare{}).
		SetError(responseuploadcloudflare{}).
		SetAuthToken("8x02SSARJt_A5B77KnL2oW74qwDPFKA_9DORcf1-").
		Get("https://api.cloudflare.com/client/v4/accounts/dc5ba4b3b061907a5e1f8cdf1ae1ec96/images/v1")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responseuploadcloudflare)
	return c.JSON(fiber.Map{
		"status": result.Status,
		"record": result.Record,
	})
}
