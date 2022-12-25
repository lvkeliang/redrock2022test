package main

import (
	"redrock2022test/business/api"
	"redrock2022test/business/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
