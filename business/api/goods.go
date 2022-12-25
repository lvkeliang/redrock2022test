package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"redrock2022test/business/service"
	"redrock2022test/business/util"
)

func Addgoods(c *gin.Context) {

}

func OffShelveGoods(c *gin.Context) {

}

func InventoryIn(c *gin.Context) {
	//检验登录
	cookie, err := c.Cookie("name")
	_, err = service.SearchUserByUserName(cookie)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "未登录")
		} else {
			log.Printf("search user error : %v", err)
			util.RespInternalErr(c)
			return
		}
		return
	}

	goodsName := c.PostForm("goodsName")
	err := service.InventoryIn(c, goodsName)
	if err != nil {
		util.RespParamErr(c)
		return
	}
	util.RespOK(c)
}

func InventoryOff(c *gin.Context) {

}

func Deploy(c *gin.Context) {

}
