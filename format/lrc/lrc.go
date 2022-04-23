package lrc

import (
	"SubtitleConverter/format/base_format"
	"regexp"
)

type LrcTime struct {
	base_format.BaseTime
}

var (
	ParseRegexp     = regexp.MustCompile(`^(?P<Time>\[\d{2}:\d{2}\.\d{2}\])(?P<Data>.*)`)
	TimeParseRegexp = regexp.MustCompile(`(?P<Min>\d{2}):(?P<Sec>\d{2})\.(?P<MSec>\d{2})`)
)

func (b *LrcTime) FromLrcTime(lrcTime string) {

	b.Hour = "00"
	matches := TimeParseRegexp.FindStringSubmatch(lrcTime)
	b.Min = matches[TimeParseRegexp.SubexpIndex("Min")]
	b.Sec = matches[TimeParseRegexp.SubexpIndex("Sec")]
	b.MSec = matches[TimeParseRegexp.SubexpIndex("MSec")]
}

func ParseLrcLine(lrc string) (time LrcTime, data string) {
	matches := ParseRegexp.FindStringSubmatch(lrc)
	if len(matches) < 2 {
		// TODO: Parse a invalid format line here.
		// fmt.Println("error parsing lrc line: invalid format")
		return LrcTime{}, ""
	}
	var b LrcTime
	b.FromLrcTime(matches[ParseRegexp.SubexpIndex("Time")])
	d := matches[ParseRegexp.SubexpIndex("Data")]
	return b, d
}
