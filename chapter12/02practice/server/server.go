package main

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"

	"go.learn/chapter12/02practice/frinterface"
	"go.learn/pkg/apis"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func connectDb() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("root:123456@tcp(localhost)/testdb"))
	if err != nil {
		log.Fatal("数据库链接失败", err)
	}
	fmt.Println("链接数据库成功")
	return conn
}

func main() {
	conn := connectDb()
	var rankServer frinterface.ServerInterface = NewDbRank(conn, NewFatRateRank())

	if initRank, ok := rankServer.(frinterface.RankInitInterface); ok {
		if err := initRank.Init(); err != nil {
			log.Fatal("初始化失败", err)
		}
	}
	r := gin.Default()
	pprof.Register(r)
	r.POST("/register", func(c *gin.Context) {
		var pi *apis.PersonalInformation
		if err := c.BindJSON(&pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "无法解析注册信息" + err.Error(),
			})
			return
		}
		if err := rankServer.RegisterPersonalInformation(pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "注册失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Message": "注册成功",
		})
	})
	r.PUT("/personalinfo", func(c *gin.Context) {
		var pi *apis.PersonalInformation
		if err := c.BindJSON(&pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "无法解析注册信息",
			})
			return
		}
		if resp, err := rankServer.UpdatePersonalInformation(pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "注册失败",
			})
			return
		} else {
			c.JSON(http.StatusOK, resp)
		}
	})
	r.GET("/rank/:name", func(c *gin.Context) {
		name := c.Param("name")
		if fr, err := rankServer.GetFatRate(name); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "获取个人排行失败",
			})
			return
		} else {
			c.JSON(http.StatusOK, fr)
		}
	})
	r.GET("/ranktop", func(c *gin.Context) {
		if frTop, err := rankServer.GetTop(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "获取排行失败",
			})
			return
		} else {
			c.JSON(http.StatusOK, frTop)
		}
	})
	if err := r.Run(":8081"); err != nil {
		log.Fatal("error")
	}

}
