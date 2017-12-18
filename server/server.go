package server

import (
	"github.com/oaStuff/clusteredBigCache/Cluster"
	"github.com/oaStuff/logservice"
)

//server configuration
type Config struct {
	Name 			string
	Join 			string
	LocalPort 		int
	ConfigFile 		string
	EnableLog 		bool
	LogFile 		string
	EnableHttp 		bool
	HttpPort 		int
}

//server description
type Server struct {
	Config 		*Config
	cache		*clusteredBigCache.ClusteredBigCache
	logger 		*asyncLogger.Logger
}


//create a new server
func New(config *Config, logger *asyncLogger.Logger) *Server {

	cacheConfig := clusteredBigCache.DefaultClusterConfig()
	cacheConfig.LocalPort = config.LocalPort
	cacheConfig.Id = config.Name
	if config.Join != "" {
		cacheConfig.Join = true
		cacheConfig.JoinIp = config.Join
	}


	return &Server{
		Config: config,
		logger: logger,
		cache: clusteredBigCache.New(cacheConfig, logger),
	}
}

//start the server
func (svr *Server) Start()  {
	svr.cache.Start()
}
