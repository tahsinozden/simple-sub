package subutils

import (
	"bytes"
	"fmt"
	"log"
)

const metadata = `
[Script Info]
Title: MovieTitle
Original Script: simple-sub
ScriptType: v1.00
Collisions: Normal

[V4 Styles]
Format: Name,Fontname,Fontsize,PrimaryColour,SecondaryColour,TertiaryColour,BackColour,Bold,Italic,BorderStyle,Outline,Shadow,Alignment,MarginL,MarginR,MarginV,AlphaLevel,Encoding
Style: StyleA,Arial,20,&H00FFFFFF,&H00FFFFFF,0,0,0,0,1,2,0,6,30,30,10,0,0
Style: StyleB,Arial,20,&H00FFFFFF,&H00FFFFFF,0,0,0,0,1,2,0,2,30,30,10,0,0

[Events]
Format: Marked, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text`

const subTopLineTemplate = "Dialogue: Marked=0,%s,%s,StyleA,NTP,0000,0000,0000,!Effect,%s \n"
const subBottomLineTemplate = "Dialogue: Marked=0,%s,%s,StyleB,NTP,0000,0000,0000,!Effect,%s \n"

// MergeSubtitles : merges subtitles and creates a new file.
func MergeSubtitles(c CommandArgs) {
	if !hasAllSubMergeParams(c) {
		log.Fatal("Missing merge params!")
	}

	subTop, subBottom := createEntries(c)
	merged := Merge(subTop, subBottom)
	writeToFile(c.FileSubTop+".merged.ssa", merged)
}

// Merge : merges two subtitles
func Merge(subUp []SubtitleEntry, subDown []SubtitleEntry) string {
	var buffer bytes.Buffer
	buffer.WriteString(metadata + "\n")
	mergeSub(&buffer, subTopLineTemplate, subUp)
	mergeSub(&buffer, subBottomLineTemplate, subDown)
	return buffer.String()
}

func mergeSub(b *bytes.Buffer, template string, entries []SubtitleEntry) {
	for _, item := range entries {
		b.WriteString(fmt.Sprintf(template, item.StartTime, item.EndTime, item.Text))
	}
}

func createEntries(c CommandArgs) ([]SubtitleEntry, []SubtitleEntry) {
	return createEntry(c.FileSubTop, c.EncSubTop), createEntry(c.FileSubBottom, c.EncSubBottom)
}

func createEntry(fileName string, enc string) []SubtitleEntry {
	if len(fileName) == 0 {
		return []SubtitleEntry{}
	}

	lines := readFile(FileInfo{FileName: fileName, Encoding: enc})
	return CreateSubEntries(lines)
}
