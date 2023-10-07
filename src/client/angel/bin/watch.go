package main 

import "os/exec"
import "fmt"

import "github.com/jluong102/projectguardian/logger"

func newLogger(watchSettings *Watch) *logger.LogTool {
	// Create new logger for thread
	logTool := logger.CreateLogger(watchSettings.Name)
	logTool.LogLevel = logger.LOG_NONE

	if !watchSettings.NoLog {
		if watchSettings.LogPath != "" {
			logTool.LogPath = watchSettings.LogPath
		} else {
			logTool.LogPath = "/var/guardian/angel/logs"
		}

		if !watchSettings.Debug {
			logTool.LogLevel = logTool.LogLevel | logger.LOG_DEBUG
		}

		if !watchSettings.NoInfo {
			logTool.LogLevel = logTool.LogLevel | logger.LOG_INFO
		}

		if !watchSettings.NoWarning {
			logTool.LogLevel = logTool.LogLevel | logger.LOG_WARNING
		}

		if !watchSettings.NoError {
			logTool.LogLevel = logTool.LogLevel | logger.LOG_ERROR
		}
		
		if !watchSettings.NoSuccess {
			logTool.LogLevel = logTool.LogLevel | logger.LOG_SUCCESS
		}

		logTool.Print = !watchSettings.NoConsole
	}

	return logTool
}

func RunCheck(script string) int {
	cmd := exec.Command(script)
	err := cmd.Run()

	if err != nil {
		if exitCode, ok := err.(*exec.ExitError); ok {
			return exitCode.ExitCode()
		}
	}

	return 0
}

func StartWatch(watchSettings *Watch) {
	logTool := newLogger(watchSettings)

	logTool.WriteInfo(fmt.Sprintf("%s => Running check script: %s", watchSettings.Name, watchSettings.CheckScript))
}