package models

import (
	"context"
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/buger/jsonparser"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/isbpanel_backend/configs"
	"github.com/nikitamirzani323/isbpanel_backend/db"
	"github.com/nikitamirzani323/isbpanel_backend/entities"
	"github.com/nikitamirzani323/isbpanel_backend/helpers"
	"github.com/nleeper/goment"
)

func Fetch_movieHome(search string, page int) (helpers.Responsemovie, error) {
	var obj entities.Model_movie
	var arraobj []entities.Model_movie
	var res helpers.Responsemovie
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	perpage := 50
	totalrecord := 0
	offset := 0
	if page < 2 {
		offset = 0
	} else {
		offset = perpage * page
	}

	sql_selectcount := `SELECT 
		COUNT(movieid) as totalmovie 
		FROM ` + configs.DB_tbl_trx_movie + ` 
	`
	row_selectcount := con.QueryRowContext(ctx, sql_selectcount)
	switch e_selectcount := row_selectcount.Scan(&totalrecord); e_selectcount {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e_selectcount)
	}
	sql_select := ""
	sql_select += ""
	sql_select += "SELECT "
	sql_select += "movieid , movietitle, description, movietype, "
	sql_select += "rating , imdb, year, views, enabled, COALESCE(posted_id,0),label, urlthumbnail, slug,   "
	sql_select += "createmovie, COALESCE(createdatemovie,''), updatemovie, COALESCE(updatedatemovie,'') "
	sql_select += "FROM " + configs.DB_tbl_trx_movie + "  "
	if search == "" {
		sql_select += "ORDER BY createdatemovie DESC  LIMIT " + strconv.Itoa(offset) + " , " + strconv.Itoa(perpage)
	} else {
		sql_select += "WHERE movietitle LIKE '%" + search + "%' "
		sql_select += "ORDER BY createdatemovie DESC LIMIT " + strconv.Itoa(perpage)
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			movieid_db, year_db, views_db, enabled_db, posted_id_db                         int
			rating_db, imdb_db                                                              float32
			movietitle_db, movietype_db, description_db, label_db, urlthumbnail_db, slug_db string
			createmovie_db, createdatemovie_db, updatemovie_db, updatedatemovie_db          string
		)

		err = row.Scan(
			&movieid_db, &movietitle_db, &description_db, &movietype_db, &rating_db, &imdb_db, &year_db, &views_db,
			&enabled_db, &posted_id_db, &label_db, &urlthumbnail_db, &slug_db,
			&createmovie_db, &createdatemovie_db, &updatemovie_db, &updatedatemovie_db)

		helpers.ErrorCheck(err)
		status := "HIDE"
		statuscss := configs.STATUS_CANCEL
		create := ""
		update := ""
		if createmovie_db != "" {
			create = createmovie_db + ", " + createdatemovie_db
		}
		if updatemovie_db != "" {
			update = updatemovie_db + ", " + updatedatemovie_db
		}
		if enabled_db > 0 {
			status = "SHOW"
			statuscss = configs.STATUS_RUNNING
		}
		//GENRE
		var objmoviegenre entities.Model_moviegenre
		var arraobjmoviegenre []entities.Model_moviegenre
		sql_selectmoviegenre := `SELECT 
			A.idgenre, B.nmgenre 
			FROM ` + configs.DB_tbl_trx_moviegenre + ` as A 
			JOIN ` + configs.DB_tbl_mst_moviegenre + ` as B ON B.idgenre = A.idgenre 
			WHERE A.movieid = ?   
		`
		row_moviegenre, err := con.QueryContext(ctx, sql_selectmoviegenre, movieid_db)
		helpers.ErrorCheck(err)
		for row_moviegenre.Next() {
			var (
				idgenre_db int
				nmgenre_db string
			)
			err = row_moviegenre.Scan(&idgenre_db, &nmgenre_db)
			objmoviegenre.Moviegenre_id = idgenre_db
			objmoviegenre.Moviegenre_name = nmgenre_db
			arraobjmoviegenre = append(arraobjmoviegenre, objmoviegenre)
		}
		//SOURCE
		var objmoviesource entities.Model_moviesource
		var arraobjmoviesource []entities.Model_moviesource
		if movietype_db == "movie" {
			sql_selectmoviesource := `SELECT 
				id, url 
				FROM ` + configs.DB_tbl_mst_moviesource + ` 
				WHERE poster_id = ?   
			`
			row_moviesource, err := con.QueryContext(ctx, sql_selectmoviesource, movieid_db)
			helpers.ErrorCheck(err)
			nosource := 0
			for row_moviesource.Next() {
				nosource = nosource + 1
				var (
					id_db  int
					url_db string
				)
				err = row_moviesource.Scan(&id_db, &url_db)
				objmoviesource.Moviesource_id = id_db
				objmoviesource.Moviesource_stream = "STREAM-" + strconv.Itoa(nosource)
				objmoviesource.Moviesource_url = url_db
				arraobjmoviesource = append(arraobjmoviesource, objmoviesource)
			}
		}

		path_image := ""
		if urlthumbnail_db == "" {
			poster_image, poster_extension := _GetMedia(posted_id_db)
			path_image = "https://duniafilm.b-cdn.net/uploads/cache/poster_thumb/uploads/" + poster_extension + "/" + poster_image
		} else {
			path_image = urlthumbnail_db
		}

		obj.Movie_id = movieid_db
		obj.Movie_date = createdatemovie_db
		obj.Movie_type = movietype_db
		obj.Movie_title = movietitle_db
		obj.Movie_label = label_db
		obj.Movie_slug = slug_db
		obj.Movie_descp = description_db
		obj.Movie_thumbnail = path_image
		obj.Movie_year = year_db
		obj.Movie_rating = rating_db
		obj.Movie_imdb = imdb_db
		obj.Movie_view = views_db
		obj.Movie_genre = arraobjmoviegenre
		obj.Movie_source = arraobjmoviesource
		obj.Movie_status = status
		obj.Movie_statuscss = statuscss
		obj.Movie_create = create
		obj.Movie_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Perpage = perpage
	res.Totalrecord = totalrecord
	res.Time = time.Since(start).String()

	return res, nil
}
func Save_movie(admin, name, label, slug, tipemovie, descp, urlthum, listgenre, listsource, sdata string, idrecord, year, status int, imdb float32) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	temp_idrecord := ""
	if sdata == "New" {
		sql_insert := `
			insert into
			` + configs.DB_tbl_trx_movie + ` (
				movieid , movietitle, label, movietype, description, imdb, year, slug, enabled, urlthumbnail,    
				createmovie, createdatemovie
			) values (
				?,?,?,?,?,?,?,?,?,?, 
				?, ?
			)
		`
		stmt_insert, e_insert := con.PrepareContext(ctx, sql_insert)
		helpers.ErrorCheck(e_insert)
		defer stmt_insert.Close()
		field_column := configs.DB_tbl_trx_movie + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		temp_idrecord = tglnow.Format("YY") + tglnow.Format("MM") + strconv.Itoa(idrecord_counter)
		res_newrecord, e_newrecord := stmt_insert.ExecContext(
			ctx,
			temp_idrecord,
			name, label, tipemovie, descp, imdb, year, slug, status, urlthum,
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
			//GENRE
			json := []byte(listgenre)
			jsonparser.ArrayEach(json, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				movie_genre_id, _ := jsonparser.GetInt(value, "movie_genre_id")
				movie_genre_name, _ := jsonparser.GetString(value, "movie_genre_name")

				log.Printf("%d - %s", int(movie_genre_id), movie_genre_name)
				sql_insertgenre := `
				insert into
					` + configs.DB_tbl_trx_moviegenre + ` (
						idmoviegenre , movieid, idgenre
					) values (
						?,?,?
					)
				`
				stmt_insertgenre, e_insertgenre := con.PrepareContext(ctx, sql_insertgenre)
				helpers.ErrorCheck(e_insertgenre)
				defer stmt_insertgenre.Close()
				field_column := configs.DB_tbl_trx_moviegenre + tglnow.Format("YYYY")
				idrecord_counter := Get_counter(field_column)
				res_newrecordgenre, e_newrecordgenre := stmt_insertgenre.ExecContext(
					ctx,
					tglnow.Format("YY")+tglnow.Format("MM")+strconv.Itoa(idrecord_counter),
					temp_idrecord, movie_genre_id)
				helpers.ErrorCheck(e_newrecordgenre)
				insertgenre, e := res_newrecordgenre.RowsAffected()
				helpers.ErrorCheck(e)
				if insertgenre > 0 {
					flag = true
					msg = "Succes"
					log.Println("Data GENERE Berhasil di save")
				}
			})

			//SOURCE
			jsonsource := []byte(listsource)
			jsonparser.ArrayEach(jsonsource, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				movie_source_id, _ := jsonparser.GetInt(value, "movie_source_id")
				movie_source_name, _ := jsonparser.GetString(value, "movie_source_name")

				log.Printf("%d - %s", int(movie_source_id), movie_source_name)
				sql_insertsource := `
				insert into
					` + configs.DB_tbl_mst_moviesource + ` (
						id , poster_id, type, url
					) values (
						?,?,?,?
					)
				`
				stmt_insertsource, e_insertsource := con.PrepareContext(ctx, sql_insertsource)
				helpers.ErrorCheck(e_insertsource)
				defer stmt_insertsource.Close()
				field_column := configs.DB_tbl_mst_moviesource + tglnow.Format("YYYY")
				idrecord_counter := Get_counter(field_column)
				res_newrecordsource, e_newrecordsource := stmt_insertsource.ExecContext(
					ctx,
					tglnow.Format("YY")+tglnow.Format("MM")+strconv.Itoa(idrecord_counter),
					temp_idrecord, "embed", movie_source_name)
				helpers.ErrorCheck(e_newrecordsource)
				insertsource, e := res_newrecordsource.RowsAffected()
				helpers.ErrorCheck(e)
				if insertsource > 0 {
					flag = true
					msg = "Succes"
					log.Println("Data SOURCE Berhasil di save")
				}
			})
		}
	} else {
		sql_update := `
			UPDATE 
			` + configs.DB_tbl_trx_movie + ` 
			SET movietitle=?, description=?, slug=?, urlthumbnail=?, 
			updatemovie=?, updatedatemovie=? 
			WHERE movieid=? 
		`
		stmt_update, e_update := con.PrepareContext(ctx, sql_update)
		helpers.ErrorCheck(e_update)
		defer stmt_update.Close()
		res_newrecord, e_newrecord := stmt_update.ExecContext(
			ctx,
			name, descp, slug, urlthum,
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
		if flag {
			//DELETE GENRE
			stmt_genre_delete, e_genre_delete := con.PrepareContext(ctx, `
					DELETE FROM  
					`+configs.DB_tbl_trx_moviegenre+`   
					WHERE movieid=?  
			`)

			helpers.ErrorCheck(e_genre_delete)
			rec_genre_delete, e_rec_genre_delete := stmt_genre_delete.ExecContext(ctx, idrecord)
			helpers.ErrorCheck(e_rec_genre_delete)

			affect_genre_delete, err_affer_genre_delete := rec_genre_delete.RowsAffected()
			helpers.ErrorCheck(err_affer_genre_delete)

			defer stmt_genre_delete.Close()
			if affect_genre_delete > 0 {
				log.Printf("Delete genre : %d\n", idrecord)
			} else {
				log.Printf("Delete genre : %d Failed\n ", idrecord)
			}

			//DELETE SOURCE
			stmt_source_delete, e_source_delete := con.PrepareContext(ctx, `
					DELETE FROM  
					`+configs.DB_tbl_mst_moviesource+`   
					WHERE poster_id=?  
			`)

			helpers.ErrorCheck(e_source_delete)
			rec_source_delete, e_rec_source_delete := stmt_source_delete.ExecContext(ctx, idrecord)
			helpers.ErrorCheck(e_rec_source_delete)

			affect_source_delete, err_affer_source_delete := rec_source_delete.RowsAffected()
			helpers.ErrorCheck(err_affer_source_delete)

			defer stmt_source_delete.Close()
			if affect_source_delete > 0 {
				log.Printf("Delete source : %d\n", idrecord)
			} else {
				log.Printf("Delete source : %d Failed\n ", idrecord)
			}

			//GENRE
			json := []byte(listgenre)
			jsonparser.ArrayEach(json, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				movie_genre_id, _ := jsonparser.GetInt(value, "movie_genre_id")
				movie_genre_name, _ := jsonparser.GetString(value, "movie_genre_name")

				log.Printf("%d - %s", int(movie_genre_id), movie_genre_name)
				sql_insertgenre := `
					insert into
						` + configs.DB_tbl_trx_moviegenre + ` (
							idmoviegenre , movieid, idgenre
						) values (
							?,?,?
						)
					`
				stmt_insertgenre, e_insertgenre := con.PrepareContext(ctx, sql_insertgenre)
				helpers.ErrorCheck(e_insertgenre)
				defer stmt_insertgenre.Close()
				field_column := configs.DB_tbl_trx_moviegenre + tglnow.Format("YYYY")
				idrecord_counter := Get_counter(field_column)
				res_newrecordgenre, e_newrecordgenre := stmt_insertgenre.ExecContext(
					ctx,
					tglnow.Format("YY")+tglnow.Format("MM")+strconv.Itoa(idrecord_counter),
					idrecord, movie_genre_id)
				helpers.ErrorCheck(e_newrecordgenre)
				insertgenre, e := res_newrecordgenre.RowsAffected()
				helpers.ErrorCheck(e)
				if insertgenre > 0 {
					flag = true
					msg = "Succes"
					log.Println("Data GENERE Berhasil di save")
				}
			})

			//SOURCE
			jsonsource := []byte(listsource)
			jsonparser.ArrayEach(jsonsource, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				movie_source_id, _ := jsonparser.GetInt(value, "movie_source_id")
				movie_source_name, _ := jsonparser.GetString(value, "movie_source_name")

				log.Printf("%d - %s", int(movie_source_id), movie_source_name)
				sql_insertsource := `
				insert into
					` + configs.DB_tbl_mst_moviesource + ` (
						id , poster_id, type, url
					) values (
						?,?,?,?
					)
				`
				stmt_insertsource, e_insertsource := con.PrepareContext(ctx, sql_insertsource)
				helpers.ErrorCheck(e_insertsource)
				defer stmt_insertsource.Close()
				field_column := configs.DB_tbl_mst_moviesource + tglnow.Format("YYYY")
				idrecord_counter := Get_counter(field_column)
				res_newrecordsource, e_newrecordsource := stmt_insertsource.ExecContext(
					ctx,
					tglnow.Format("YY")+tglnow.Format("MM")+strconv.Itoa(idrecord_counter),
					idrecord, "embed", movie_source_name)
				helpers.ErrorCheck(e_newrecordsource)
				insertsource, e := res_newrecordsource.RowsAffected()
				helpers.ErrorCheck(e)
				if insertsource > 0 {
					flag = true
					msg = "Succes"
					log.Println("Data SOURCE Berhasil di save")
				}
			})
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
func Delete_movie(admin string, idrecord int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	flag := false

	flag = CheckDB(configs.DB_tbl_trx_movie, "movieid", strconv.Itoa(idrecord))

	if flag {
		//DELETE GENRE
		sql_deletegenre := `
			DELETE FROM
			` + configs.DB_tbl_trx_moviegenre + ` 
			WHERE movieid=? 
		`
		stmt_deletegenre, e_deletegenre := con.PrepareContext(ctx, sql_deletegenre)
		helpers.ErrorCheck(e_deletegenre)
		defer stmt_deletegenre.Close()
		rec_deletegenre, e_deletegenre := stmt_deletegenre.ExecContext(ctx, idrecord)

		helpers.ErrorCheck(e_deletegenre)
		deletegenre, egenre := rec_deletegenre.RowsAffected()
		helpers.ErrorCheck(egenre)
		if deletegenre > 0 {
			flag = true
			msg = "Succes"
			log.Println("Data Genre Berhasil di delete")
		}

		//DELETE SOURCE
		sql_deletesource := `
			DELETE FROM
			` + configs.DB_tbl_mst_moviesource + ` 
			WHERE poster_id=? 
		`
		stmt_deletesource, e_deletesource := con.PrepareContext(ctx, sql_deletesource)
		helpers.ErrorCheck(e_deletesource)
		defer stmt_deletesource.Close()
		rec_deletesource, e_deletesource := stmt_deletesource.ExecContext(ctx, idrecord)

		helpers.ErrorCheck(e_deletesource)
		deletesource, esource := rec_deletesource.RowsAffected()
		helpers.ErrorCheck(esource)
		if deletesource > 0 {
			flag = true
			msg = "Succes"
			log.Println("Data Source Berhasil di delete")
		}

		//DELETE MOVIE
		sql_delete := `
			DELETE FROM
			` + configs.DB_tbl_trx_movie + ` 
			WHERE movieid=? 
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
			log.Println("Data Movie Berhasil di delete")
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
func Fetch_genre() (helpers.Response, error) {
	var obj entities.Model_genre
	var arraobj []entities.Model_genre
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			idgenre , nmgenre, genredisplay, 
			creategenre, COALESCE(createdategenre,""), updategenre, COALESCE(updatedategenre,"")  
			FROM ` + configs.DB_tbl_mst_moviegenre + ` 
			ORDER BY genredisplay ASC   
		`

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idgenre_db, genredisplay_db                                            int
			nmgenre_db                                                             string
			creategenre_db, createdategenre_db, updategenre_db, updatedategenre_db string
		)

		err = row.Scan(
			&idgenre_db, &nmgenre_db, &genredisplay_db,
			&creategenre_db, &createdategenre_db, &updategenre_db, &updatedategenre_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		if creategenre_db != "" {
			create = creategenre_db + ", " + createdategenre_db
		}
		if updategenre_db != "" {
			update = updategenre_db + ", " + updatedategenre_db
		}

		obj.Genre_id = idgenre_db
		obj.Genre_name = nmgenre_db
		obj.Genre_display = genredisplay_db
		obj.Genre_create = create
		obj.Genre_update = update
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
func Save_genre(admin, name, sdata string, idrecord, display int) (helpers.Response, error) {
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
			` + configs.DB_tbl_mst_moviegenre + ` (
				idgenre , nmgenre, genredisplay,  
				creategenre, createdategenre
			) values (
				? ,?, ?, 
				?, ?
			)
		`
		stmt_insert, e_insert := con.PrepareContext(ctx, sql_insert)
		helpers.ErrorCheck(e_insert)
		defer stmt_insert.Close()
		field_column := configs.DB_tbl_mst_moviegenre + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		res_newrecord, e_newrecord := stmt_insert.ExecContext(
			ctx,
			tglnow.Format("YY")+strconv.Itoa(idrecord_counter),
			name, display,
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
			` + configs.DB_tbl_mst_moviegenre + ` 
			SET nmgenre=?, genredisplay=?, 
			updategenre=?, updatedategenre=? 
			WHERE idgenre=? 
		`
		stmt_update, e_update := con.PrepareContext(ctx, sql_update)
		helpers.ErrorCheck(e_update)
		defer stmt_update.Close()
		res_newrecord, e_newrecord := stmt_update.ExecContext(
			ctx,
			name, display,
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
func Delete_genre(admin string, idrecord int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	flag := false
	flag_movie := false

	flag = CheckDB(configs.DB_tbl_mst_moviegenre, "idgenre", strconv.Itoa(idrecord))
	flag_movie = CheckDB(configs.DB_tbl_trx_moviegenre, "idgenre", strconv.Itoa(idrecord))
	if flag {
		if flag_movie {
			sql_delete := `
				DELETE FROM
				` + configs.DB_tbl_mst_moviegenre + ` 
				WHERE idgenre=? 
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
			msg = "Cannot Delete"
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
func _GetMedia(idrecord int) (string, string) {
	con := db.CreateCon()
	ctx := context.Background()
	url := ""
	extension := ""

	sql_select := `SELECT
		url, extension   
		FROM ` + configs.DB_tbl_mst_mediatable + `  
		WHERE idmediatable = ? 
	`
	row := con.QueryRowContext(ctx, sql_select, idrecord)
	switch e := row.Scan(&url, &extension); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}
	return url, extension
}
