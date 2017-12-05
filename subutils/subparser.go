package subutils

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

// SubtitleEntry : simple subtitle entry
type SubtitleEntry struct {
	StartTime string
	EndTime   string
	Text      string
}

func (s SubtitleEntry) String() string {
	return fmt.Sprintf("%s %s %s", s.StartTime, s.EndTime, s.Text)
}

// 00:01:01,833
var reTime = regexp.MustCompile("[0-9]+:[0-9]+:[0-9]+,[0-9]+")
var reText = regexp.MustCompile("[a-z]+")

// ParseSub : creates subtitle entries from text and writes to a file
func ParseSub(c CommandArgs) {
	if len(c.FileName) == 0 {
		return
	}

	txt := readFile(FileInfo{FileName: c.FileName, Encoding: c.Encoding})
	subs := CreateSubEntries(txt)
	var buffer bytes.Buffer
	for _, item := range subs {
		buffer.WriteString(item.String())
		buffer.WriteString("\n")
	}
	writeToFile(c.FileName+".parsed", buffer.String())
}

// CreateSubEntries : creates subtitle entries from text
func CreateSubEntries(text string) []SubtitleEntry {
	lines := filterText(text)
	time := lines[0]
	var subs []SubtitleEntry
	var buffer bytes.Buffer
	lineCounter := 0
	for _, item := range lines[1:] {
		if isTimeLine(item) {
			times := parseTimes(time)
			subs = append(subs, createSubtitleEntry(times, buffer.String()))
			time = item
			buffer.Reset()
			lineCounter = 0
			continue
		}
		// FIXME: fix formatting issue, i.e. next line combined with previous one
		if lineCounter > 1 {
			buffer.WriteString("\\N") // new line in subtitle form
		}
		buffer.WriteString(item)
		lineCounter++
	}
	times := parseTimes(time)
	subs = append(subs, createSubtitleEntry(times, buffer.String()))
	return subs
}

func filterText(text string) []string {
	allLines := strings.Split(text, "\n")
	filtered := make([]string, 1)
	for _, item := range allLines {
		tmp := strings.TrimSpace(item)
		if isTimeLine(tmp) || isTextLine(tmp) {
			filtered = append(filtered, tmp)
		}
	}

	return filtered[2:]
}

func createSubtitleEntry(times []string, text string) SubtitleEntry {
	t1, t2 := getFormattedTime(times)
	start := strings.Replace(t1, ",", ".", -1)
	end := strings.Replace(t2, ",", ".", -1)
	return SubtitleEntry{start, end, text}
}

func getFormattedTime(times []string) (string, string) {
	if len(times) < 2 {
		return "", ""
	}
	t1, t2 := times[0], times[1]
	t1, t2 = t1[:len(t1)-1], t2[:len(t2)-1]
	return t1, t2
}

func isTimeLine(line string) bool {
	return len(reTime.FindAllString(line, -1)) > 0
}

func isTextLine(line string) bool {
	return len(reText.FindAllString(line, -1)) > 0
}

func parseTimes(time string) []string {
	return reTime.FindAllString(time, 2)
}
