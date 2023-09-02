package main

import (
	"github.com/aparnasukesh/RESTapi-GO/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	controller.Routes(r)

	r.Run(":2000")
}
