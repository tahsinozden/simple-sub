package main

import (
	"github.com/golang/glog"
	"simple-sub/subutils"
	"simple-sub/utils"
	"simple-sub/webapi"
)

var validModes = map[string]func(c utils.CommandArgs){
	"remove-accent": subutils.RemoveAccentLetters,
	"parse":         subutils.ParseSub,
	"merge":         subutils.MergeSubtitles,
	"serve":         webapi.Serve,
}

// GetValidModes returns the valid modes
func GetValidModes() []string {
	m := make([]string, 1)
	for k := range validModes {
		m = append(m, k)
	}
	return m
}

// Run runs in accordance with commandline arguments
func Run(c utils.CommandArgs) {
	fn, ok := validModes[c.Mode]
	if !ok {
		glog.Error("provide a valid mode. Valid modes : ", GetValidModes())
		return
	}
	fn(c)
}
