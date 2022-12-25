package service

import (
	"fmt"
	"redrock2022test/business/dao"
	"redrock2022test/business/model"
)

func SearchWarehose(name string) (u model.Warehose, err error) {
	fmt.Println("执行service.SearchWarehose")
	u, err = dao.SearchWarehose(name)
	return
}

func CreateWarehose(u model.Warehose) error {
	fmt.Println("执行service.CreateWarehose")
	err := dao.InsertWarehose(u)
	return err
}
