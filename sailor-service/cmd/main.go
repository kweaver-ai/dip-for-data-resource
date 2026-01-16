package main

import (
	"log"
	"os"
	"runtime"
)

var (
	Name = "af_glossary_service"
	// Version is the version of the compiled software.
	Version = "1.0"
)

// @title       af-sailor-service
// @version     0.0
// @description AnyFabric cognitive assistant
// @BasePath    /api
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ExecuteCmd()
}

func ExecuteCmd() {
	if err := rootCmd.Execute(); err != nil {
		log.Printf("command exec  failed, err: %+v", err)
		os.Exit(1)
	}
}
