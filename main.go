package main

import (
	"SubtitleConverter/cli"
)

const (
	testLrcFile = `./tmp/test_lrc.lrc`
	testSrtFile = `./tmp/test_srt.srt`
)

func main() {
	cli.Run()
}
