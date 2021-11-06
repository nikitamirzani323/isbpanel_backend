package entities

type Model_pasaran struct {
	Pasaran_id        string `json:"pasaran_id"`
	Pasaran_name      string `json:"pasaran_name"`
	Pasaran_url       string `json:"pasaran_url"`
	Pasaran_diundi    string `json:"pasaran_diundi"`
	Pasaran_jamjadwal string `json:"pasaran_jamjadwal"`
	Pasaran_keluaran  string `json:"pasaran_keluaran"`
	Pasaran_create    string `json:"pasaran_create"`
	Pasaran_update    string `json:"pasaran_update"`
}
type Model_keluaran struct {
	Keluaran_id      int    `json:"keluaran_id"`
	Keluaran_tanggal string `json:"keluaran_tanggal"`
	Keluaran_periode string `json:"keluaran_periode"`
	Keluaran_nomor   string `json:"keluaran_nomor"`
}
type Controller_pasaransave struct {
	Sdata             string `json:"sdata" validate:"required"`
	Page              string `json:"page" validate:"required"`
	Pasaran_id        string `json:"pasaran_id"`
	Pasaran_name      string `json:"pasaran_name"`
	Pasaran_url       string `json:"pasaran_url"`
	Pasaran_diundi    string `json:"pasaran_diundi"`
	Pasaran_jamjadwal string `json:"pasaran_jamjadwal"`
}
type Controller_keluaran struct {
	Page       string `json:"page" validate:"required"`
	Pasaran_id string `json:"pasaran_id" validate:"required"`
}
type Controller_keluaransave struct {
	Sdata            string `json:"sdata" validate:"required"`
	Page             string `json:"page" validate:"required"`
	Pasaran_id       string `json:"pasaran_id"`
	Keluaran_tanggal string `json:"keluaran_tanggal"`
	Keluaran_nomor   string `json:"keluaran_nomor"`
}
type Controller_keluarandelete struct {
	Page        string `json:"page" validate:"required"`
	Pasaran_id  string `json:"pasaran_id"`
	Keluaran_id int    `json:"keluaran_id"`
}
