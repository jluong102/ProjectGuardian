package main 

import "os"
import "os/exec"
import "fmt"

import "github.com/jluong102/projectguardian/logger"
import "github.com/jluong102/projectguardian/permissions"

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

// Returns -9999 if script fails to run
// Otherwise should return the correct exit code
func runCheck(script string) int {
	cmd := exec.Command(script)
	err := cmd.Run()

	if err != nil {
		if exitCode, ok := err.(*exec.ExitError); ok {
			return exitCode.ExitCode()
		} else {
			// Not sure how I want to handle this yet, but for now
			fmt.Printf("UNEXPECTED ERROR: %s\n", err)

			return -9999
		}
	} 

	return 0
}

func StartWatch(watchSettings *Watch) {
	logTool := newLogger(watchSettings)
	logTool.WriteInfo(fmt.Sprintf("%s => Running check script: %s", watchSettings.Name, watchSettings.CheckScript))

	// Confirm check file exists
	if info, err := os.Stat(watchSettings.CheckScript); os.IsNotExist(err) {
		logTool.WriteError(fmt.Sprintf("Check script not found: %s", err))
		logTool.WriteInfo(fmt.Sprintf("Exiting watch %s", watchSettings.Name))

		return
	} else {
		if permissions.IsExecutableCurrentUser(info) {
			exitCode := runCheck(watchSettings.CheckScript)
			fmt.Printf("Status %d\n", exitCode)
		} else { // Make sure script is executable
			logTool.WriteError(fmt.Sprintf("Execution permission denied: %s", watchSettings.CheckScript))
			logTool.WriteInfo(fmt.Sprintf("Exiting watch %s", watchSettings.Name))
			
			return 
		}
	}
}