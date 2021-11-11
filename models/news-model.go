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

func Fetch_newsHome(search string) (helpers.Response, error) {
	var obj entities.Model_news
	var arraobj []entities.Model_news
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := ""
	sql_select += ""
	sql_select += "SELECT "
	sql_select += "A.idnews , A.title_news, A.descp_news, "
	sql_select += "A.url_news , A.img_news, B.nmcatenews, A.idcatenews, "
	sql_select += "A.createnews, COALESCE(A.createdatenews,''), A.updatenews, COALESCE(A.updatedatenews,'') "
	sql_select += "FROM " + configs.DB_tbl_trx_news + " as A "
	sql_select += "JOIN " + configs.DB_tbl_mst_category + " as B ON B.idcatenews = A.idcatenews "
	if search == "" {
		sql_select += "ORDER BY A.idnews DESC  LIMIT 500  "
	} else {
		sql_select += "WHERE title_news LIKE '%" + search + "%' "
		sql_select += "OR title_news LIKE '%" + search + "%' "
		sql_select += "ORDER BY A.idnews DESC  LIMIT 500  "
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idnews_db, idcatenews_db                                              int
			title_news_db, descp_news_db, url_news_db, img_news_db, nmcatenews_db string
			createnews_db, createdatenews_db, updatenews_db, updatedatenews_db    string
		)

		err = row.Scan(
			&idnews_db, &title_news_db, &descp_news_db, &url_news_db, &img_news_db, &nmcatenews_db, &idcatenews_db,
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
		obj.News_idcategory = idcatenews_db
		obj.News_category = nmcatenews_db
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
func Save_news(admin, sdata, title, descp, url, image string, idrecord, category int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sdata == "New" {
		sql_insert := `
			insert into
			` + configs.DB_tbl_trx_news + ` (
				idnews , idcatenews, title_news, descp_news, url_news, img_news, 
				createnews, createdatenews
			) values (
				? ,?, ?, ?, ?, ?, 
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
			category, title, descp, url, image,
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
	} else {
		sql_update := `
			UPDATE 
			` + configs.DB_tbl_trx_news + ` 
			SET idcatenews=?, title_news=?, descp_news=?, 
			url_news=?,img_news=?,
			updatenews=?, updatedatenews=? 
			WHERE idnews=? 
		`
		stmt_update, e_update := con.PrepareContext(ctx, sql_update)
		helpers.ErrorCheck(e_update)
		defer stmt_update.Close()
		res_newrecord, e_newrecord := stmt_update.ExecContext(
			ctx,
			category, title, descp, url, image,
			admin,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)
		helpers.ErrorCheck(e_newrecord)
		update, e := res_newrecord.RowsAffected()
		helpers.ErrorCheck(e)
		if update > 0 {
			flag = true
			msg = "Succes"
			log.Println("Data Berhasil di update")
		}
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
func Fetch_category() (helpers.Response, error) {
	var obj entities.Model_category
	var arraobj []entities.Model_category
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			idcatenews , nmcatenews, displaycatenews,statuscategory, 
			createcatenews, COALESCE(createdatecatenews,""), updatecatenews, COALESCE(updatedatecatenews,"")  
			FROM ` + configs.DB_tbl_mst_category + ` 
			ORDER BY displaycatenews ASC   
		`

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idcatenews_db, displaycatenews_db                                                  int
			nmcatenews_db, statuscategory_db                                                   string
			createcatenews_db, createdatecatenews_db, updatecatenews_db, updatedatecatenews_db string
		)

		err = row.Scan(
			&idcatenews_db, &nmcatenews_db, &displaycatenews_db, &statuscategory_db,
			&createcatenews_db, &createdatecatenews_db, &updatecatenews_db, &updatedatecatenews_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		statuscss := configs.STATUS_CANCEL
		if createcatenews_db != "" {
			create = createcatenews_db + ", " + createdatecatenews_db
		}
		if updatecatenews_db != "" {
			update = updatecatenews_db + ", " + updatedatecatenews_db
		}
		if statuscategory_db == "Y" {
			statuscss = configs.STATUS_COMPLETE
		}

		obj.Category_id = idcatenews_db
		obj.Category_name = nmcatenews_db
		obj.Category_display = displaycatenews_db
		obj.Category_status = statuscategory_db
		obj.Category_statuscss = statuscss
		obj.Category_create = create
		obj.Category_update = update
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
func Save_category(admin, name, status, sdata string, idrecord, display int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sdata == "New" {
		sql_insert := `
			insert into
			` + configs.DB_tbl_mst_category + ` (
				idcatenews , nmcatenews, displaycatenews, statuscategory, 
				createcatenews, createdatecatenews
			) values (
				? ,?, ?, ?, 
				?, ?
			)
		`
		stmt_insert, e_insert := con.PrepareContext(ctx, sql_insert)
		helpers.ErrorCheck(e_insert)
		defer stmt_insert.Close()
		field_column := configs.DB_tbl_mst_category + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		res_newrecord, e_newrecord := stmt_insert.ExecContext(
			ctx,
			tglnow.Format("YY")+strconv.Itoa(idrecord_counter),
			name, display, status,
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
	} else {
		sql_update := `
			UPDATE 
			` + configs.DB_tbl_mst_category + ` 
			SET nmcatenews=?, displaycatenews=?, statuscategory=?, 
			updatecatenews=?, updatedatecatenews=? 
			WHERE idcatenews=? 
		`
		stmt_update, e_update := con.PrepareContext(ctx, sql_update)
		helpers.ErrorCheck(e_update)
		defer stmt_update.Close()
		res_newrecord, e_newrecord := stmt_update.ExecContext(
			ctx,
			name, display, status,
			admin,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)
		helpers.ErrorCheck(e_newrecord)
		update, e := res_newrecord.RowsAffected()
		helpers.ErrorCheck(e)
		if update > 0 {
			flag = true
			msg = "Succes"
			log.Println("Data Berhasil di update")
		}
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
func Delete_category(admin string, idrecord int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	flag := false

	flag = CheckDB(configs.DB_tbl_mst_category, "idcatenews", strconv.Itoa(idrecord))
	if flag {
		sql_delete := `
			DELETE FROM
			` + configs.DB_tbl_mst_category + ` 
			WHERE idcatenews=? 
		`
		stmt_delete, e_delete := con.PrepareContext(ctx, sql_delete)
		helpers.ErrorCheck(e_delete)
		defer stmt_delete.Close()
		rec_delete, e_delete := stmt_delete.ExecContext(ctx, idrecord)

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
