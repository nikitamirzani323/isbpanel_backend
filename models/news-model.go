package models

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/isbpanel_backend/configs"
	"github.com/nikitamirzani323/isbpanel_backend/db"
	"github.com/nikitamirzani323/isbpanel_backend/entities"
	"github.com/nikitamirzani323/isbpanel_backend/helpers"
	"github.com/nleeper/goment"
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
			ORDER BY idnews DESC  LIMIT 100 
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
func Save_news(admin, title, descp, url, image string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	sql_insert := `
		insert into
		` + configs.DB_tbl_trx_news + ` (
			idnews , title_news, descp_news, url_news, img_news, 
			createnews, createdatenews
		) values (
			? ,?, ?, ?, ?,
			?, ?
		)
	`
	stmt_insert, e_insert := con.PrepareContext(ctx, sql_insert)
	helpers.ErrorCheck(e_insert)
	defer stmt_insert.Close()
	field_column := configs.DB_tbl_trx_news + tglnow.Format("YYYY")
	idrecord_counter := Get_counter(field_column)
	res_newrecord, e_newrecord := stmt_insert.ExecContext(
		ctx,
		tglnow.Format("YY")+strconv.Itoa(idrecord_counter),
		title, descp, url, image,
		admin,
		tglnow.Format("YYYY-MM-DD HH:mm:ss"))
	helpers.ErrorCheck(e_newrecord)
	insert, e := res_newrecord.RowsAffected()
	helpers.ErrorCheck(e)
	if insert > 0 {
		flag = true
		msg = "Succes"
		log.Println("Data Berhasil di save")
	}

	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = nil
		res.Time = time.Since(render_page).String()
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = nil
		res.Time = time.Since(render_page).String()
	}

	return res, nil
}
func Delete_news(admin string, idnews int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	flag := false

	flag = CheckDB(configs.DB_tbl_trx_news, "idnews", strconv.Itoa(idnews))
	if flag {
		sql_delete := `
			DELETE FROM
			` + configs.DB_tbl_trx_news + ` 
			WHERE idnews=? 
		`
		stmt_delete, e_delete := con.PrepareContext(ctx, sql_delete)
		helpers.ErrorCheck(e_delete)
		defer stmt_delete.Close()
		rec_delete, e_delete := stmt_delete.ExecContext(ctx, idnews)

		helpers.ErrorCheck(e_delete)
		delete, e := rec_delete.RowsAffected()
		helpers.ErrorCheck(e)
		if delete > 0 {
			flag = true
			msg = "Succes"
			log.Println("Data Berhasil di delete")
		}
	} else {
		msg = "Data Not Found"
	}

	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = nil
		res.Time = time.Since(render_page).String()
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = nil
		res.Time = time.Since(render_page).String()
	}

	return res, nil
}
