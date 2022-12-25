package dao

import (
	"fmt"
	"redrock2022test/business/model"
)

func SearchWarehose(warehosesName string) (u model.Warehose, err error) {
	row := DB.QueryRow("select id,name,goods from warehoses where warehosesName = ?", warehosesName)
	fmt.Println("执行dao.SearchWarehose")
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.ID, &u.WarehoseName, &u.Goods)
	return
}

func InsertWarehose(u model.Warehose) (err error) {
	fmt.Println("执行InsertWarehose")
	_, err = DB.Exec("insert into warehoses (name,password) values (?)", u.WarehoseName)
	return
}
