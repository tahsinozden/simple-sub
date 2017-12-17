package subutils

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"

	"github.com/golang/glog"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"tubtitle/utils"
)

var validEncodings = map[string]*charmap.Charmap{
	"pl": charmap.Windows1250,
	"tr": charmap.Windows1254,
}

func RemoveAccentLetters(c utils.CommandArgs) {
	if len(c.FileName) == 0 || len(c.Encoding) == 0 {
		glog.Error("Missing remove-accent params!")
		return
	}

	txt := readWithEncoding(c.FileName, GetEncoding(c.Encoding))
	writeToFile(c.FileName+".accents-removed.srt", txt)
}

func GetEncoding(enc string) *charmap.Charmap {
	if enc, ok := validEncodings[enc]; ok {
		return enc
	}
	return charmap.ISO8859_1
}

func readFile(f utils.FileInfo) string {
	if len(f.FileName) == 0 {
		return ""
	}

	if len(f.Encoding) > 0 {
		return readWithEncoding(f.FileName, GetEncoding(f.Encoding))
	}
	return simpleRead(f.FileName)
}

func readWithEncoding(filename string, charmap *charmap.Charmap) string {
	var buffer bytes.Buffer
	f, err := os.Open(filename)
	if err != nil {
		glog.Fatal("couldn't open file : ", filename)
	}
	defer f.Close()
	r := transform.NewReader(f, charmap.NewDecoder())
	sc := bufio.NewScanner(r)

	for sc.Scan() {
		buffer.WriteString(sc.Text())
		buffer.WriteString("\n")
	}
	return GetConvertAccentText(buffer.String())
}

func simpleRead(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		glog.Fatal("couldn't open file : ", filename)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	var buffer bytes.Buffer

	for sc.Scan() {
		buffer.WriteString(sc.Text())
		buffer.WriteString("\n")
	}
	return GetConvertAccentText(buffer.String())
}

func writeToFile(fileName string, text string) {
	perm := os.FileMode(0644)
	ioutil.WriteFile(fileName, []byte(text), perm)
}

func GetConvertAccentText(text string) string {
	var buffer bytes.Buffer
	for _, runeValue := range text {
		converted := Convert2NonAccent(string(runeValue))
		buffer.WriteString(converted)
	}
	return buffer.String()
}

func hasAllSubMergeParams(c utils.CommandArgs) bool {
	return len(c.FileSubTop) > 0 && len(c.FileSubBottom) > 0
}
