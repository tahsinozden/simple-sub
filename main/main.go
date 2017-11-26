package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"simple-sub/converter"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

// for Polish
var windows1250 = charmap.Windows1250

// for Turkish
var windows1254 = charmap.Windows1254

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		fileName := args[0]
		fileName, enc := parseCmd()
		if len(fileName) > 0 && len(enc) > 0 {
			fmt.Println(enc, "is selected.")
			txt := readWithEncoding(fileName, getEncoding(enc))
			writeToFile(fileName+".new", txt)
		}
	} else {
		fmt.Println("use -help")
	}

}

func parseCmd() (string, string) {
	var fileName = flag.String("f", "", "file name")
	var enc = flag.String("enc", "", "encoding type, 'pl' for Polish, 'tr' for Turkish. i.e. -enc pl")
	flag.Parse()
	return *fileName, *enc
}

func getEncoding(cmdStr string) *charmap.Charmap {
	switch cmdStr {
	case "pl":
		return windows1250
	case "tr":
		return windows1254
	default:
		return windows1250
	}
}

func readWithEncoding(filename string, charmap *charmap.Charmap) string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	r := transform.NewReader(f, charmap.NewDecoder())

	sc := bufio.NewScanner(r)
	allText := ""
	for sc.Scan() {
		allText += sc.Text() + "\n"
	}
	defer f.Close()
	return getConvertAccentText(allText)
}

func writeToFile(fileName string, text string) {
	perm := os.FileMode(0644)
	ioutil.WriteFile(fileName, []byte(text), perm)
}

func getConvertAccentText(text string) string {
	newText := ""
	for _, runeValue := range text {
		newText += converter.Convert2NonAccent(string(runeValue))
	}
	return newText
}
