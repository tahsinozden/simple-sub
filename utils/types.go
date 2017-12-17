package utils

import (
	"bytes"
)

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

type SubtitleForm struct {
	Name string
	File *bytes.Buffer
	Enc  string
}
