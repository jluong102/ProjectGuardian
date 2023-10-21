package main 

import "os/exec"
import "fmt"
import "time"

import "github.com/jluong102/projectguardian/logger"
// import "github.com/jluong102/projectguardian/permissions"

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

	// Keep running until we tell it to stop
	for {
		// Confirm check file exists and is executable
		if err := CheckExecutableScript(watchSettings.CheckScript); err == nil {
			logTool.WriteInfo(fmt.Sprintf("%s => Running check script: %s", watchSettings.Name, watchSettings.CheckScript))
			exitCode := runCheck(watchSettings.CheckScript)
			codeFound := false

			fmt.Printf("Status %d\n", exitCode)

			for _, i := range watchSettings.SuccessCodes {
				if int(i) == exitCode {
					codeFound = true
					fmt.Printf("good\n")
					time.Sleep(time.Duration(watchSettings.Interval) * time.Minute)
					break
				}
			}

			// Only check failure codes if no success codes are found
			if !codeFound {
				for _, i := range watchSettings.FailureCodes {
					if int(i) == exitCode {
						fmt.Printf("Bad\n")
						time.Sleep(time.Duration(watchSettings.Interval) * time.Minute)
						break
					}
				}
			}
		} else {
			logTool.WriteError(fmt.Sprintf("Unable to execute script: %s", err))

			// temporary 
			return
		}
	}
}