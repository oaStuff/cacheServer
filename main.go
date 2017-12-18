package main

import (
	"flag"
	"strings"
	"os"
	"os/signal"
	"syscall"
	"github.com/oaStuff/logservice"
	"github.com/oaStuff/cacheServer/server"
	"github.com/spf13/viper"
	"fmt"
)



func main() {

	svrConfig := parseProgramArgument()
	logger := asyncLogger.New(asyncLogger.LoggerConfig{Enabled:true, AllowFileLog:svrConfig.EnableLog,
										AllowConsoleLog:true, Filename:svrConfig.LogFile})

	logger.Info("cacheServer starting ...")
	svr := server.New(svrConfig, logger)
	svr.Start()
	logger.Info("cacheServer started")


	//wait for the terminating signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
	logger.Info("cacheServer stopped")

}

func parseProgramArgument() *server.Config {
	name := flag.String("name", "", "server name")
	join := flag.String("join", "", "ipAddr:port number of remote server")
	localPort := flag.Int("port", -1, "local server port to bind to")
	configFile := flag.String("config", "", "server config file")
	enableLog := flag.Bool("log", false, "enable server logging")
	logFile := flag.String("logfile", "", "server log file")
	enableHttp := flag.Bool("http", false, "enable the http endpoint")
	httpPort := flag.Int("httpport", -1, "http endpoint port")

	flag.Parse()

	svr := &server.Config{}
	svr.LocalPort = *localPort
	svr.Join = *join
	svr.ConfigFile = *configFile
	svr.EnableHttp = *enableHttp
	svr.EnableLog = *enableLog
	svr.LogFile = *logFile
	svr.HttpPort = *httpPort
	svr.Name = *name

	svr.LogFile = strings.Trim(svr.LogFile,"'")
	svr.LogFile = strings.Trim(svr.LogFile,"\"")

	svr.Name = strings.Trim(svr.Name,"'")
	svr.Name = strings.Trim(svr.Name,"\"")

	if svr.EnableLog && svr.LogFile == "" {
		panic("log file must be specified since logging is enabled")
	}

	if svr.EnableHttp {
		if svr.HttpPort < 1 {
			panic("http port must be specified since http is enabled")
		}

		if svr.HttpPort < 1024 {
			panic("specify a port higher than 1024")
		}

	}

	svr.ConfigFile = strings.Trim(svr.ConfigFile,"'")
	svr.ConfigFile = strings.Trim(svr.ConfigFile,"\"")

	if svr.ConfigFile != "" {
		//viper.AddConfigPath("./data/")
		//viper.SetConfigName("sampleConfig")
		viper.SetConfigFile(svr.ConfigFile)
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(viper.Get("name"))
	}

	return svr
}
