package models

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/isbpanel_backend/configs"
	"github.com/nikitamirzani323/isbpanel_backend/db"
	"github.com/nikitamirzani323/isbpanel_backend/entities"
	"github.com/nikitamirzani323/isbpanel_backend/helpers"
)

func Fetch_newsHome() (helpers.Response, error) {
	var obj entities.Model_news
	var arraobj []entities.Model_news
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			idnews , title_news, descp_news, 
			url_news , img_news, 
			createnews, COALESCE(createdatenews,""), updatenews, COALESCE(updatedatenews,"")  
			FROM ` + configs.DB_tbl_trx_news + ` 
			ORDER BY idnews DESC  
		`

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idnews_db                                                          int
			title_news_db, descp_news_db, url_news_db, img_news_db             string
			createnews_db, createdatenews_db, updatenews_db, updatedatenews_db string
		)

		err = row.Scan(
			&idnews_db, &title_news_db, &descp_news_db, &url_news_db, &img_news_db,
			&createnews_db, &createdatenews_db, &updatenews_db, &updatedatenews_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		if createnews_db != "" {
			create = createnews_db + ", " + createdatenews_db
		}
		if updatenews_db != "" {
			update = updatenews_db + ", " + updatedatenews_db
		}

		obj.News_id = idnews_db
		obj.News_title = title_news_db
		obj.News_descp = descp_news_db
		obj.News_url = url_news_db
		obj.News_image = img_news_db
		obj.News_create = create
		obj.News_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}
