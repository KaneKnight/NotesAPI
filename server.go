package main

import "github.com/gin-gonic/gin"
import "net/http"


func main() {
	r := gin.Default()
	r.GET("/", home)
	r.Run(":8080")
}

func home(c* gin.Context) {
	c.String(http.StatusOK, "Hello.")
}
