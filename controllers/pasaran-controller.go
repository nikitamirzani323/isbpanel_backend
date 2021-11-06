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

const Field_home_redis = "LISTPASARAN_BACKEND_ISBPANEL"
const Field_keluaran_redis = "LISTKELUARAN_BACKEND_ISBPANEL"

func Pasaranhome(c *fiber.Ctx) error {
	var obj entities.Model_pasaran
	var arraobj []entities.Model_pasaran
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Field_home_redis)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		pasaran_id, _ := jsonparser.GetString(value, "pasaran_id")
		pasaran_name, _ := jsonparser.GetString(value, "pasaran_name")
		pasaran_url, _ := jsonparser.GetString(value, "pasaran_url")
		pasaran_diundi, _ := jsonparser.GetString(value, "pasaran_diundi")
		pasaran_jamjadwal, _ := jsonparser.GetString(value, "pasaran_jamjadwal")
		pasaran_keluaran, _ := jsonparser.GetString(value, "pasaran_keluaran")
		pasaran_create, _ := jsonparser.GetString(value, "pasaran_create")
		pasaran_update, _ := jsonparser.GetString(value, "pasaran_update")

		obj.Pasaran_id = pasaran_id
		obj.Pasaran_name = pasaran_name
		obj.Pasaran_url = pasaran_url
		obj.Pasaran_diundi = pasaran_diundi
		obj.Pasaran_jamjadwal = pasaran_jamjadwal
		obj.Pasaran_keluaran = pasaran_keluaran
		obj.Pasaran_create = pasaran_create
		obj.Pasaran_update = pasaran_update
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_pasaranHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Field_home_redis, result, 0)
		log.Println("PASARAN MYSQL")
		return c.JSON(result)
	} else {
		log.Println("PASARAN CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Pasaransave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_pasaransave)
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

	result, err := models.Save_pasaran(
		client_admin,
		client.Pasaran_id, client.Pasaran_name, client.Pasaran_url,
		client.Pasaran_diundi, client.Pasaran_jamjadwal, client.Sdata)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val_pasaran := helpers.DeleteRedis(Field_home_redis)
	log.Printf("Redis Delete BACKEND PASARAN : %d", val_pasaran)
	return c.JSON(result)
}
func Keluaranhome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_keluaran)
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

	var obj entities.Model_keluaran
	var arraobj []entities.Model_keluaran
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Field_keluaran_redis)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		keluaran_id, _ := jsonparser.GetInt(value, "keluaran_id")
		keluaran_tanggal, _ := jsonparser.GetString(value, "keluaran_tanggal")
		keluaran_periode, _ := jsonparser.GetString(value, "keluaran_periode")
		keluaran_nomor, _ := jsonparser.GetString(value, "keluaran_nomor")

		obj.Keluaran_id = int(keluaran_id)
		obj.Keluaran_tanggal = keluaran_tanggal
		obj.Keluaran_periode = keluaran_periode
		obj.Keluaran_nomor = keluaran_nomor
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_keluaran(client.Pasaran_id)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Field_keluaran_redis+"_"+client.Pasaran_id, result, 0)
		log.Println("KELUARAN MYSQL")
		return c.JSON(result)
	} else {
		log.Println("KELUARAN CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Keluaransave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_keluaransave)
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

	result, err := models.Save_keluaran(
		client_admin,
		client.Pasaran_id, client.Keluaran_tanggal, client.Keluaran_nomor)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val_pasaran := helpers.DeleteRedis(Field_home_redis)
	val_keluaran := helpers.DeleteRedis(Field_keluaran_redis + "_" + client.Pasaran_id)
	log.Printf("Redis Delete BACKEND PASARAN : %d", val_pasaran)
	log.Printf("Redis Delete BACKEND KELUARAN : %d", val_keluaran)
	return c.JSON(result)
}
func Keluarandelete(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_keluarandelete)
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

	result, err := models.Delete_keluaran(client_admin, client.Pasaran_id, client.Keluaran_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val_pasaran := helpers.DeleteRedis(Field_home_redis)
	val_keluaran := helpers.DeleteRedis(Field_keluaran_redis + "_" + client.Pasaran_id)
	log.Printf("Redis Delete BACKEND PASARAN : %d", val_pasaran)
	log.Printf("Redis Delete BACKEND KELUARAN : %d", val_keluaran)
	return c.JSON(result)
}
