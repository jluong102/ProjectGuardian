package main

type Master struct {
	LogPath string `json:"log_path"` // Optional
	Debug bool `json:"debug"` // Optional
	NoLog bool `json:"no_log"` // Optional
	NoInfo bool `json:"no_info"` // Optional
	NoSuccess bool `json:"no_success"` // Optional
	NoWarning bool `json:"no_warning"` // Optional
	NoError bool `json:"no_error"` // Optional
	NoConsole bool `json:"no_console"` // Optional
	Watches map[string]*Watch `json:"watches"` // Required
}

type Watch struct {
	Name string `json:"name"` // Auto
	Interval uint32 `json:"interval"` // Required
	LogPath string `json:"log_path"` // Optional
	Debug bool `json:"debug"` // Optional
	NoLog bool `json:"no_log"` // Optional
	NoInfo bool `json:"no_info"` // Optional
	NoSuccess bool `json:"no_success"` // Optional
	NoWarning bool `json:"no_warning"` // Optional
	NoError bool `json:"no_error"` // Optional
	NoConsole bool `json:"no_console"` // Optional
	SuccessCodes []int32 `json:"success_codes"` // Required
	FailureCodes []int32 `json:"failure_codes"` // Required
	CheckScript string `json:"check_script"` // Required
}

type Remedies struct {

}