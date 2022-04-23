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

type TimeFormat struct {
	StartTime BaseTime
	EndTime   BaseTime
}

type SubtitleFormat struct {
	Time TimeFormat
	Data string
}
