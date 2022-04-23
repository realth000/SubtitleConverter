package format

import (
	. "SubtitleConverter/format/base_format"
	"SubtitleConverter/format/lrc"
	"bufio"
	"fmt"
	"os"
)

func ParseLrc(filePath string) []SubtitleFormat {
	lrcFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return []SubtitleFormat{}
	}
	lrcScanner := bufio.NewScanner(lrcFile)

	var result []SubtitleFormat

	// Handle the first line.
	if lrcScanner.Scan() {
		firstTime, firstData := lrc.ParseLrcLine(lrcScanner.Text())
		result = append(result, SubtitleFormat{
			Time: TimeFormat{StartTime: firstTime.BaseTime},
			Data: firstData,
		})
	}

	// Handle middle lines.
	var lastLineBehind = 0
	for lrcScanner.Scan() {
		t, d := lrc.ParseLrcLine(lrcScanner.Text())
		var s = SubtitleFormat{Time: TimeFormat{StartTime: t.BaseTime}, Data: d}
		result = append(result, s)
		result[lastLineBehind].Time.EndTime = t.BaseTime
		lastLineBehind++
	}

	// Handle the last line.
	result[lastLineBehind].Time.EndTime = BaseTime{
		Hour: "99",
		Min:  "59",
		Sec:  "59",
		MSec: "99",
	}

	return result
}
