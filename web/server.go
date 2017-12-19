package web

import (
	"net"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
	"github.com/oaStuff/cacheServer/server"
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

	g.Run(net.JoinHostPort("",strconv.Itoa(svr.Config.WebPort)))
}
