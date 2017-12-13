package utils

// CommandArgs commandline args
type CommandArgs struct {
	Mode          string
	FileName      string
	Encoding      string
	FileSubTop    string
	FileSubBottom string
	EncSubTop     string
	EncSubBottom  string
	Port          string
}

// FileInfo
type FileInfo struct {
	FileName string
	Encoding string
}
