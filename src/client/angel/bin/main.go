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
import "sync"

import "github.com/jluong102/projectguardian/logger"

type cmdline struct {
	master string
	debug bool
}

func checkArguments(cmdArgs *cmdline) error {
	// currentPath, err := os.Executable() // For setting default paths
	var err error = nil 

	// Parse cmdline arguments
	flag.StringVar(&cmdArgs.master, "config", "/etc/guardian/angel/master.json", "Master config file")
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
				ShowMasterSettings(logTool, masterConfig)
				var wg sync.WaitGroup
				
				// Validate watches
				for i, j := range masterConfig.Watches {
					err := SetupWatch(j)

					if err == nil {
						logTool.WriteDebug(fmt.Sprintf("Watch %d", i))
						ShowWatchSettings(logTool, j)
					} else {
						logTool.WriteError(fmt.Sprintf("Watch settings %d error: %s", i, err))
						os.Exit(WATCH_SETTINGS_ERROR)
					}
				}

				// Start watch threads 
				for _, i := range masterConfig.Watches {
					// Override debug settings 
					if cmdArgs.debug {
						i.Debug = cmdArgs.debug
					}

					wg.Add(1)
					logTool.WriteInfo(fmt.Sprintf("Starting watch thread: %s", i.Name))
					
					go func(watch *Watch) { 
						StartWatch(watch)
						wg.Done() 
					}(i)
				}

				wg.Wait()
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
