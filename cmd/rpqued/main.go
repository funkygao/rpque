package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"github.com/funkygao/gafka/diagnostics/agent"
	"github.com/funkygao/golib/version"
	"github.com/funkygao/log4go"
	"github.com/funkygao/rpque/cmd/rpqued/server"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()

	t0 := time.Now()

	agent.HttpAddr = ":9017"
	log4go.Info("pprof agent ready on %s", agent.Start())
	go func() {
		log4go.Error("%s", <-agent.Errors)
	}()

	s := server.New()
	if err := s.Start(); err != nil {
		log4go.Error(err.Error())
		log4go.Info("rpqued[%s@%s] bye!", version.Revision, version.BuildDate)
		log4go.Close()
		os.Exit(1)
	}

	s.ServeForever()
	log4go.Info("rpqued[%s@%s] %s, bye!", version.Revision, version.BuildDate, time.Since(t0))
	log4go.Close()
}
