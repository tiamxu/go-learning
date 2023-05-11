package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func api_test() {
	//创建路由
	r := gin.Default()
	//绑定路由规则，执行函数
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, "hello world,"+name+" "+action)
	})

	r.Run(":8080")
}
func url_test() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		name := c.DefaultQuery("name", "徐亮")
		c.String(http.StatusOK, "hello "+name)
	})
	r.Run(":8080")
}

func form_test() {
	r := gin.Default()

	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.String(http.StatusOK,
			fmt.Sprintf("username:%s password:%s", username, password))
	})
	r.Run(":8080")
}

type Login struct {
	Username string `form:"username" json:"username" uri:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func json_test() {
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		var loginJson Login
		if err := c.ShouldBindJSON(&loginJson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if loginJson.Username != "root" || loginJson.Password != "123456" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	r.Run(":8081")
}

func main() {
	//api_test()
	// url_test()
	// form_test()
	//json_test()
	// env := "stage_hello"
	// namespace := strings.Split(env, "_")[0]
	// fmt.Printf("env:%s,namespace:%s\n", env, namespace)

}
