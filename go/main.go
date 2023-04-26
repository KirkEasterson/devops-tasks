package main

import (
	"flag"

	"gocounter/api/server"
)

type CmdArgs struct {
	Url  string
	Port string
}

func main() {
	cmdArgs := getArgs()
	serverConfig := server.ServerConfig{
		Url:  cmdArgs.Url,
		Port: cmdArgs.Port,
	}
	server.Serve(serverConfig)
}

func getArgs() CmdArgs {
	url := flag.String("u", "127.0.0.1", "the URL for the server to use")
	port := flag.String("p", "8080", "the port for the server to use")
	flag.Parse()

	cmdArgs := CmdArgs{
		Url:  *url,
		Port: *port,
	}
	return cmdArgs
}
