package main

import (
	"context"
	"gins/app/helper"
	"gins/app/model"
	_ "gins/config/db"
	"gins/config/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var ctx = context.Background()

//GIN_MODE=release
func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.String(200, "")
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"result": "hello"})
	})
	r.GET("/redis", func(c *gin.Context) {
		rdb := redis.GetRedis()
		key := "setKey:" + strconv.Itoa(helper.Random(99999999))
		err := rdb.Set(ctx, key, helper.Random(99999999), 0).Err()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"result": err.Error()})
		}

		_, err2 := rdb.Get(ctx, key).Result()
		if err2 != nil {
			c.JSON(http.StatusOK, gin.H{"result": err2.Error()})
		}

		res, err3 := rdb.Del(ctx, key).Result()
		if err3 != nil {
			c.JSON(http.StatusOK, gin.H{"result": err3.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"result": res})
	})
	r.GET("/mysql", func(c *gin.Context) {
		// 写入数据
		err, id := (&model.Tests{}).CreateTests("name-" + strconv.Itoa(helper.Random(9999)))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"result": err.Error()})
		}

		_, err2 := (&model.Tests{}).GetIdByTests(id)
		if err2 != nil {
			c.JSON(http.StatusOK, gin.H{"result": err.Error()})
		}

		err3 := (&model.Tests{}).DeleteTests(model.Tests{ID: id})
		if err3 != nil {
			c.JSON(http.StatusOK, gin.H{"result": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"result": id})
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
