package models

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/isbpanel_backend/configs"
	"github.com/nikitamirzani323/isbpanel_backend/db"
	"github.com/nikitamirzani323/isbpanel_backend/entities"
	"github.com/nikitamirzani323/isbpanel_backend/helpers"
)

func Fetch_adminruleHome() (helpers.Response, error) {
	var obj entities.Model_adminruleall
	var arraobj []entities.Model_adminruleall
	var res helpers.Response
	msg := "Error"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			idadmin , ruleadmingroup 
			FROM ` + configs.DB_tbl_admingroup + ` 
			ORDER BY idadmin ASC  
		`

	row, err := con.QueryContext(ctx, sql_select)

	var no int = 0
	helpers.ErrorCheck(err)
	for row.Next() {
		no += 1
		var (
			idadmin_db, ruleadmingroup_db string
		)

		err = row.Scan(&idadmin_db, &ruleadmingroup_db)

		helpers.ErrorCheck(err)

		obj.Idadmin = idadmin_db
		obj.Ruleadmingroup = ruleadmingroup_db
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
func Save_adminrule(admin, idadmin, rule, sData string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDB(configs.DB_tbl_admingroup, "idadmin ", idadmin)
		if !flag {
			sql_insert := `
			insert into
			` + configs.DB_tbl_admingroup + ` (
				idadmin 
			) values (
				?
			)
		`
			stmt_insert, e_insert := con.PrepareContext(ctx, sql_insert)
			helpers.ErrorCheck(e_insert)
			defer stmt_insert.Close()
			res_newpasaran, e_newpasaran := stmt_insert.ExecContext(ctx, idadmin)
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
		sql_update2 := `
				UPDATE 
				` + configs.DB_tbl_admingroup + `   
				SET ruleadmingroup =?
				WHERE idadmin  =? 
			`
		stmt_admin, e := con.PrepareContext(ctx, sql_update2)
		helpers.ErrorCheck(e)
		rec_admin, e_admin := stmt_admin.ExecContext(
			ctx,
			rule,
			idadmin)
		helpers.ErrorCheck(e_admin)

		update_admin, e_admin := rec_admin.RowsAffected()
		helpers.ErrorCheck(e_admin)

		defer stmt_admin.Close()
		if update_admin > 0 {
			flag = true
			msg = "Succes"
			log.Println("Update tbl_adminrule Success")
		} else {
			log.Println("Update tbl_adminrule failed")
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
