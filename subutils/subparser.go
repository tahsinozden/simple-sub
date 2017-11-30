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
	for _, item := range lines[2:] {
		if isTimeLine(item) {
			times := parseTimes(time)
			subs = append(subs, SubtitleEntry{times[0], times[1], buffer.String()})
			time = item
			buffer.Reset()
			continue
		}
		buffer.WriteString(item)
	}
	times := parseTimes(time)
	subs = append(subs, SubtitleEntry{times[0], times[1], buffer.String()})
	return subs
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
