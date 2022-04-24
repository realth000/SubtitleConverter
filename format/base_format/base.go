package base_format

import "fmt"

type BaseTime struct {
	Hour string
	Min  string
	Sec  string
	MSec string
}

func (b *BaseTime) ToString() string {
	return fmt.Sprintf("%s:%s:%s.%s", b.Hour, b.Min, b.Sec, b.MSec)
}

func (b *BaseTime) ToLrcFormat() string {
	return fmt.Sprintf("%s:%s.%s", b.Min, b.Sec, b.MSec)
}

func (b *BaseTime) ToLrcFormatWithHour() string {
	return fmt.Sprintf("%s:%s:%s.%s", b.Hour, b.Min, b.Sec, b.MSec)
}

func (b *BaseTime) ToSrtFormat() string {
	return fmt.Sprintf("%s:%s:%s,%s", b.Hour, b.Min, b.Sec, b.MSec)
}

type TimeFormat struct {
	StartTime BaseTime
	EndTime   BaseTime
}

type SubtitleFormat struct {
	Index int // for srt format.
	Time  TimeFormat
	Data  string
}
