package srt

import (
	"SubtitleConverter/format/base_format"
	"errors"
	"regexp"
)

type SrtTime struct {
	StartTime base_format.BaseTime
	EndTime   base_format.BaseTime
}

const (
	startTimeRegexpString = `^(?P<SHour>\d{2}):(?P<SMin>\d{2}):(?P<SSec>\d{2}),(?P<SMSec>\d{3})`
	timeSeparator         = ` --> `
	endTimeRegexpString   = `(?P<EHour>\d{2}):(?P<EMin>\d{2}):(?P<ESec>\d{2}),(?P<EMSec>\d{3})$`
	errorInvalidTime      = `error: invalid srt format time line`
)

var (
	TimeParseRegexp = regexp.MustCompile(startTimeRegexpString + timeSeparator + endTimeRegexpString)
)

func (s *SrtTime) FromSrtTime(srtTime string) error {
	matches := TimeParseRegexp.FindStringSubmatch(srtTime)
	// Check if line is valid.
	if len(matches) < TimeParseRegexp.NumSubexp()+1 {
		return errors.New(errorInvalidTime)
	}
	s.StartTime.Hour = matches[TimeParseRegexp.SubexpIndex("SHour")]
	s.StartTime.Min = matches[TimeParseRegexp.SubexpIndex("SMin")]
	s.StartTime.Sec = matches[TimeParseRegexp.SubexpIndex("SSec")]
	s.StartTime.MSec = matches[TimeParseRegexp.SubexpIndex("SMSec")]
	s.EndTime.Hour = matches[TimeParseRegexp.SubexpIndex("EHour")]
	s.EndTime.Min = matches[TimeParseRegexp.SubexpIndex("EMin")]
	s.EndTime.Sec = matches[TimeParseRegexp.SubexpIndex("ESec")]
	s.EndTime.MSec = matches[TimeParseRegexp.SubexpIndex("EMSec")]

	return nil
}
