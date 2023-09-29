package main

type Master struct {
	LogPath string `json:"log_path"` // Optional
	Debug bool `json:"debug"` // Optional
	NoLog bool `json:"no_log"` // Optional
	NoInfo bool `json:"no_info"` // Optional
	NoSuccess bool `json:"no_success"` // Optional
	NoWarning bool `json:"no_warning"` // Optional
	NoError bool `json:"no_error"` // Optional
}