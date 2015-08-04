package main

import (
	"flag"
	"strings"
)

var config struct {
	Pid       string
	Timeout   int
	Port      int
	LogPath   string
	Producers producers
}

type producers []string

func (i *producers) Set(value string) error {
	for _, proc := range strings.Split(value, ",") {
		*i = append(*i, proc)
	}
	return nil
}

func (i *producers) String() string {
	return strings.Join(*i, ",")
}

func init() {
	config.Producers = make([]string, 0, 0)
	flag.IntVar(&config.Timeout, "timeout", 30, "timeout, in seconds")
	flag.IntVar(&config.Port, "port", 8080, "HTTP service port")
	flag.StringVar(&config.Pid, "pid", "demo-pid-01", "participant ID")
	flag.StringVar(&config.LogPath, "log", "liveqa.log", "path to log file")
	flag.Var(&config.Producers, "producer", "comma separated list of processors to use on this server")
}
