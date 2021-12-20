package controllers

import (
	"log"
	"strconv"
	"time"

	"bitbucket.org/isbtotogroup/isbpanel_backend/entities"
	"bitbucket.org/isbtotogroup/isbpanel_backend/helpers"
	"bitbucket.org/isbtotogroup/isbpanel_backend/models"
	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

const Fieldcrmisbtv_home_redis = "LISTCRMISBTV_BACKEND_ISBPANEL"
const Fieldcrmduniafilm_home_redis = "LISTCRMDUNIAFILM_BACKEND_ISBPANEL"

func Crmisbtvhome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_crmisbtv)
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
	if client.Crmisbtv_search != "" {
		val_news := helpers.DeleteRedis(Fieldcrmisbtv_home_redis + "_" + strconv.Itoa(client.Crmisbtv_page) + "_" + client.Crmisbtv_search)
		log.Printf("Redis Delete BACKEND NEWS : %d", val_news)
	}
	var obj entities.Model_crmisbtv
	var arraobj []entities.Model_crmisbtv
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldcrmisbtv_home_redis + "_" + strconv.Itoa(client.Crmisbtv_page) + "_" + client.Crmisbtv_search)
	jsonredis := []byte(resultredis)
	perpage_RD, _ := jsonparser.GetInt(jsonredis, "perpage")
	totalrecord_RD, _ := jsonparser.GetInt(jsonredis, "totalrecord")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		crmisbtv_username, _ := jsonparser.GetString(value, "crmisbtv_username")
		crmisbtv_name, _ := jsonparser.GetString(value, "crmisbtv_name")
		crmisbtv_coderef, _ := jsonparser.GetString(value, "crmisbtv_coderef")
		crmisbtv_point, _ := jsonparser.GetInt(value, "crmisbtv_point")
		crmisbtv_status, _ := jsonparser.GetString(value, "crmisbtv_status")
		crmisbtv_lastlogin, _ := jsonparser.GetString(value, "crmisbtv_lastlogin")
		crmisbtv_create, _ := jsonparser.GetString(value, "crmisbtv_create")
		crmisbtv_update, _ := jsonparser.GetString(value, "crmisbtv_update")

		obj.Crmisbtv_username = crmisbtv_username
		obj.Crmisbtv_name = crmisbtv_name
		obj.Crmisbtv_coderef = crmisbtv_coderef
		obj.Crmisbtv_point = int(crmisbtv_point)
		obj.Crmisbtv_status = crmisbtv_status
		obj.Crmisbtv_lastlogin = crmisbtv_lastlogin
		obj.Crmisbtv_create = crmisbtv_create
		obj.Crmisbtv_update = crmisbtv_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_crmisbtv(client.Crmisbtv_search, client.Crmisbtv_page)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldcrmisbtv_home_redis+"_"+strconv.Itoa(client.Crmisbtv_page)+"_"+client.Crmisbtv_search, result, 60*time.Minute)
		log.Println("CRM ISBTV MYSQL")
		return c.JSON(result)
	} else {
		log.Println("CRM ISBTV CACHE")
		return c.JSON(fiber.Map{
			"status":      fiber.StatusOK,
			"message":     "Success",
			"record":      arraobj,
			"perpage":     perpage_RD,
			"totalrecord": totalrecord_RD,
			"time":        time.Since(render_page).String(),
		})
	}
}
func Crmduniafilm(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_crmisbtv)
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
	if client.Crmisbtv_search != "" {
		val_news := helpers.DeleteRedis(Fieldcrmduniafilm_home_redis + "_" + strconv.Itoa(client.Crmisbtv_page) + "_" + client.Crmisbtv_search)
		log.Printf("Redis Delete BACKEND NEWS : %d", val_news)
	}
	var obj entities.Model_crmduniafilm
	var arraobj []entities.Model_crmduniafilm
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldcrmduniafilm_home_redis + "_" + strconv.Itoa(client.Crmisbtv_page) + "_" + client.Crmisbtv_search)
	jsonredis := []byte(resultredis)
	perpage_RD, _ := jsonparser.GetInt(jsonredis, "perpage")
	totalrecord_RD, _ := jsonparser.GetInt(jsonredis, "totalrecord")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		crmduniafilm_username, _ := jsonparser.GetString(value, "crmduniafilm_username")
		crmduniafilm_name, _ := jsonparser.GetString(value, "crmduniafilm_name")

		obj.Crmduniafilm_username = crmduniafilm_username
		obj.Crmduniafilm_name = crmduniafilm_name
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_crmduniafilm(client.Crmisbtv_search, client.Crmisbtv_page)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldcrmduniafilm_home_redis+"_"+strconv.Itoa(client.Crmisbtv_page)+"_"+client.Crmisbtv_search, result, 60*time.Minute)
		log.Println("CRM DUNIA FILM MYSQL")
		return c.JSON(result)
	} else {
		log.Println("CRM DUNIA FILM CACHE")
		return c.JSON(fiber.Map{
			"status":      fiber.StatusOK,
			"message":     "Success",
			"record":      arraobj,
			"perpage":     perpage_RD,
			"totalrecord": totalrecord_RD,
			"time":        time.Since(render_page).String(),
		})
	}
}
