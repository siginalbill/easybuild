package utils

import (
	"../models"
	"os"
	"text/template"
)

func CreateDaos() error {

	// 添加base.go|dboperation.go

	err := CreateFile("base.template", "./source/base.template", "./src/models/base.go")
	if err != nil {
		return err
	}
	err = CreateFile("dao.template", "./source/dao.template", "./src/dao/dboperation.go")
	if err != nil {
		return err
	}
	return nil

	// 生成daos

	//rd, err := ioutil.ReadDir(pathname)
	//for _, fi := range rd {
	//	if fi.IsDir() {
	//		// 目录跳过
	//		fmt.Printf("[%s]\n", pathname+"\\"+fi.Name())
	//	} else {
	//		fmt.Println(fi.Name())
	//	}
	//}
	//return err
}

func CreateFile (temp string, temppath string, filepath string) error {
	t := template.New(temp)
	t, err := t.ParseFiles(temppath)
	if err != nil {
		return err
	}
	v := models.Db{}
	file, err := os.Create(filepath)
	defer file.Close()
	if err != nil {
		return err
	}
	err = t.Execute(file, v)
	if err != nil {
		return err
	}
	return nil
}
