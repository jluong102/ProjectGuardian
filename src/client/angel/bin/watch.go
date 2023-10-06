package main 

import "github.com/jluong102/projectguardian/logger"

func newLogger(watchSettings *Watch) *logger.LogTool {
	// Create new logger for thread
	logTool := logger.CreateLogger(watchSettings.Name)

	if !watchSettings.NoLog {
		logTool.LogPath = watchSettings.LogPath
	}

	return logTool
}

func StartWatch(watchSettings *Watch) {
	newLogger(watchSettings)
}