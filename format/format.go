package format

import (
	. "SubtitleConverter/format/base_format"
	"SubtitleConverter/format/lrc"
	"SubtitleConverter/format/srt"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	errorInvalidIndex = `error: invalid srt line index`
	errorInvalidTime  = `error: invalid srt format time`
	errorEOF          = `error: unexpected EOF`
)

var (
	srtLineIndex = 0
)

func ParseLrc(filePath string) []SubtitleFormat {
	lrcFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return []SubtitleFormat{}
	}
	lrcScanner := bufio.NewScanner(lrcFile)

	var result []SubtitleFormat
	var lrcLineIndex = 0

	// Handle the first line.
	if lrcScanner.Scan() {
		lrcLineIndex++
		firstTime, firstData, err := lrc.ParseLrcLine(lrcScanner.Text())
		if err != nil && firstData != "" {
			// First line failed.
			fmt.Println(err)
			return []SubtitleFormat{}
		}
		result = append(result, SubtitleFormat{
			Index: lrcLineIndex,
			Time:  TimeFormat{StartTime: firstTime.BaseTime},
			Data:  firstData,
		})
	}

	// Handle middle lines.
	var lastLineBehind = 0
	for lrcScanner.Scan() {
		lrcLineIndex++
		t, d, err := lrc.ParseLrcLine(lrcScanner.Text())
		if err != nil && d != "" {
			fmt.Println(err)
			continue
		}
		var s = SubtitleFormat{Index: lrcLineIndex, Time: TimeFormat{StartTime: t.BaseTime}, Data: d}
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

func makeParseSrtError(errorType string) error {
	fmt.Println(srtLineIndex)
	return errors.New(errorType + " at line " + string(srtLineIndex))
}

func ParseSrt(filePath string) ([]SubtitleFormat, error) {
	lrcFile, err := os.Open(filePath)
	if err != nil {
		return []SubtitleFormat{}, err
	}
	srtScanner := bufio.NewScanner(lrcFile)
	srtLineIndex = 0
	var result []SubtitleFormat
	for srtScanner.Scan() {
		var s SubtitleFormat
		var t srt.SrtTime

		// Check if is an index line.
		index, err := strconv.Atoi(srtScanner.Text())
		if err != nil {
			return []SubtitleFormat{}, makeParseSrtError(errorInvalidIndex)
		}
		s.Index = index

		if !srtScanner.Scan() {
			return []SubtitleFormat{}, makeParseSrtError(errorEOF)
		}
		srtLineIndex++

		// Check if is a srt format time.
		if err := t.FromSrtTime(srtScanner.Text()); err != nil {
			return []SubtitleFormat{}, makeParseSrtError(errorInvalidTime)
		}
		s.Time.StartTime = t.StartTime
		s.Time.EndTime = t.EndTime

		// Store the next line as data.
		if !srtScanner.Scan() {
			return []SubtitleFormat{}, makeParseSrtError(errorEOF)
		}
		srtLineIndex++
		s.Data += srtScanner.Text()
		// TODO: Whether multiple lines of data exist?
		// Try to store most lines of data.
		for srtScanner.Scan() {
			srtLineIndex++
			if srtScanner.Text() == "" {
				break
			}
			s.Data += srtScanner.Text()
		}

		result = append(result, s)
	}

	return result, nil
}

func ToLrc(subs []SubtitleFormat) []string {
	var result []string
	for _, sub := range subs {
		result = append(result, fmt.Sprintf("%s%s", lrc.ToLrcTime(sub.Time.StartTime), sub.Data))
	}
	return result
}

func ToSrt(subs []SubtitleFormat) []string {
	var result []string
	for _, sub := range subs {
		result = append(result, strconv.Itoa(sub.Index), srt.ToSrtTime(sub.Time), sub.Data, "")
	}

	return result
}
