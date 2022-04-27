package cli

import (
	"SubtitleConverter/format"
	"SubtitleConverter/format/base_format"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Flags.
var (
	flagSrcPath string
	flagDstPath string
	flagSrcFmt  string
	flagDstFmt  string
)

func init() {
	flag.StringVar(&flagSrcPath, "s", "", "Source file")
	flag.StringVar(&flagDstPath, "d", "", "Destination file. Print to screen if not set.")
	flag.StringVar(&flagSrcFmt, "sf", "", "-sf [lrc, srt] Specify source format")
	flag.StringVar(&flagDstFmt, "df", "", "-df [lrc, srt] Specify destination format")
	flag.Parse()
	if !checkFlag() {
		os.Exit(0)
	}
}

func checkFlag() bool {
	if flagSrcPath == "" {
		fmt.Println("error: source file not specified")
		return false
	}
	if flagSrcFmt != "" && flagSrcFmt != "lrc" && flagSrcFmt != "srt" {
		fmt.Println("invalid source format")
		return false
	}

	if flagDstFmt != "" && flagDstFmt != "lrc" && flagDstFmt != "srt" {
		fmt.Println("invalid destination format")
		return false
	}

	return true
}

func parseSource() []base_format.SubtitleFormat {
	if flagSrcFmt == "lrc" || strings.HasSuffix(flagSrcPath, ".lrc") {
		return format.ParseLrc(flagSrcPath)
	}
	if flagSrcFmt == "srt" || strings.HasSuffix(flagSrcPath, ".srt") {
		ret, err := format.ParseSrt(flagSrcPath)
		if err != nil {
			fmt.Println(err)
			return []base_format.SubtitleFormat{}
		}
		return ret
	}
	fmt.Println("error: unknown source format")
	os.Exit(0)
	return nil
}

func convert(base []base_format.SubtitleFormat) []string {
	if flagDstFmt == "lrc" || strings.HasSuffix(flagDstPath, ".lrc") {
		return format.ToLrc(base)
	}
	if flagDstFmt == "srt" || strings.HasSuffix(flagDstPath, ".srt") {
		return format.ToSrt(base)
	}
	fmt.Println("error: can not specify destination format")
	os.Exit(0)
	return nil
}

func Run() {
	base := parseSource()
	result := convert(base)
	if len(result) < 1 {
		fmt.Println("Error: convert result too short, maybe some error occurred")
		return
	}
	if flagDstPath == "" {
		for _, r := range result {
			fmt.Println(r)
		}
	} else {
		err := ioutil.WriteFile(flagDstPath, []byte(strings.Join(result, "\n")), 0644)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Success: write to file %s\n", flagDstPath)
		}
	}
}
