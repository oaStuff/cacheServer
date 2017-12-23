package web

import (
	"net"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
	"github.com/oaStuff/cacheServer/server"
	"net/http"
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

	g.POST("/data", func(c *gin.Context) {
		c.String(http.StatusNoContent,"not yet implemented")
	})

	g.GET("/data/:key", func(c *gin.Context) {
		c.String(http.StatusNoContent,"not yet implemented")
	})

	g.DELETE("/data/:key", func(c *gin.Context) {
		c.String(http.StatusNoContent,"not yet implemented")
	})

	g.Run(net.JoinHostPort("",strconv.Itoa(svr.Config.WebPort)))
}
