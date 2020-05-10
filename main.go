package main

import (
	"github.com/gin-gonic/gin"
	"test.blog.com/testBlog/common"
)

func main() {
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = CollectRouter(r)
	r.Run("0.0.0.0:9090") // listen and serve on 0.0.0.0:8080
}
