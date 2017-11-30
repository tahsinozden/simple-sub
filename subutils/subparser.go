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

func filterText(text string) []string {
	allLines := strings.Split(text, "\n")
	filtered := make([]string, 1)
	for _, item := range allLines {
		tmp := strings.TrimSpace(item)
		if isTimeLine(tmp) || isTextLine(tmp) {
			filtered = append(filtered, tmp)
		}
	}

	return filtered
}

// CreateSubEntries : creates subtitle entries from text
func CreateSubEntries(text string) []SubtitleEntry {
	lines := filterText(text)
	time := lines[1]
	subs := []SubtitleEntry{}
	var buffer bytes.Buffer
	lineCounter := 0
	for _, item := range lines[2:] {
		if isTimeLine(item) {
			times := parseTimes(time)
			subs = append(subs, createSubtitleEntry(times, buffer.String()))
			time = item
			buffer.Reset()
			lineCounter = 0
			continue
		}
		// TODO: fix formating issue, i.e. next line combined with previous one
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

func createSubtitleEntry(times []string, text string) SubtitleEntry {
	start := strings.Replace(times[0], ",", ".", -1)
	end := strings.Replace(times[1], ",", ".", -1)
	return SubtitleEntry{start, end, text}
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
