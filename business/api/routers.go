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

	wa := r.Group("/addwarehose", addWarehose)
	{
		wa.POST("/new")
		wa.PUT("/alterationgoods")
		wa.PUT("/inventory")
		wa.PUT("/deploy")
	}
}
