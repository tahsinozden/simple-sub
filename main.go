package main

import (
	"flag"
	"fmt"
	"os"
	"simple-sub/subutils"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		commandArgs := parseCmd()
		fmt.Println(commandArgs)
		commandArgs.Run()
	} else {
		fmt.Println("use -help")
	}
}

func parseCmd() subutils.CommandArgs {
	var mode = flag.String("mode", "", "modes of operations : "+strings.Join(subutils.GetValidModes(), ", ")[1:])
	var fileName = flag.String("f", "", "file name")
	var fUp = flag.String("fup", "", "subtitle for the top side of the screen")
	var fDown = flag.String("fdown", "", "subtitle for the downside of the screen")
	var enc = flag.String("enc", "", "encoding type, 'pl' for Polish, 'tr' for Turkish. i.e. -enc pl")
	flag.Parse()
	return subutils.CommandArgs{Mode: *mode, FileName: *fileName, Encoding: *enc, FileUp: *fUp, FileDown: *fDown}
}
