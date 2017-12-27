package web

import (
	"net"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
	"github.com/oaStuff/cacheServer/server"
	"net/http"
	"time"
)

func StartHttpServer(svr *server.Server) {


	if !svr.Config.WebDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	g := gin.Default()
	g.Use(static.Serve("/", static.LocalFile("./web", true)))
	g.NoRoute(func(c *gin.Context) {
		c.File("./web/index.html")
	})

	g.POST("/data/:key", func(c *gin.Context) {
		key := c.Param("key")
		d := c.GetHeader("duration")
		var duration time.Duration
		if d == "" {
			duration = 0
		} else {
			v, _ := strconv.Atoi(d)
			duration = time.Second * time.Duration(v)
		}
		data, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusExpectationFailed, map[string]string{"code":"c_500", "msg":err.Error()})
		}
		svr.Put(key, data, duration)
		c.JSON(http.StatusOK, map[string]string{"code":"c_200", "msg":"successful"})
	})

	g.GET("/data/:key", func(c *gin.Context) {
		key := c.Param("key")
		data, err := svr.Get(key, time.Millisecond * 300)
		if err != nil {
			c.JSON(http.StatusNotFound, map[string]string{"code":"c_404", "msg":err.Error()})
		}

		c.Status(http.StatusOK)
		c.Writer.Header().Set("Content-Type","application/octet-stream")
		c.Writer.Write(data)

	})

	g.DELETE("/data/:key", func(c *gin.Context) {
		key := c.Param("key")
		if err := svr.Delete(key); err != nil {
			c.JSON(http.StatusExpectationFailed, map[string]string{"code":"c_500", "msg":err.Error()})
		}

		c.JSON(http.StatusOK, map[string]string{"code":"c_200", "msg":"successful"})
	})

	g.Run(net.JoinHostPort("",strconv.Itoa(svr.Config.WebPort)))
}
