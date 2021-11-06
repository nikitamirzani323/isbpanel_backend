package models

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/isbpanel_backend/configs"
	"github.com/nikitamirzani323/isbpanel_backend/db"
	"github.com/nikitamirzani323/isbpanel_backend/entities"
	"github.com/nikitamirzani323/isbpanel_backend/helpers"
	"github.com/nleeper/goment"
)

func Fetch_adminHome() (helpers.ResponseAdmin, error) {
	var obj entities.Model_admin
	var arraobj []entities.Model_admin
	var res helpers.ResponseAdmin
	msg := "Error"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			username , name, idadmin,
			statuslogin, lastlogin, joindate, 
			ipaddress, timezone  
			FROM ` + configs.DB_tbl_admin + ` 
			ORDER BY lastlogin DESC 
		`

	row, err := con.QueryContext(ctx, sql_select)

	var no int = 0
	helpers.ErrorCheck(err)
	for row.Next() {
		no += 1
		var (
			username_db, name_db, idadminlevel_db                                string
			statuslogin_db, lastlogin_db, joindate_db, ipaddress_db, timezone_db string
		)

		err = row.Scan(
			&username_db, &name_db, &idadminlevel_db,
			&statuslogin_db, &lastlogin_db, &joindate_db,
			&ipaddress_db, &timezone_db)

		helpers.ErrorCheck(err)
		if statuslogin_db == "Y" {
			statuslogin_db = "ACTIVE"
		}
		if lastlogin_db == "0000-00-00 00:00:00" {
			lastlogin_db = ""
		}
		obj.Username = username_db
		obj.Nama = name_db
		obj.Rule = idadminlevel_db
		obj.Joindate = joindate_db
		obj.Timezone = timezone_db
		obj.Lastlogin = lastlogin_db
		obj.LastIpaddress = ipaddress_db
		obj.Status = statuslogin_db
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	var objRule entities.Model_adminrule
	var arraobjRule []entities.Model_adminrule
	sql_listrule := `SELECT 
		idadmin 	
		FROM ` + configs.DB_tbl_admingroup + ` 
	`
	row_listrule, err_listrule := con.QueryContext(ctx, sql_listrule)

	helpers.ErrorCheck(err_listrule)
	for row_listrule.Next() {
		var (
			idruleadmin_db string
		)

		err = row_listrule.Scan(&idruleadmin_db)

		helpers.ErrorCheck(err)

		objRule.Idrule = idruleadmin_db
		arraobjRule = append(arraobjRule, objRule)
		msg = "Success"
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Listrule = arraobjRule
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_adminDetail(username string) (helpers.ResponseAdmin, error) {
	var obj entities.Model_adminsave
	var arraobj []entities.Model_adminsave
	var res helpers.ResponseAdmin
	msg := "Error"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()
	flag := true

	sql_detail := `SELECT 
		idadmin, name, statuslogin  
		createadmin, createdateadmin, updateadmin, updatedateadmin  
		FROM ` + configs.DB_tbl_admin + `
		WHERE username = ? 
	`
	var (
		idadmin_db, name_db, statuslogin_db                                    string
		createadmin_db, createdateadmin_db, updateadmin_db, updatedateadmin_db string
	)
	rows := con.QueryRowContext(ctx, sql_detail, username)

	switch err := rows.Scan(
		&idadmin_db, &name_db, &statuslogin_db,
		&createadmin_db, &createdateadmin_db, &updateadmin_db, &updatedateadmin_db); err {
	case sql.ErrNoRows:
		flag = false
	case nil:
		if createdateadmin_db == "0000-00-00 00:00:00" {
			createdateadmin_db = ""
		}
		if updatedateadmin_db == "0000-00-00 00:00:00" {
			updatedateadmin_db = ""
		}
		create := ""
		update := ""
		if createdateadmin_db != "" {
			create = createadmin_db + ", " + createdateadmin_db
		}
		if updateadmin_db != "" {
			create = updateadmin_db + ", " + updatedateadmin_db
		}

		obj.Username = username
		obj.Nama = name_db
		obj.Rule = idadmin_db
		obj.Status = statuslogin_db
		obj.Create = create
		obj.Update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	default:
		flag = false
		helpers.ErrorCheck(err)
	}

	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = arraobj
		res.Time = time.Since(start).String()
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = nil
		res.Time = time.Since(start).String()
	}

	return res, nil
}
func Save_adminHome(admin, username, password, nama, rule, status, sData string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDB(configs.DB_tbl_admin, "username", username)
		if !flag {
			sql_insert := `
			insert into
			` + configs.DB_tbl_admin + ` (
				username , password, idadmin, name, statuslogin, joindate, 
				createadmin, createdateadmin
			) values (
				?, ?, ?, ?, ?, ?, 
				?, ?
			)
		`
			stmt_insert, e_insert := con.PrepareContext(ctx, sql_insert)
			helpers.ErrorCheck(e_insert)
			defer stmt_insert.Close()
			hashpass := helpers.HashPasswordMD5(password)
			res_newpasaran, e_newpasaran := stmt_insert.ExecContext(
				ctx,
				username, hashpass,
				rule, nama, "Y",
				tglnow.Format("YYYY-MM-DD HH:mm:ss"),
				admin,
				tglnow.Format("YYYY-MM-DD HH:mm:ss"))
			helpers.ErrorCheck(e_newpasaran)
			insert, e := res_newpasaran.RowsAffected()
			helpers.ErrorCheck(e)
			if insert > 0 {
				flag = true
				msg = "Succes"
				log.Println("Data Berhasil di save")
			}
		} else {
			msg = "Duplicate Entry"
		}
	} else {
		if password == "" {
			sql_update := `
				UPDATE 
				` + configs.DB_tbl_admin + `  
				SET name =?, idadmin=?, statuslogin=?,  
				updateadmin=?, updatedateadmin=? 
				WHERE username =? 
			`
			stmt_admin, e := con.PrepareContext(ctx, sql_update)
			helpers.ErrorCheck(e)
			rec_admin, e_admin := stmt_admin.ExecContext(
				ctx,
				nama,
				rule,
				status,
				admin,
				tglnow.Format("YYYY-MM-DD HH:mm:ss"),
				username)
			helpers.ErrorCheck(e_admin)
			update_admin, e_admin := rec_admin.RowsAffected()
			helpers.ErrorCheck(e_admin)

			defer stmt_admin.Close()
			if update_admin > 0 {
				flag = true
				msg = "Succes"
				log.Printf("Update tbl_admin Success : %s\n", username)
			} else {
				log.Println("Update tbl_admin failed")
			}
		} else {
			hashpass := helpers.HashPasswordMD5(password)
			sql_update2 := `
				UPDATE 
				` + configs.DB_tbl_admin + `   
				SET name =?, password=?, idadmin=?, statuslogin=?,  
				updateadmin=?, updatedateadmin=? 
				WHERE username =? 
			`
			stmt_admin, e := con.PrepareContext(ctx, sql_update2)
			helpers.ErrorCheck(e)
			rec_admin, e_admin := stmt_admin.ExecContext(
				ctx,
				nama,
				hashpass,
				rule,
				status,
				admin,
				tglnow.Format("YYYY-MM-DD HH:mm:ss"),
				username)
			helpers.ErrorCheck(e_admin)

			update_admin, e_admin := rec_admin.RowsAffected()
			helpers.ErrorCheck(e_admin)

			defer stmt_admin.Close()
			if update_admin > 0 {
				flag = true
				msg = "Succes"
				log.Println("Update tbl_admin Success")
			} else {
				log.Println("Update tbl_admin failed")
			}
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
