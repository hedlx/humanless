package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		emu := v1.Group("/emu")
		{
			emu.GET("/", func(c *gin.Context) {
				c.JSON(200, "[List] /v1/emu")
			})
			emu.GET("/:id", func(c *gin.Context) {
				c.JSON(200, "[Get] /v1/emu/"+c.Param("id"))
			})
			emu.POST("/:id/register", func(c *gin.Context) {
				c.JSON(200, "[Post] /v1/emu/"+c.Param("id")+"/register")
			})
		}

		sess := v1.Group("/sessions")
		{
			sess.GET("/", func(c *gin.Context) {
				c.JSON(200, "[List] /v1/sessions")
			})
			sess.GET("/:id", func(c *gin.Context) {
				c.JSON(200, "[Get] /v1/sessions/"+c.Param("id"))
			})
			sess.GET("/:id/replay", func(c *gin.Context) {
				c.JSON(200, "[Get] /v1/sessions/"+c.Param("id")+"/replay")
			})
		}
	}

	r.Run()
}
