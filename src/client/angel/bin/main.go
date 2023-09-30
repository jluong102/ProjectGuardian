//       _/_/                                  _/   
//    _/    _/  _/_/_/      _/_/_/    _/_/    _/    
//   _/_/_/_/  _/    _/  _/    _/  _/_/_/_/  _/     
//  _/    _/  _/    _/  _/    _/  _/        _/      
// _/    _/  _/    _/    _/_/_/    _/_/_/  _/       
//                          _/                      
//                     _/_/ 
package main

import "fmt"
import "flag"
import "os"
import "io/ioutil"
import "encoding/json"

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

func fileRead(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	defer file.Close()

	if err == nil {
		data, err := ioutil.ReadAll(file)

		if err == nil {
			return data, err
		}
	}

	return nil, err
}

func main() {
	var cmdArgs *cmdline = new(cmdline)
	logTool := logger.CreateLogger("angel")
	logTool.LogLevel = logger.LOG_NONE // Disable logging until config is read
	err := checkArguments(cmdArgs)
	
	fmt.Printf("Initalizing...\n")

	if err != nil {
		logTool.WriteError(fmt.Sprintf("Startup Error: %s", err))
		os.Exit(ARGUMENT_PARSE_ERROR)
	}

	logTool.WriteInfo(fmt.Sprintf("Attempting to load master config: %s", cmdArgs.master))
	rawMaster, err := fileRead(cmdArgs.master)

	if err == nil {
		logTool.WriteSuccess(fmt.Sprintf("Opened %s", cmdArgs.master))
		logTool.WriteInfo(fmt.Sprintf("Attempting to parse json from master config: %s", cmdArgs.master))

		var masterConfig *Master = new(Master)
		SetMasterDefaults(masterConfig)
		err := json.Unmarshal(rawMaster, masterConfig)

		if err == nil {
			logTool.WriteSuccess(fmt.Sprintf("Parsed json from master config: %s", cmdArgs.master))
			logTool.WriteInfo("Loading master settings")

			if cmdArgs.debug { 
				masterConfig.Debug = cmdArgs.debug
			}

			err := SetupMaster(masterConfig, logTool)

			if err == nil {
				logTool.WriteSuccess("Master settings loaded")
				logTool.WriteDebug(fmt.Sprintf("Log Path: %s", masterConfig.LogPath))
				logTool.WriteDebug(fmt.Sprintf("Debug: %t", masterConfig.Debug))
				logTool.WriteDebug(fmt.Sprintf("Disable Log Info: %t", masterConfig.NoInfo))
				logTool.WriteDebug(fmt.Sprintf("Disable Log Warning: %t", masterConfig.NoWarning))
				logTool.WriteDebug(fmt.Sprintf("Disable Log Error: %t", masterConfig.NoError))
				logTool.WriteDebug(fmt.Sprintf("Disable Log Success: %t", masterConfig.NoSuccess))

				for i, j := range masterConfig.Watches {
					logTool.WriteDebug(fmt.Sprintf("Watch %d -> Name: %s ", i, j.Name))
				}
			} else {
				logTool.WriteError(fmt.Sprintf("Master settings error: %s", err))
				os.Exit(MASTER_SETTINGS_ERROR)
			}

		} else {
			logTool.WriteError(fmt.Sprintf("Unable to parse json from master config: %s", err))
			os.Exit(JSON_PARSE_ERROR)
		}
	} else {
		logTool.WriteError(fmt.Sprintf("Unable to load master config: %s", err))
		os.Exit(MASTER_CONFIG_ERROR)
	}
}
