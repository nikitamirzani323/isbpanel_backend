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

const Fieldalbum_home_redis = "LISTALBUM_BACKEND_ISBPANEL"

func Albumhome(c *fiber.Ctx) error {
	var obj entities.Model_album
	var arraobj []entities.Model_album
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldalbum_home_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		album_id, _ := jsonparser.GetInt(value, "album_id")
		album_idcloud, _ := jsonparser.GetString(value, "album_idcloud")
		album_name, _ := jsonparser.GetString(value, "album_name")
		album_signed, _ := jsonparser.GetString(value, "album_signed")
		album_varian, _ := jsonparser.GetString(value, "album_varian")
		album_movieid, _ := jsonparser.GetInt(value, "album_movieid")
		album_movie, _ := jsonparser.GetString(value, "album_movie")
		album_moviestatus, _ := jsonparser.GetString(value, "album_moviestatus")
		album_moviestatuscss, _ := jsonparser.GetString(value, "album_moviestatuscss")
		album_createdate, _ := jsonparser.GetString(value, "album_createdate")

		obj.Album_id = int(album_id)
		obj.Album_idcloud = album_idcloud
		obj.Album_name = album_name
		obj.Album_signed = album_signed
		obj.Album_varian = album_varian
		obj.Album_movieid = int(album_movieid)
		obj.Album_movie = album_movie
		obj.Album_moviestatus = album_moviestatus
		obj.Album_moviestatuscss = album_moviestatuscss
		obj.Album_createdate = album_createdate
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_albumHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldalbum_home_redis, result, 0)
		log.Println("ALBUM MYSQL")
		return c.JSON(result)
	} else {
		log.Println("ALBUM CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Albumsave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_cloudflaresave)
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
	client_admin, idruleadmin := helpers.Parsing_Decry(temp_decp, "==")

	ruleadmin := models.Get_AdminRule("ruleadmingroup", idruleadmin)
	flag := models.Get_listitemsearch(ruleadmin, ",", client.Page)

	if !flag {
		c.Status(fiber.StatusForbidden)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusForbidden,
			"message": "Anda tidak bisa akses halaman ini",
			"record":  nil,
		})
	} else {
		result, err := models.Save_album(
			client_admin,
			string(client.Album_data), client.Sdata)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		val_album := helpers.DeleteRedis(Fieldalbum_home_redis)
		log.Printf("Redis Delete BACKEND ALBUM : %d", val_album)
		return c.JSON(result)
	}
}
