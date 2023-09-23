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
	LOG_DEBUG = 1
	LOG_INFO = 1 << 1
	LOG_WARNING = 2 << 1
	LOG_ERROR = 3 << 1
)

// PUBLIC
func (this LogTool) WriteInfo(msg string) {
	now := time.Now().Format(time.ANSIC)
	msg = fmt.Sprintf("[%s][INFO]%s\n", now, msg)

	if this.Print {
		fmt.Printf(msg)
	}

	if (this.LogLevel & LOG_INFO) != 0 {
		AppendFile(this.LogPath + this.Filename, msg)
	}
}

func (this LogTool) WriteWarning(msg string) {
	now := time.Now().Format(time.ANSIC)
	msg = fmt.Sprintf("[%s][WARNING]%s\n", now, msg)

	if this.Print {
		fmt.Printf(msg)
	}

	if (this.LogLevel & LOG_WARNING) != 0 {
		AppendFile(this.LogPath + this.Filename, msg)
	}
}

func (this LogTool) WriteError(msg string) {
	now :=  time.Now().Format(time.ANSIC)
	msg = fmt.Sprintf("[%s][ERROR]%s\n", now, msg)

	if this.Print {
		fmt.Printf(msg)
	}

	if (this.LogLevel & LOG_ERROR) != 0 {
		AppendFile(this.LogPath + this.Filename, msg)
	}
}

func (this LogTool) WriteDebug(msg string) {
	now := time.Now().Format(time.ANSIC)
	msg = fmt.Sprintf("[%s][DEBUG]%s\n", now, msg)

	if this.Print {
		fmt.Printf(msg)
	}

	if (this.LogLevel & LOG_DEBUG) != 0 {
		AppendFile(this.LogPath + this.Filename, msg)
	}
}

// Constructor
func CreateLogger(filename string) *LogTool {
	var logger *LogTool = new(LogTool)
	logger.Filename = filename

	// Set defaults 
	logger.LogPath = "/var/guardian/misc/"
	logger.LogLevel = LOG_INFO | LOG_WARNING | LOG_ERROR
	logger.AutoRotate = true
	logger.Print = true

	return logger
}

func AppendFile(filename string, msg string) {
	stream, err := os.OpenFile(filename, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0755)

	if err != nil {
		now := time.Now().Format(time.ANSIC)
		fmt.Printf("[%s][FAIL]Error: %s", now, err)
	} else {
		_, err := stream.Write([]byte(msg))

		if err != nil {
			now := time.Now().Format(time.ANSIC)
			fmt.Printf("[%s][FAIL]Error: %s", now, err)
		}
	}

	stream.Close()
}