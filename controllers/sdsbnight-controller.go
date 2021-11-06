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

func Sdsbnighthome(c *fiber.Ctx) error {
	field_redis := "LISTSDSBNIGHT_SDSB4D"

	var obj entities.Responseredis_sdsbnight
	var arraobj []entities.Responseredis_sdsbnight
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		sdsbnight_id, _ := jsonparser.GetInt(value, "sdsbnight_id")
		sdsbnight_date, _ := jsonparser.GetString(value, "sdsbnight_date")
		sdsbnight_prize1, _ := jsonparser.GetString(value, "sdsbnight_prize1")
		sdsbnight_prize2, _ := jsonparser.GetString(value, "sdsbnight_prize2")
		sdsbnight_prize3, _ := jsonparser.GetString(value, "sdsbnight_prize3")
		sdsbnight_create, _ := jsonparser.GetString(value, "sdsbnight_create")
		sdsbnight_update, _ := jsonparser.GetString(value, "sdsbnight_update")

		obj.Sdsbnight_id = int(sdsbnight_id)
		obj.Sdsbnight_date = sdsbnight_date
		obj.Sdsbnight_prize1 = sdsbnight_prize1
		obj.Sdsbnight_prize2 = sdsbnight_prize2
		obj.Sdsbnight_prize3 = sdsbnight_prize3
		obj.Sdsbnight_create = sdsbnight_create
		obj.Sdsbnight_update = sdsbnight_update
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_sdsbnightHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(field_redis, result, 0)
		log.Println("SDSBNIGHT MYSQL")
		return c.JSON(result)
	} else {
		log.Println("SDSBNIGHT CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func SdsbnightSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_sdsbnightsave)
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

	result, err := models.Save_sdsbnightHome(
		client_admin,
		client.Tanggal, client.Sdata, client.Idrecord)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	field_redis := "LISTSDSBNIGHT_SDSB4D"
	val_master := helpers.DeleteRedis(field_redis)
	log.Printf("Redis Delete BACKEND LISTSDSBNIGHT_SDSB4D : %d", val_master)
	field_redis_api := "SDSB4D_LISTSDSBNIGHT_API"
	val_api := helpers.DeleteRedis(field_redis_api)
	log.Printf("Redis Delete API LISTSDSBDAY_SDSB4D : %d", val_api)
	return c.JSON(result)
}
func SdsbnightGeneratorSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_sdsbnightprizesave)
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

	field := ""
	switch client.Tipe {
	case "prize1":
		field = "prize1_sdsb4dnight"
	case "prize2":
		field = "prize2_sdsb4dnight"
	case "prize3":
		field = "prize3_sdsb4dnight"
	}

	result, err := models.Save_sdsbnightGenerator(
		client_admin,
		field, client.Prize, client.Sdata, client.Idrecord)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	field_redis := "LISTSDSBNIGHT_SDSB4D"
	val_master := helpers.DeleteRedis(field_redis)
	log.Printf("Redis Delete BACKEND LISTSDSBNIGHT_SDSB4D : %d", val_master)
	field_redis_api := "SDSB4D_LISTSDSBNIGHT_API"
	val_api := helpers.DeleteRedis(field_redis_api)
	log.Printf("Redis Delete API LISTSDSBNIGHT_SDSB4D : %d", val_api)
	return c.JSON(result)
}
func SdsbnightGeneratorNumber(c *fiber.Ctx) error {
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_admin, _ := helpers.Parsing_Decry(temp_decp, "==")

	result, err := models.Save_Generatornight(client_admin)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	field_redis := "LISTSDSBNIGHT_SDSB4D"
	val_master := helpers.DeleteRedis(field_redis)
	log.Printf("Redis Delete BACKEND LISTSDSBNIGHT_SDSB4D : %d", val_master)
	field_redis_api := "SDSB4D_LISTSDSBNIGHT_API"
	val_api := helpers.DeleteRedis(field_redis_api)
	log.Printf("Redis Delete API LISTSDSBNIGHT_SDSB4D : %d", val_api)
	return c.JSON(result)
}
