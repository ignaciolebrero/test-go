package controller

import "github.com/gin-gonic/gin"

func Prueba (c *gin.Context) {
	c.JSON(200, "listooooo")
}