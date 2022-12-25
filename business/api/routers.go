package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()

	u := r.Group("/user")
	{
		u.GET("/register", func(c *gin.Context) {
			c.String(200, "ok\n")
			return
		})
		u.POST("/login", Login)
	}

	wa := r.Group("/addwarehose", AddWarehose)
	{
		wa.PUT("/alterationgoods/add", Addgoods)
		wa.PUT("/alterationgoods/off", OffShelveGoods)
		wa.PUT("/inventory/in", InventoryIn)
		wa.PUT("/inventory/off", InventoryOff)
		wa.PUT("/deploy", Deploy)
	}
}
