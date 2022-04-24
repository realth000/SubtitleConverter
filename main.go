package main

import (
	"SubtitleConverter/format"
	"fmt"
)

const (
	testLrcFile = `./tmp/test_lrc.lrc`
	testSrtFile = `./tmp/test_srt.srt`
)

func main() {
	//result := format.ParseLrc(testLrcFile)
	//// Test lrc output.
	//for _, r := range result {
	//	fmt.Printf("%s -> %s : %s\n", r.Time.StartTime.ToString(), r.Time.EndTime.ToString(), r.Data)
	//}
	//
	//result2, err := format.ParseSrt(testSrtFile)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//// Test srt output.
	//for _, r := range result2 {
	//	fmt.Printf("%s -> %s : %s\n", r.Time.StartTime.ToString(), r.Time.EndTime.ToString(), r.Data)
	//}

	result := format.ToSrt(format.ParseLrc(testLrcFile))
	// Test lrc output.
	for _, r := range result {
		fmt.Println(r)
	}

	t, _ := format.ParseSrt(testSrtFile)
	result2 := format.ToLrc(t)
	for _, r := range result2 {
		fmt.Println(r)
	}

}
