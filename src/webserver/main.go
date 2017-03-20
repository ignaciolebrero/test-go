package main

import "github.com/gin-gonic/gin"
import "github.com/mercadolibre/test-go/src/controller"

func main() {
  router := gin.Default()
  router.GET("/internal", func (c *gin.Context) {
    c.JSON(200, gin.H{ "message":"asd" })
  })

  router.GET("/prueba", controller.Prueba)
  router.Run(":8080")
}

// func main() {
//   router := gin.Default()
//   router.GET("/ping", func(c *gin.Context) {
//     c.JSON(200, gin.H{
//       "message": "pong",
//     })
//   })
//   router.Run() // listen and serve on 0.0.0.0:8080
// }

//set up mapping of URLs
// func MapUrl(router *gin.Engine) {

//   // Add health check
//   router.GET("/ping", func(c *gin.Context) {
//     c.String(200, "pong")
//   })

//   //documentación de swagger
//   router.GET("/internal/users/:user_id/:resources/messages/unread", controller.GetByUserId)
//   router.POST("/internal/users/:user_id/:resources/:resource_id/messages/unread/:message_id", controller.PostByUserId)
//   router.DELETE("/internal/users/:user_id/:resources/:resource_id/messages/unread/:message_id", controller.DeleteByUserId)
// }