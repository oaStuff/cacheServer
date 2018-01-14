package web

import (
	"net"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-contrib/cors"
	"github.com/oaStuff/cacheServer/server"
	"net/http"
	"time"
	"net/http/pprof"
	"fmt"
)

func StartHttpServer(svr *server.Server) {


	if !svr.Config.WebDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	g := gin.Default()
	g.Use(cors.Default())
	g.Use(static.Serve("/", static.LocalFile("./web", true)))
	g.NoRoute(func(c *gin.Context) {
		c.File("./web/index.html")
	})

	g.POST("/data/:key", func(c *gin.Context) {
		key := c.Param("key")
		d := c.GetHeader("duration")
		fmt.Println("duration is ", d)
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

	g.Any("/debug/pprof/", gin.WrapF(pprof.Index))
	g.Any("/debug/pprof/cmdline", gin.WrapF(pprof.Cmdline))
	g.Any("/debug/pprof/profile", gin.WrapF(pprof.Profile))
	g.Any("/debug/pprof/symbol", gin.WrapF(pprof.Symbol))
	g.Any("/debug/pprof/trace", gin.WrapF(pprof.Trace))

	g.Any("/debug/pprof/heap", gin.WrapH(pprof.Handler("heap")))
	g.Any("/debug/pprof/block", gin.WrapH(pprof.Handler("block")))
	g.Any("/debug/pprof/goroutine", gin.WrapH(pprof.Handler("goroutine")))
	g.Any("/debug/pprof/threadcreate", gin.WrapH(pprof.Handler("threadcreate")))





	g.Run(net.JoinHostPort("",strconv.Itoa(svr.Config.WebPort)))
}
