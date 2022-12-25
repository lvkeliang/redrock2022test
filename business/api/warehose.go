package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"redrock2022test/business/model"
	"redrock2022test/business/service"
	"redrock2022test/business/util"
)

func addWarehose(c *gin.Context) {
	warehoseName := c.PostForm("name")
	if warehoseName == "" {
		fmt.Println("输入空字符")
		util.RespParamErr(c)
		return
	}

	u, err := service.SearchWarehose(warehoseName)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("service.SearchWarehose出问题了")
		log.Printf("search Warehose error : %v", err)
		util.RespInternalErr(c)
		return
	}
	if u.WarehoseName != "" {
		util.NormErr(c, 300, "仓库已存在")
		return
	}

	err = service.CreateWarehose(model.Warehose{
		WarehoseName: warehoseName,
	})

	if err != nil {
		fmt.Println("创建仓库出问题了")
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}
