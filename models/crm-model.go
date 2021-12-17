package models

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/isbpanel_backend/configs"
	"github.com/nikitamirzani323/isbpanel_backend/db"
	"github.com/nikitamirzani323/isbpanel_backend/entities"
	"github.com/nikitamirzani323/isbpanel_backend/helpers"
)

func Fetch_crmisbtv(search string, page int) (helpers.Responsemovie, error) {
	var obj entities.Model_crmisbtv
	var arraobj []entities.Model_crmisbtv
	var res helpers.Responsemovie
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	perpage := 250
	totalrecord := 0
	offset := page
	sql_selectcount := ""
	sql_selectcount += ""
	sql_selectcount += "SELECT "
	sql_selectcount += "COUNT(username) as totalmember  "
	sql_selectcount += "FROM " + configs.DB_tbl_mst_user + "  "
	if search != "" {
		sql_selectcount += "WHERE username LIKE '%" + search + "%' "
		sql_selectcount += "OR username LIKE '%" + search + "%' "
	}

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
	sql_select += "username , nmuser, coderef, "
	sql_select += "point_in , point_out, statususer,  "
	sql_select += "COALESCE(lastlogin,''), createdateuser, COALESCE(updatedateuser,'') "
	sql_select += "FROM " + configs.DB_tbl_mst_user + "  "
	if search == "" {
		sql_select += "ORDER BY createdateuser DESC  LIMIT " + strconv.Itoa(offset) + " , " + strconv.Itoa(perpage)
	} else {
		sql_select += "WHERE username LIKE '%" + search + "%' "
		sql_select += "OR username LIKE '%" + search + "%' "
		sql_select += "ORDER BY createdateuser DESC  LIMIT " + strconv.Itoa(perpage)
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			point_in_db, point_out_db                          int
			username_db, nmuser_db, coderef_db, statususer_db  string
			lastlogin_db, createdateuser_db, updatedateuser_db string
		)

		err = row.Scan(
			&username_db, &nmuser_db, &coderef_db, &point_in_db, &point_out_db, &statususer_db,
			&lastlogin_db, &createdateuser_db, &updatedateuser_db)

		helpers.ErrorCheck(err)

		obj.Crmisbtv_username = username_db
		obj.Crmisbtv_name = nmuser_db
		obj.Crmisbtv_coderef = coderef_db
		obj.Crmisbtv_point = point_in_db - point_out_db
		obj.Crmisbtv_status = statususer_db
		obj.Crmisbtv_lastlogin = lastlogin_db
		obj.Crmisbtv_create = createdateuser_db
		obj.Crmisbtv_update = updatedateuser_db
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
func Fetch_crmduniafilm(search string, page int) (helpers.Responsemovie, error) {
	var obj entities.Model_crmduniafilm
	var arraobj []entities.Model_crmduniafilm
	var res helpers.Responsemovie
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	perpage := 250
	totalrecord := 0
	offset := page
	sql_selectcount := ""
	sql_selectcount += ""
	sql_selectcount += "SELECT "
	sql_selectcount += "COUNT(username) as totalmember  "
	sql_selectcount += "FROM " + configs.DB_VIEW_MEMBER_DUNIAFILM + "  "
	if search != "" {
		sql_selectcount += "WHERE username LIKE '%" + search + "%' "
		sql_selectcount += "OR name LIKE '%" + search + "%' "
	}

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
	sql_select += "username , name "
	sql_select += "FROM " + configs.DB_VIEW_MEMBER_DUNIAFILM + "  "
	if search == "" {
		sql_select += "ORDER BY username ASC  LIMIT " + strconv.Itoa(offset) + " , " + strconv.Itoa(perpage)
	} else {
		sql_select += "WHERE username LIKE '%" + search + "%' "
		sql_select += "OR name LIKE '%" + search + "%' "
		sql_select += "ORDER BY username ASC  LIMIT " + strconv.Itoa(perpage)
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			username_db, name_db string
		)

		err = row.Scan(&username_db, &name_db)

		helpers.ErrorCheck(err)

		obj.Crmduniafilm_username = username_db
		obj.Crmduniafilm_name = name_db
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
