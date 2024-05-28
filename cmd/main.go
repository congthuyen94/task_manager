package main

import (
	"flag"
	"os"
	"task_manager/internal/config"

	"github.com/romana/rlog"
)

type options struct {
	SocketIOServer string `yaml:"socketIOServer" json:"socketIOServer" env:"SOCKETIO_SERVER" env-default:"localhost:8000"`
}

func getConfigFileFromArgs() string {
	var configFile string
	flag.StringVar(&configFile, "c", "", "Config File Location")
	flag.Parse()

	return configFile
}

func main() {
	rlog.Info("Start of program")
	var opts options
	configFile := getConfigFileFromArgs()
	if err := config.LoadConfig(configFile, &opts); err != nil {
		os.Exit(1)
	}
	rlog.Info(opts.SocketIOServer)
}
