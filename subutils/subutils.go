package subutils

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

// CommandArgs : commandline args
type CommandArgs struct {
	Mode     string
	FileName string
	Encoding string
	FileUp   string
	FileDown string
}

var validModes = map[string]func(c CommandArgs){
	"remove-accent": removeAccentLetters,
	"parse":         parseSub,
	"merge":         mergeSubtitles,
}

var validEncodings = map[string]*charmap.Charmap{
	"pl": charmap.Windows1250,
	"tr": charmap.Windows1254,
}

// GetValidModes : returns the valid modes
func GetValidModes() []string {
	m := make([]string, 1)
	for k := range validModes {
		m = append(m, k)
	}
	return m
}

// Run : runs in accordance with commandline arguments
func (c *CommandArgs) Run() {
	if fn, ok := validModes[c.Mode]; ok {
		fn(*c)
	} else {
		fmt.Println("provide a valid mode. Valid modes : ", GetValidModes())
	}
}

func (c CommandArgs) String() string {
	return fmt.Sprintf("Mode: %s\nFileName: %s\nEncoding: %s\nFileUp: %s\nFileDown: %s", c.Mode, c.FileName, c.Encoding, c.FileUp, c.FileDown)
}

func removeAccentLetters(commandArgs CommandArgs) {
	if len(commandArgs.FileName) > 0 && len(commandArgs.Encoding) > 0 {
		txt := readWithEncoding(commandArgs.FileName, getEncoding(commandArgs.Encoding))
		writeToFile(commandArgs.FileName+".new", txt)
	}
}

func getEncoding(cmdStr string) *charmap.Charmap {
	if enc, ok := validEncodings[cmdStr]; ok {
		return enc
	}
	return charmap.ISO8859_1
}

func readWithEncoding(filename string, charmap *charmap.Charmap) string {
	var buffer bytes.Buffer
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := transform.NewReader(f, charmap.NewDecoder())
	sc := bufio.NewScanner(r)

	for sc.Scan() {
		buffer.WriteString(sc.Text())
		buffer.WriteString("\n")
	}
	return getConvertAccentText(buffer.String())
}

func writeToFile(fileName string, text string) {
	perm := os.FileMode(0644)
	ioutil.WriteFile(fileName, []byte(text), perm)
}

func getConvertAccentText(text string) string {
	var buffer bytes.Buffer
	for _, runeValue := range text {
		buffer.WriteString(Convert2NonAccent(string(runeValue)))
	}
	return buffer.String()
}

func parseSub(commandArgs CommandArgs) {
	if len(commandArgs.FileName) > 0 && len(commandArgs.Encoding) > 0 {
		txt := readWithEncoding(commandArgs.FileName, getEncoding(commandArgs.Encoding))
		subs := CreateSubEntries(txt)
		var buffer bytes.Buffer
		for _, item := range subs {
			buffer.WriteString(item.String())
			buffer.WriteString("\n")
		}
		writeToFile(commandArgs.FileName+".parsed", buffer.String())
	}
}

func mergeSubtitles(c CommandArgs) {
	if len(c.FileUp) > 0 && len(c.FileDown) > 0 && len(c.Encoding) > 0 {
		txt := readWithEncoding(c.FileUp, getEncoding(c.Encoding))
		subUp := CreateSubEntries(txt)
		txt = readWithEncoding(c.FileDown, getEncoding(c.Encoding))
		subDown := CreateSubEntries(txt)
		merged := Merge(subUp, subDown)
		writeToFile(c.FileUp+"merged.ssa", merged)
	}
}
