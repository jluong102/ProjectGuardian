package main

// import "fmt"

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
	}

	return nil
}