package main

import (
	"flag"
	"runtime"
	"time"

	"math/rand"

	"wangqingang/domain/common"
	"wangqingang/domain/handler"
)

const (
	cmdKeyConfig  = "config"
	cmdHelpConfig = "config file's path"
)

func cmdConfigHandler(config string) {
	if err := common.InitConfig(config); err != nil {
		panic(err)
	}
	if common.Config.Gomaxprocs >= 1 {
		runtime.GOMAXPROCS(common.Config.Gomaxprocs)
	}
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	config := flag.String(cmdKeyConfig, "", cmdHelpConfig)
	flag.Parse()

	cmdConfigHandler(*config)
	router := handler.ServerEngine()
	router.Run(common.Config.Listen)
}
