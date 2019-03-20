package utils

import (
	"../models"
	"fmt"
	"text/template"
	"os"
)

func CreateDBhelper() error {
	// 读取json文件
	JsonParse := NewJsonStruct()
	v := models.Db{}
	JsonParse.Load("./configs/db.json", &v)
	_, InputError := os.Open("./configs/db.json")
	if InputError != nil {
		return InputError
	}
	// 获取模板,模板文件名应与template一致
	t := template.New("dbhelper.template")
	t, err := t.ParseFiles("./source/dbhelper.template")
	if err != nil {
		return err
	}
	file, err1 := os.Create("./src/datasource/dbhelper.go");
	defer file.Close();
	if err1 != nil {
		return err1
	}
	err = t.Execute(file, v)
	if err != nil {
		return err
	}
	return nil
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
