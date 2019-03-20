package main

import (
	"../utils"
	"../models"
	"fmt"
)

func main(){
	// 读取json文件
	JsonParse := utils.NewJsonStruct()
	v := models.Dic{}
	JsonParse.Load("./configs/dic.json", &v)
	fmt.Println("读取配置")
	// 创建目录
	mk := utils.NewDirStruct(&v)
	mk.CreateDirs()
	// 生成代码
	err := utils.CreateDBhelper()
	utils.CheckError(err)
	err = utils.CreateDaos()
	utils.CheckError(err)
}


