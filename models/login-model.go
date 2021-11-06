package models

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/nikitamirzani323/isbpanel_backend/configs"
	"github.com/nikitamirzani323/isbpanel_backend/db"
	"github.com/nikitamirzani323/isbpanel_backend/helpers"
	"github.com/nleeper/goment"
)

func Login_Model(username, password, ipaddress, timezone string) (bool, int, error) {
	con := db.CreateCon()
	ctx := context.Background()
	flag := false
	tglnow, _ := goment.New()
	var passwordDB, idadminDB string
	var idruleadminDB int
	sql_select := `
			SELECT
			password, idadmin    
			FROM ` + configs.DB_tbl_admin + ` 
			WHERE username  = ?
			AND statuslogin = "Y" 
		`

	row := con.QueryRowContext(ctx, sql_select, username)
	switch e := row.Scan(&passwordDB, &idadminDB); e {
	case sql.ErrNoRows:
		return false, 0, errors.New("Username and Password Not Found")
	case nil:
		flag = true
	default:
		return false, 0, errors.New("Username and Password Not Found")
	}

	hashpass := helpers.HashPasswordMD5(password)
	log.Println("Password : " + hashpass)
	log.Println("Hash : " + passwordDB)
	if hashpass != passwordDB {
		return false, 0, nil
	}

	if flag {
		sql_update := `
			UPDATE ` + configs.DB_tbl_admin + ` 
			SET lastlogin=?, ipaddress=? , timezone=?, 
			updateadmin=?,  updatedateadmin=?  
			WHERE username  = ? 
			AND statuslogin = "Y" 
		`
		rows_update, err_update := con.PrepareContext(ctx, sql_update)
		helpers.ErrorCheck(err_update)
		res_update, err_update := rows_update.ExecContext(ctx,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			ipaddress,
			timezone,
			username,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			username)
		helpers.ErrorCheck(err_update)
		update, e := res_update.RowsAffected()
		helpers.ErrorCheck(e)
		if update > 0 {
			flag = true
			log.Println("LOGIN Data Berhasil di update")
		}
	}

	return true, idruleadminDB, nil
}
