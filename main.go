package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"tubtitle/utils"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("use -help for available options.")
		return
	}

	c := parseCmd()
	fmt.Printf("%+v\n", c)
	Run(c)
}

func parseCmd() utils.CommandArgs {
	var mode = flag.String("mode", "", "modes of operations : "+strings.Join(GetValidModes(), ", ")[1:])
	var fileName = flag.String("f", "", "file name")
	var fUp = flag.String("f1", "", "subtitle for the top side of the screen")
	var fDown = flag.String("f2", "", "subtitle for the bottom side of the screen")
	var encTop = flag.String("e1", "", "encoding for subtitle (the top side of the screen)")
	var encBottom = flag.String("e2", "", "encoding for subtitle (the bottom side of the screen)")
	var enc = flag.String("enc", "", "encoding type, 'pl' for Polish, 'tr' for Turkish. i.e. -enc pl")
	var port = flag.String("p", "3000", "port number for server.")
	flag.Parse()
	return utils.CommandArgs{Mode: *mode, FileName: *fileName, Encoding: *enc, FileSubTop: *fUp, FileSubBottom: *fDown, EncSubTop: *encTop, EncSubBottom: *encBottom, Port: *port}
}
