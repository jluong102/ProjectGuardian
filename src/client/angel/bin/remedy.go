package main

import "github.com/jluong102/projectguardian/logger"

func findRemedy(watchSettings *Watch, exitCode int) string {
	for i, j := range watchSettings.Remedies {
		if j.OnCode == exitCode {
			return i
		}
	}

	return "" // Empty string if nothing found
}