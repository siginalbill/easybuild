package models

type Db struct {
	User string `json:"user"`
	Pass string `json:"pass"`
	Data string `json:"data"`
	Conn conn `json:"conn"`
}

type conn struct {
	Type string `json:"type"`
	Ip string `json:"ip"`
	Port string `json:"port"`
	DataBase string `json:"data_base"`
}
