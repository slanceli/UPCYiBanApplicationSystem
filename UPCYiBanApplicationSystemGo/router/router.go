package router

import (
	"UPCYiBanApplicationSystemGo/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Router *gin.Engine

func VerifyIdentidy() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		v := session.Get("name")
		if v == nil {
			c.String(http.StatusUnauthorized, "未登录")
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			c.Next()
		}
	}
}

func InitRouter () {
	Router = gin.Default()
	store := cookie.NewStore([]byte("Vhh0w3JBEiE6E4sqfjos"))
	Router.Use(sessions.Sessions("YiBanHouTai", store))
	Router.POST("/application", VerifyIdentidy(), func(c *gin.Context) {
		var body utils.Application
		session := sessions.Default(c)
		v := session.Get("name")
		if err := c.ShouldBindJSON(&body); err != nil {
			fmt.Println("Bind json failed, err: ", err)
			c.String(http.StatusInternalServerError, "failed")
		} else {
			fmt.Println(body)
			if body.AddApplication(v.(string)) {
				c.String(http.StatusOK, "successful")
			} else {
				c.String(http.StatusInternalServerError, "failed")
			}
		}
	})
	Router.POST("/login", func(c *gin.Context) {
		userName := c.PostForm("username")
		passwd := c.PostForm("password")
		if utils.UPClogin(userName, passwd) {
			session := sessions.Default(c)
			session.Set("name", userName)
			err := session.Save()
			if err != nil {
				fmt.Println("Save session failed, err:", err)
			}
			c.String(http.StatusOK, "successful")
		} else {
			c.String(http.StatusInternalServerError, "failed")
		}
	})
}

func RunRouter () {
	_ = Router.Run("127.0.0.1:6987")
}