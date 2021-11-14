package models

import (
	"context"
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/isbpanel_backend/configs"
	"github.com/nikitamirzani323/isbpanel_backend/db"
	"github.com/nikitamirzani323/isbpanel_backend/entities"
	"github.com/nikitamirzani323/isbpanel_backend/helpers"
)

func Fetch_movieHome(search string) (helpers.Responsemovie, error) {
	var obj entities.Model_movie
	var arraobj []entities.Model_movie
	var res helpers.Responsemovie
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	perpage := 50
	totalrecord := 0

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
	log.Println("TOTALRECORD : ", totalrecord)
	sql_select := ""
	sql_select += ""
	sql_select += "SELECT "
	sql_select += "movieid , movietitle, description, movietype, "
	sql_select += "rating , imdb, year, views, enabled, posted_id, "
	sql_select += "createmovie, COALESCE(createdatemovie,''), updatemovie, COALESCE(updatedatemovie,'') "
	sql_select += "FROM " + configs.DB_tbl_trx_movie + "  "
	if search == "" {
		sql_select += "ORDER BY createdatemovie DESC LIMIT " + strconv.Itoa(perpage)
	} else {
		sql_select += "WHERE movietitle '%" + search + "%' "
		sql_select += "OR movietitle LIKE '%" + search + "%' "
		sql_select += "ORDER BY createdatemovie DESC  LIMIT " + strconv.Itoa(perpage)
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			movieid_db, year_db, views_db, enabled_db, posted_id_db                int
			rating_db, imdb_db                                                     float32
			movietitle_db, movietype_db, description_db                            string
			createmovie_db, createdatemovie_db, updatemovie_db, updatedatemovie_db string
		)

		err = row.Scan(
			&movieid_db, &movietitle_db, &description_db, &movietype_db, &rating_db, &imdb_db, &year_db, &views_db, &enabled_db, &posted_id_db,
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

		var objmoviegenre entities.Model_moviegenre
		var arraobjmoviegenre []entities.Model_moviegenre
		sql_selectmoviegenre := `SELECT 
			B.nmgenre 
			FROM ` + configs.DB_tbl_trx_moviegenre + ` as A 
			JOIN ` + configs.DB_tbl_mst_moviegenre + ` as B ON B.idgenre = A.idgenre 
			WHERE A.movieid = ?   
		`
		row_moviegenre, err := con.QueryContext(ctx, sql_selectmoviegenre, movieid_db)
		helpers.ErrorCheck(err)
		for row_moviegenre.Next() {
			var (
				nmgenre_db string
			)
			err = row_moviegenre.Scan(&nmgenre_db)
			objmoviegenre.Moviegenre_name = nmgenre_db
			arraobjmoviegenre = append(arraobjmoviegenre, objmoviegenre)
		}
		poster_image, poster_extension := _GetMedia(posted_id_db)
		path_image := "https://duniafilm.b-cdn.net/uploads/cache/poster_thumb/uploads/" + poster_extension + "/" + poster_image

		obj.Movie_id = movieid_db
		obj.Movie_type = movietype_db
		obj.Movie_title = movietitle_db
		obj.Movie_descp = description_db
		obj.Movie_thumbnail = path_image
		obj.Movie_year = year_db
		obj.Movie_rating = rating_db
		obj.Movie_imdb = imdb_db
		obj.Movie_view = views_db
		obj.Movie_genre = arraobjmoviegenre
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
