package entities

type Model_crmisbtv struct {
	Crmisbtv_username  string `json:"crmisbtv_username"`
	Crmisbtv_name      string `json:"crmisbtv_name"`
	Crmisbtv_coderef   string `json:"crmisbtv_coderef"`
	Crmisbtv_point     int    `json:"crmisbtv_point"`
	Crmisbtv_status    string `json:"crmisbtv_status"`
	Crmisbtv_lastlogin string `json:"crmisbtv_lastlogin"`
	Crmisbtv_create    string `json:"crmisbtv_create"`
	Crmisbtv_update    string `json:"crmisbtv_update"`
}
type Model_crmduniafilm struct {
	Crmduniafilm_username string `json:"crmduniafilm_username"`
	Crmduniafilm_name     string `json:"crmduniafilm_name"`
}

type Controller_crmisbtv struct {
	Crmisbtv_search string `json:"crmisbtv_search"`
	Crmisbtv_page   int    `json:"crmisbtv_page"`
}
