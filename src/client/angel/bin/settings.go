package main

import "fmt"

import "github.com/jluong102/projectguardian/logger"

// Use this to setup inital values as needed
// then overwrite with user provided settings from json
func SetMasterDefaults(masterData *Master) {
	masterData.LogPath = "/var/guardian/angel/logs"
	masterData.NoLog = false
	masterData.Debug = false
	masterData.NoInfo = false
	masterData.NoWarning = false
	masterData.NoError = false
	masterData.NoSuccess = false
}

func SetupMaster(masterData *Master, logTool *logger.LogTool) error {
	// Setup logging settings
	if !masterData.NoLog {
		if masterData.LogPath != "" {
			logTool.LogPath = masterData.LogPath
		}

		if masterData.Debug {
			logTool.LogLevel = logTool.LogLevel | logger.LOG_DEBUG
		}

		if !masterData.NoInfo {
			logTool.LogLevel = logTool.LogLevel | logger.LOG_INFO
		}

		if !masterData.NoWarning {
			logTool.LogLevel = logTool.LogLevel | logger.LOG_WARNING
		}

		if !masterData.NoError {
			logTool.LogLevel = logTool.LogLevel | logger.LOG_ERROR
		}
		
		if !masterData.NoSuccess {
			logTool.LogLevel = logTool.LogLevel | logger.LOG_SUCCESS
		}

		logTool.LogPath = masterData.LogPath
		logTool.Print = !masterData.NoConsole
	}

	if len(masterData.Watches) < 1 {
		return fmt.Errorf("No watches found")
	}

	return nil
}

func ShowMasterSettings(logTool *logger.LogTool, masterSettings *Master) {
	output := fmt.Sprintf("Log Path: %s", masterSettings.LogPath)
	output += fmt.Sprintf("\n\t -> Debug: %t", masterSettings.Debug)
	output += fmt.Sprintf("\n\t -> Disable Log Info: %t", masterSettings.NoInfo)
	output += fmt.Sprintf("\n\t -> Disable Log Warning: %t", masterSettings.NoWarning)
	output += fmt.Sprintf("\n\t -> Disable Log Error: %t", masterSettings.NoError)
	output += fmt.Sprintf("\n\t -> Disable Log Success: %t", masterSettings.NoSuccess)
	output += fmt.Sprintf("\n\t -> Disable Log Console: %t", masterSettings.NoConsole)

	logTool.WriteDebug(output)
}

/*func SetWatchDefaults(watchData *Watch) {
	watchData.Name = ""
	watchData.Interval = -1
	watchData.LogPath = "/var/guardian/angel/logs"
	watchData.Debug = false
	watchData.NoLog = false
	watchData.NoInfo = false
	watchData.NoLog = false
	watchData.NoSuccess = false
	watchData.NoWarning = false
	watchData.NoError = false
	watchData.NoConsole = false	
}*/

func SetupWatch(watchData *Watch) error {
	if watchData.Interval < 1 {
		return fmt.Errorf("Interval must be set to 1 minute or longer")
	}

	return nil
}

func ShowWatchSettings(logTool *logger.LogTool, watchSettings *Watch) {
	output := fmt.Sprintf("Name: %s", watchSettings.Name)
	output += fmt.Sprintf("\n\t -> Interval: %d minutes", watchSettings.Interval)
	output += fmt.Sprintf("\n\t -> Debug: %t", watchSettings.Debug)
	output += fmt.Sprintf("\n\t -> Disable Log Info: %t", watchSettings.NoInfo)
	output += fmt.Sprintf("\n\t -> Disable Log Warning: %t", watchSettings.NoWarning)
	output += fmt.Sprintf("\n\t -> Disable Log Error: %t", watchSettings.NoError)
	output += fmt.Sprintf("\n\t -> Disable Log Success: %t", watchSettings.NoSuccess)
	output += fmt.Sprintf("\n\t -> Disable Log Console: %t", watchSettings.NoConsole)

	logTool.WriteDebug(output)
}