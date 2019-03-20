package utils

import (
	"../models"
	"fmt"
	"os"
)

type DirStruct struct {
	root string
	dirs [] string
}

func NewDirStruct(dic *models.Dic) *DirStruct {
	var list []string
	for _,v := range dic.Dics{
		list = append(list, v.Name)
	}
	return &DirStruct{
		root: dic.Root.Name,
		dirs: list,
	}
}

func (ds *DirStruct)  CreateDirs ()  {
	err := os.MkdirAll("./" + ds.root, os.ModePerm)
	if err!=nil{
		fmt.Println(err)
		return
	}
	for _, v := range ds.dirs{
		err:=os.Mkdir("./" + ds.root+ "/" + v, os.ModePerm)
		if err!=nil{
			fmt.Println(err)
			return
		}
	}
}
