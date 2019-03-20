package models

type Dic struct {
	Dics []dic_paths `json:"dics"`
	Root root `json:"root"`
}

type dic_paths struct {
	Name string `json:"name"`
	Describe string `json:"describe"`
}

type root struct {
	Name string `json:"name"`
	Describe string `json:"describe"`
}
