package service

import (
	"fmt"
	"redrock2022test/business/dao"
	"redrock2022test/business/model"
)

func SearchUserByUserName(name string) (u model.User, err error) {
	fmt.Println("执行service.SearchUserByUserName")
	u, err = dao.SearchUserByUserName(name)
	return
}
