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
	"github.com/nleeper/goment"
)

func Fetch_pasaranHome() (helpers.Response, error) {
	var obj entities.Model_pasaran
	var arraobj []entities.Model_pasaran
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			idpasarantogel , nmpasarantogel, 
			urlpasaran , pasarandiundi, jamjadwal, 
			createpasarantogel, COALESCE(createdatepasarantogel,""), updatepasarantogel, COALESCE(updatedatepasarantogel,"")  
			FROM ` + configs.DB_tbl_mst_pasaran + ` 
			ORDER BY nmpasarantogel DESC 
		`

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idpasarantogel_db, nmpasarantogel_db, urlpasaran_db, pasarandiundi_db, jamjadwal_db                string
			createpasarantogel_db, createdatepasarantogel_db, updatepasarantogel_db, updatedatepasarantogel_db string
		)

		err = row.Scan(
			&idpasarantogel_db, &nmpasarantogel_db, &urlpasaran_db, &pasarandiundi_db, &jamjadwal_db,
			&createpasarantogel_db, &createdatepasarantogel_db, &updatepasarantogel_db, &updatedatepasarantogel_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		if createpasarantogel_db != "" {
			create = createpasarantogel_db + ", " + createdatepasarantogel_db
		}
		if updatepasarantogel_db != "" {
			update = updatepasarantogel_db + ", " + updatedatepasarantogel_db
		}

		var (
			datekeluaran_db, nomorkeluaran_db string
		)
		sql_selectpasaran := `SELECT 
			datekeluaran , nomorkeluaran
			FROM ` + configs.DB_tbl_trx_keluaran + ` 
			WHERE idpasarantogel = ? 
			ORDER BY datekeluaran DESC LIMIT 1
		`
		row_keluaran := con.QueryRowContext(ctx, sql_selectpasaran, idpasarantogel_db)
		switch e_keluaran := row_keluaran.Scan(&datekeluaran_db, &nomorkeluaran_db); e_keluaran {
		case sql.ErrNoRows:
		case nil:
		default:
			helpers.ErrorCheck(e_keluaran)
		}

		obj.Pasaran_id = idpasarantogel_db
		obj.Pasaran_name = nmpasarantogel_db
		obj.Pasaran_url = urlpasaran_db
		obj.Pasaran_diundi = pasarandiundi_db
		obj.Pasaran_jamjadwal = jamjadwal_db
		obj.Pasaran_keluaran = datekeluaran_db + " - " + nomorkeluaran_db
		obj.Pasaran_create = create
		obj.Pasaran_update = update
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
func Save_pasaran(admin, idrecord, nmpasarantogel, urlpasaran, pasarandiundi, jamjadwal, sData string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDB(configs.DB_tbl_mst_pasaran, "idpasarantogel", idrecord)
		if !flag {
			sql_insert := `
				insert into
				` + configs.DB_tbl_mst_pasaran + ` (
					idpasarantogel , nmpasarantogel, urlpasaran, pasarandiundi, jamjadwal 
					createpasarantogel, createdatepasarantogel
				) values (
					? ,?, ?, ?, ?,
					?, ?
				)
			`
			stmt_insert, e_insert := con.PrepareContext(ctx, sql_insert)
			helpers.ErrorCheck(e_insert)
			defer stmt_insert.Close()
			res_newrecord, e_newrecord := stmt_insert.ExecContext(
				ctx,
				idrecord, nmpasarantogel, urlpasaran, pasarandiundi, jamjadwal,
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
			msg = "Duplicate Entry"
		}
	} else {
		sql_update := `
				UPDATE 
				` + configs.DB_tbl_mst_pasaran + `  
				SET nmpasarantogel=?,urlpasaran=?,pasarandiundi=?, jamjadwal=?, 
				updatepasarantogel=?, updatedatepasarantogel=? 
				WHERE idpasarantogel =? 
			`
		stmt_record, e := con.PrepareContext(ctx, sql_update)
		helpers.ErrorCheck(e)
		rec_record, e_record := stmt_record.ExecContext(
			ctx,
			nmpasarantogel, urlpasaran, pasarandiundi, jamjadwal,
			admin,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			idrecord)
		helpers.ErrorCheck(e_record)
		update_record, e_record := rec_record.RowsAffected()
		helpers.ErrorCheck(e_record)

		defer stmt_record.Close()
		if update_record > 0 {
			flag = true
			msg = "Succes"
			log.Printf("Update PASARAN Success : %s\n", idrecord)
		} else {
			log.Println("Update PASARAN failed")
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
func Fetch_keluaran(idpasaran string) (helpers.Response, error) {
	var obj entities.Model_keluaran
	var arraobj []entities.Model_keluaran
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			idtrxkeluaran , datekeluaran, 
			periodekeluaran , nomorkeluaran
			FROM ` + configs.DB_tbl_trx_keluaran + ` 
			WHERE idpasarantogel=? 
			ORDER BY datekeluaran DESC LIMIT 365 
		`

	row, err := con.QueryContext(ctx, sql_select, idpasaran)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idtrxkeluaran_db, periodekeluaran_db int
			datekeluaran_db, nomorkeluaran_db    string
		)

		err = row.Scan(
			&idtrxkeluaran_db, &datekeluaran_db, &periodekeluaran_db, &nomorkeluaran_db)

		helpers.ErrorCheck(err)

		obj.Keluaran_id = idtrxkeluaran_db
		obj.Keluaran_tanggal = datekeluaran_db
		obj.Keluaran_periode = idpasaran + "-" + strconv.Itoa(periodekeluaran_db)
		obj.Keluaran_nomor = nomorkeluaran_db
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
func Save_keluaran(admin, idpasaran, tanggal, nomor string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	flag = CheckDB(configs.DB_tbl_trx_keluaran, "datekeluaran", tanggal)
	if !flag {
		sql_insert := `
			insert into
			` + configs.DB_tbl_trx_keluaran + ` (
				idtrxkeluaran , idpasarantogel, datekeluaran, periodekeluaran, nomorkeluaran, 
				createkeluaran, createdatekeluaran
			) values (
				? ,?, ?, ?, ?,
				?, ?
			)
		`
		stmt_insert, e_insert := con.PrepareContext(ctx, sql_insert)
		helpers.ErrorCheck(e_insert)
		defer stmt_insert.Close()
		field_column := configs.DB_tbl_trx_keluaran + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		field_column_periode := idpasaran + "-" + tglnow.Format("YYYY")
		idperiode_counter := Get_counter(field_column_periode)
		res_newrecord, e_newrecord := stmt_insert.ExecContext(
			ctx,
			tglnow.Format("YY")+strconv.Itoa(idrecord_counter),
			idpasaran, tanggal, idperiode_counter, nomor,
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
		msg = "Duplicate Entry"
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
func Delete_keluaran(admin, idpasaran string, idtrxkeluaran int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	flag := false

	flag = CheckDB(configs.DB_tbl_trx_keluaran, "idtrxkeluaran", strconv.Itoa(idtrxkeluaran))
	if flag {
		sql_delete := `
			DELETE FROM
			` + configs.DB_tbl_trx_keluaran + ` 
			WHERE idtrxkeluaran=? AND idpasarantogel=? 
		`
		stmt_delete, e_delete := con.PrepareContext(ctx, sql_delete)
		helpers.ErrorCheck(e_delete)
		defer stmt_delete.Close()
		rec_delete, e_delete := stmt_delete.ExecContext(ctx, idtrxkeluaran, idpasaran)

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
