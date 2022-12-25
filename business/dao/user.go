package dao

import (
	"fmt"
	"redrock2022test/business/model"
)

func SearchUserByUserName(name string) (u model.User, err error) {
	row := DB.QueryRow("select id,name,password from user where name = ?", name)
	fmt.Println("执行dao.SearchUserByUserName")
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.ID, &u.UserName, &u.Password)
	return
}
