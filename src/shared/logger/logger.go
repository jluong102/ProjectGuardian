package	logger

import "fmt"
import "time"
import "os"

// Logger is used as a way to have consistent output 
// to both stdout and any logfiles.
// Logger will also handle log rotations

type LogTool struct {
	LogPath string
	LogLevel int16
	Filename string
	AutoRotate bool
	Print bool
}

const (
	LOG_NONE = 0
	LOG_DEBUG = 1
	LOG_INFO = 1 << 1
	LOG_WARNING = 2 << 1
	LOG_ERROR = 3 << 1
	LOG_SUCCESS = 4 << 1
)

// PUBLIC
func (this LogTool) WriteInfo(msg string) {
	now := time.Now().Format(time.UnixDate)
	msg = fmt.Sprintf("[%s][INFO]%s\n", now, msg)

	if this.Print {
		fmt.Printf(msg)
	}

	if (this.LogLevel & LOG_INFO) != 0 {
		this.AppendFile(msg)
	}
}

func (this LogTool) WriteWarning(msg string) {
	now := time.Now().Format(time.UnixDate)
	msg = fmt.Sprintf("[%s][WARNING]%s\n", now, msg)

	if this.Print {
		fmt.Printf(msg)
	}

	if (this.LogLevel & LOG_WARNING) != 0 {
		this.AppendFile(msg)
	}
}

func (this LogTool) WriteError(msg string) {
	now :=  time.Now().Format(time.UnixDate)
	msg = fmt.Sprintf("[%s][ERROR]%s\n", now, msg)

	if this.Print {
		fmt.Printf(msg)
	}

	if (this.LogLevel & LOG_ERROR) != 0 {
		this.AppendFile(msg)
	}
}

func (this LogTool) WriteDebug(msg string) {
	now := time.Now().Format(time.UnixDate)
	msg = fmt.Sprintf("[%s][DEBUG]%s\n", now, msg)

	if this.Print {
		fmt.Printf(msg)
	}

	if (this.LogLevel & LOG_DEBUG) != 0 {
		this.AppendFile(msg)
	}
}

func (this LogTool) WriteSuccess(msg string) {
	now := time.Now().Format(time.UnixDate)
	msg = fmt.Sprintf("[%s][SUCCESS]%s\n", now, msg)

	if this.Print {
		fmt.Printf(msg)
	}

	if (this.LogLevel & LOG_SUCCESS) != 0 {
		this.AppendFile(msg)
	}
}

func (this LogTool) AppendFile(msg string) {
	filename := this.LogPath + this.Filename

	if this.AutoRotate {
		filename += "_" + time.Now().Format("2006-02-01") 
	}

	filename += ".log" // End all log files with ".log"
	stream, err := os.OpenFile(filename, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0744)

	if err != nil {
		now := time.Now().Format(time.UnixDate)
		fmt.Printf("[%s][FAIL]Error: %s\n", now, err)
	} else {
		_, err := stream.Write([]byte(msg))

		if err != nil {
			now := time.Now().Format(time.UnixDate)
			fmt.Printf("[%s][FAIL]Error: %s\n", now, err)
		}
	}

	stream.Close()
}

// Constructor
func CreateLogger(filename string) *LogTool {
	var logger *LogTool = new(LogTool)
	logger.Filename = filename

	// Set defaults 
	logger.LogPath = "/var/guardian/misc/logs"
	logger.LogLevel = LOG_INFO | LOG_WARNING | LOG_ERROR
	logger.AutoRotate = true
	logger.Print = true

	return logger
}