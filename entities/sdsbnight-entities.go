package entities

type Model_sdsbnight struct {
	Sdsbnight_id     int    `json:"sdsbnight_id"`
	Sdsbnight_date   string `json:"sdsbnight_date"`
	Sdsbnight_prize1 string `json:"sdsbnight_prize1"`
	Sdsbnight_prize2 string `json:"sdsbnight_prize2"`
	Sdsbnight_prize3 string `json:"sdsbnight_prize3"`
	Sdsbnight_create string `json:"sdsbnight_create"`
	Sdsbnight_update string `json:"sdsbnight_update"`
}
type Controller_sdsbnightsave struct {
	Sdata    string `json:"sdata" validate:"required"`
	Page     string `json:"page" validate:"required"`
	Idrecord int    `json:"idrecord"`
	Tanggal  string `json:"tanggal" validate:"required"`
}
type Controller_sdsbnightprizesave struct {
	Sdata    string `json:"sdata" validate:"required"`
	Page     string `json:"page" validate:"required"`
	Idrecord int    `json:"idrecord"`
	Tipe     string `json:"tipe" validate:"required"`
	Prize    string `json:"prize" validate:"required"`
}
type Responseredis_sdsbnight struct {
	Sdsbnight_id     int    `json:"sdsbnight_id"`
	Sdsbnight_date   string `json:"sdsbnight_date"`
	Sdsbnight_prize1 string `json:"sdsbnight_prize1"`
	Sdsbnight_prize2 string `json:"sdsbnight_prize2"`
	Sdsbnight_prize3 string `json:"sdsbnight_prize3"`
	Sdsbnight_create string `json:"sdsbnight_create"`
	Sdsbnight_update string `json:"sdsbnight_update"`
}
