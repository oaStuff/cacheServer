package main

import "flag"

func main() {

	parseProgramArgument()
}

func parseProgramArgument() {
	join := flag.String("join", "", "ipAddr:port number of remote server")
	localPort := flag.Int("port", 9911, "local server port to bind to")
	configFile := flag.String("config", "", "server config file")
	enableLog := flag.Bool("log", false, "enable server logging")
	logFile := flag.String("logfile", "", "server log file")
	enableHttp := flag.Bool("http", false, "enable the http endpoint")
	httpPort := flag.Int("httpPort", 0, "http endpoint port")

}
