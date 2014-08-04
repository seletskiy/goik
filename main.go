package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	reBegin     = regexp.MustCompile(`BEGIN:VCALENDAR`)
	reEnd       = regexp.MustCompile(`END:VCALENDAR`)
	reTimeStart = regexp.MustCompile(`DTSTART:(.+)`)
	reTimeEnd   = regexp.MustCompile(`DTEND:(.+)`)
	reOrganizer = regexp.MustCompile(`ORGANIZER:(?:MAILTO:)?(.+)`)
	reSummary   = regexp.MustCompile(`SUMMARY:(.+)`)
)

func main() {
	var (
		reader = bufio.NewReader(os.Stdin)
		buffer = ""
		line   = ""
	)

	for {
		for {
			partialLine, err := reader.ReadString('\n')
			if err != nil {
				break
			}

			if partialLine[0] != ' ' && buffer != "" {
				line, buffer = strings.TrimSpace(buffer), partialLine
				break
			} else {
				buffer += partialLine
			}
		}

		if line == "" {
			break
		}

		switch true {
		case reTimeStart.MatchString(line):
			matches := reTimeStart.FindStringSubmatch(line)
			t, err := time.Parse(`20060102T150405Z`, matches[1])
			if err != nil {
				panic(err)
			}

			fmt.Println(t)
		}
	}
}
