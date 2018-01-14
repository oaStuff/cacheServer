cacheServer
===========

This is a simple application demonstrating the use of [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) 

It takes configuration from command line or a config file. Every config from the command line overrides the same config from the configuration file.


## Installing

### Using *go get* and *dep ensure* to pull dependencies

    $ go get github.com/oaStuff/cacheServer
    $ dep ensure

## Running it

    $ $GOPATH/bin/cacheServer [--config (path to config file)]
    
    
    
## Config file

The configuration file is a simple json file with the following content

```json
{
    "name" : "server_id_1",
    "join" : "",
    "port" : 9910,
    "reconnect_on_disconnect" : false,
    "log" : false,
    "logfile" : "./logs/app.log",
    "web" : true,
    "webport" : 8080,
    "webdebug" : false
}
```

The configuration file is pretty simple. The *join* field is used for clustering. It specified the remote server's IP
and port number using the format `"localhost:9900"`. So you can run multiple server with a base config file and pass other
parameters via the command line.

```text
    ./cacheServer --port 10001
    ./cacheServer --port 10002 --join localhost:10001
    ./cacheServer --port 10003 --join localhost:10001
    ./cacheServer --port 10004 --join localhost:10001
```