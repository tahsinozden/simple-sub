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
		c := parseCmd()
		fmt.Printf("%+v\n", c)
		c.Run()
	} else {
		fmt.Println("use -help")
	}
}

func parseCmd() subutils.CommandArgs {
	var mode = flag.String("mode", "", "modes of operations : "+strings.Join(subutils.GetValidModes(), ", ")[1:])
	var fileName = flag.String("f", "", "file name")
	var fUp = flag.String("f1", "", "subtitle for the top side of the screen")
	var fDown = flag.String("f2", "", "subtitle for the bottom side of the screen")
	var encTop = flag.String("e1", "", "encoding for subtitle (the top side of the screen)")
	var encBottom = flag.String("e2", "", "encoding for subtitle (the bottom side of the screen)")
	var enc = flag.String("enc", "", "encoding type, 'pl' for Polish, 'tr' for Turkish. i.e. -enc pl")
	flag.Parse()
	return subutils.CommandArgs{Mode: *mode, FileName: *fileName, Encoding: *enc, FileSubTop: *fUp, FileSubBottom: *fDown, EncSubTop: *encTop, EncSubBottom: *encBottom}
}
