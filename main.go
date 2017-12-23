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
	"github.com/spf13/pflag"
	"fmt"
	"github.com/oaStuff/cacheServer/web"
)



func main() {

	svrConfig := parseProgramArgument()
	logger := asyncLogger.New(asyncLogger.LoggerConfig{Enabled:true, AllowFileLog:svrConfig.EnableLog,
										AllowConsoleLog:true, Filename:svrConfig.LogFile})

	logger.Info("cacheServer starting ...")
	svr := server.New(svrConfig, logger)
	svr.Start()
	logger.Info("cacheServer started")
	if svrConfig.EnableWeb {
		go web.StartHttpServer(svr)
	}


	//wait for the terminating signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
	logger.Info("cacheServer stopped")

}

func parseProgramArgument() *server.Config {

	flag.String("name", "", "server name")
	flag.String("join", "", "ipAddr:port number of remote server")
	flag.Int("port", -1, "local server port to bind to")
	configFile := flag.String("config", "", "server config file")
	flag.Bool("log", false, "enable server logging")
	flag.String("logfile", "", "server log file")
	flag.Bool("web", false, "enable the web endpoint")
	flag.Int("webport", -1, "web endpoint port")
	flag.Bool("webdebug", false, "debug web request")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	svr := &server.Config{}
	svr.ConfigFile = *configFile

	svr.ConfigFile = strings.Trim(svr.ConfigFile,"'")
	svr.ConfigFile = strings.Trim(svr.ConfigFile,"\"")

	if svr.ConfigFile != "" {
		//viper.AddConfigPath("./data/")
		//viper.SetConfigName("sampleConfig")
		viper.SetConfigFile(svr.ConfigFile)
		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}
	}

	svr.Name = viper.GetString("name")
	svr.LocalPort = viper.GetInt("port")
	svr.Join = viper.GetString("join")
	svr.ReconnectOnDisconnect = viper.GetBool("reconnect_on_disconnect")
	svr.EnableLog = viper.GetBool("log")
	svr.LogFile = viper.GetString("logfile")
	svr.EnableWeb = viper.GetBool("web")
	svr.WebPort = viper.GetInt("webport")
	svr.WebDebug = viper.GetBool("webdebug")


	svr.LogFile = strings.Trim(svr.LogFile,"'")
	svr.LogFile = strings.Trim(svr.LogFile,"\"")

	svr.Name = strings.Trim(svr.Name,"'")
	svr.Name = strings.Trim(svr.Name,"\"")

	if svr.LocalPort < 1024 {
		fmt.Fprintf(os.Stderr,"local port must be greater than 1024. current value is %d\n", svr.LocalPort)
		os.Exit(1)
	}

	if svr.EnableLog && svr.LogFile == "" {
		fmt.Fprintln(os.Stderr,"log file must be specified since logging is enabled")
		os.Exit(1)
	}

	if svr.EnableWeb {
		if svr.WebPort < 1024 {
			fmt.Fprintln(os.Stderr, "specify a web port higher than 1024")
			os.Exit(1)
		}

	}

	return svr
}
