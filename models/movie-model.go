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

func Fetch_movieHome(search string) (helpers.Response, error) {
	var obj entities.Model_movie
	var arraobj []entities.Model_movie
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := ""
	sql_select += ""
	sql_select += "SELECT "
	sql_select += "movieid , movietitle, description, movietype, "
	sql_select += "rating , imdb, year, views, enabled, "
	sql_select += "createmovie, COALESCE(createdatemovie,''), updatemovie, COALESCE(updatedatemovie,'') "
	sql_select += "FROM " + configs.DB_tbl_trx_movie + "  "
	if search == "" {
		sql_select += "ORDER BY createdatemovie DESC  LIMIT 2000  "
	} else {
		sql_select += "WHERE movietitle '%" + search + "%' "
		sql_select += "OR movietitle LIKE '%" + search + "%' "
		sql_select += "ORDER BY createdatemovie DESC  LIMIT 2000  "
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			movieid_db, year_db, views_db, enabled_db                              int
			rating_db, imdb_db                                                     float32
			movietitle_db, movietype_db, description_db                            string
			createmovie_db, createdatemovie_db, updatemovie_db, updatedatemovie_db string
		)

		err = row.Scan(
			&movieid_db, &movietitle_db, &description_db, &movietype_db, &rating_db, &imdb_db, &year_db, &views_db, &enabled_db,
			&createmovie_db, &createdatemovie_db, &updatemovie_db, &updatedatemovie_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		if createmovie_db != "" {
			create = createmovie_db + ", " + createdatemovie_db
		}
		if updatemovie_db != "" {
			update = updatemovie_db + ", " + updatedatemovie_db
		}

		obj.Movie_id = movieid_db
		obj.Movie_type = movietype_db
		obj.Movie_title = movietitle_db
		obj.Movie_descp = description_db
		obj.Movie_year = year_db
		obj.Movie_rating = rating_db
		obj.Movie_imdb = imdb_db
		obj.Movie_view = views_db
		obj.Movie_status = enabled_db
		obj.Movie_create = create
		obj.Movie_update = update
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
