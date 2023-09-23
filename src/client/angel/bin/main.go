package main

import "log"
import "flag"
import "os"

import "github.com/jluong102/projectguardian/logger"

type cmdline struct {
	master string
	debug bool
}

func checkArguments(cmdArgs *cmdline) error {
	currentPath, err := os.Executable() // For setting default paths

	// Parse cmdline arguments
	flag.StringVar(&cmdArgs.master, "config", currentPath + "/config.json", "Master config file")
	flag.BoolVar(&cmdArgs.debug, "debug", false, "Debugging output")
	flag.Parse()

	return err
}

func main() {
	var cmdArgs *cmdline = new(cmdline)
	logTool := logger.CreateLogger("angel")
	err := checkArguments(cmdArgs)
	
	logTool.WriteInfo("Inializing...")
	log.Printf("Intializing...")

	if err != nil {
		log.Printf("Startup Error: %s", err)
		os.Exit(1)
	}

	log.Printf("Checking master config: %s", cmdArgs.master)
}
