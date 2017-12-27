package server

import (
	"github.com/oaStuff/clusteredBigCache/Cluster"
	"github.com/oaStuff/logservice"
	"time"
)

//server configuration
type Config struct {
	Name                  string
	Join                  string
	ReconnectOnDisconnect bool
	LocalPort             int
	ConfigFile            string
	EnableLog             bool
	LogFile               string
	EnableWeb             bool
	WebPort               int
	WebDebug              bool
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
	cacheConfig.ReconnectOnDisconnect = config.ReconnectOnDisconnect
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

func (svr *Server) Put(key string, data []byte, duration time.Duration) {
	svr.cache.Put(key, data, duration)
}

func (svr *Server) Get(key string, duration time.Duration) ([]byte, error) {
	return svr.cache.Get(key, duration)
}

func (svr *Server) Delete(key string) error {
	return svr.cache.Delete(key)
}