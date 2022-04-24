package lrc

import (
	"SubtitleConverter/format/base_format"
	"errors"
	"fmt"
	"regexp"
)

type LrcTime struct {
	base_format.BaseTime
}

var (
	ParseRegexp      = regexp.MustCompile(`^(?P<Time>\[\d{2}:\d{2}\.\d{2}\])(?P<Data>.*)`)
	TimeParseRegexp  = regexp.MustCompile(`(?P<Min>\d{2}):(?P<Sec>\d{2})\.(?P<MSec>\d{2})`)
	errorInvalidTime = `error: invalid lrc format time line`
	errorInvalidLine = `error: invalid lrc format line`
)

func (l *LrcTime) FromLrcTime(lrcTime string) error {

	matches := TimeParseRegexp.FindStringSubmatch(lrcTime)
	if len(matches) < TimeParseRegexp.NumSubexp()+1 {
		return errors.New(errorInvalidTime)
	}
	l.Hour = "00"
	l.Min = matches[TimeParseRegexp.SubexpIndex("Min")]
	l.Sec = matches[TimeParseRegexp.SubexpIndex("Sec")]
	l.MSec = matches[TimeParseRegexp.SubexpIndex("MSec")]
	return nil
}

func (l *LrcTime) IsEmpty() bool {
	if l.Hour == "" && l.Min == "" && l.Sec == "" && l.MSec == "" {
		return true
	}
	return false
}

func ParseLrcLine(lrc string) (LrcTime, string, error) {
	matches := ParseRegexp.FindStringSubmatch(lrc)
	if len(matches) < 2 {
		// TODO: Parse a invalid format line here.
		// fmt.Println("error parsing lrc line: invalid format")
		return LrcTime{}, "", errors.New(errorInvalidTime)
	}
	var l LrcTime
	err := l.FromLrcTime(matches[ParseRegexp.SubexpIndex("Time")])
	if err != nil {
		return LrcTime{}, "", errors.New(errorInvalidLine)
	}
	d := matches[ParseRegexp.SubexpIndex("Data")]
	return l, d, nil
}

func ToLrcTime(base base_format.BaseTime) string {
	return fmt.Sprintf("[%s]", base.ToLrcFormat())
}
