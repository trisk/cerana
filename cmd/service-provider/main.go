package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/cerana/cerana/pkg/logrusx"
	"github.com/cerana/cerana/provider"
	"github.com/cerana/cerana/providers/service"
	flag "github.com/spf13/pflag"
)

func main() {
	logrus.SetFormatter(&logrusx.JSONFormatter{})

	config := service.NewConfig(nil, nil)
	flag.StringP("rollback_clone_cmd", "r", "/run/current-system/sw/bin/rollback_clone", "full path to dataset clone/rollback tool")
	flag.StringP("dataset_clone_dir", "d", "data/running-clones", "destination for dataset clones used by running services")
	flag.Parse()

	dieOnError(config.LoadConfig())
	dieOnError(config.SetupLogging())

	server, err := provider.NewServer(config.Config)
	dieOnError(err)
	s := service.New(config, server.Tracker())
	s.RegisterTasks(server)

	if len(server.RegisteredTasks()) != 0 {
		dieOnError(server.Start())
		server.StopOnSignal()
	} else {
		logrus.Warn("no registered tasks, exiting")
	}
}

func dieOnError(err error) {
	if err != nil {
		logrus.Fatal("encountered an error during startup, error:", err)
		os.Exit(1)
	}
}
