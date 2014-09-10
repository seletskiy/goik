package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

var (
	reTimeStartTz  = regexp.MustCompile(`DTSTART;TZID="([^"]+)":(.+)$`)
	reTimeStartUtc = regexp.MustCompile(`DTSTART:(.+)Z$`)
	reOrganizer    = regexp.MustCompile(`ORGANIZER:(?:MAILTO:)?(.+)`)
	reSummary      = regexp.MustCompile(`SUMMARY:(.+)`)
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	buffer := ""

	for {
		line := readBreakedLine(reader, &buffer)
		if line == "" {
			break
		}

		switch true {
		case reTimeStartTz.MatchString(line):
			matches := reTimeStartTz.FindStringSubmatch(line)
			t, err := time.Parse(`20060102T150405`, matches[2])
			if err != nil {
				panic(err)
			}

			tz := matches[1]

			showCal(t)

			fmt.Printf("Time :: %s (in <%s>)\n", t.Format("15:04"), tz)
		case reTimeStartUtc.MatchString(line):
			matches := reTimeStartUtc.FindStringSubmatch(line)
			t, err := time.Parse(`20060102T150405`, matches[1])
			if err != nil {
				continue
			}

			showCal(t)

			fmt.Println("Time ::", t.In(time.Local).Format("15:04"),
				"(in local time)")
		case reOrganizer.MatchString(line):
			matches := reOrganizer.FindStringSubmatch(line)
			fmt.Println("Organizer ::", matches[1])
		case reSummary.MatchString(line):
			matches := reSummary.FindStringSubmatch(line)
			fmt.Println("Summary ::", matches[1])

		}
	}
}

func readBreakedLine(reader *bufio.Reader, buffer *string) (line string) {
	for {
		partialLine, err := reader.ReadString('\n')
		if err != nil || partialLine[0] != ' ' && *buffer != "" {
			line = strings.TrimSpace(*buffer)
			*buffer = strings.TrimSpace(partialLine)
			break
		} else {
			if partialLine[0] == ' ' {
				partialLine = partialLine[1:]
			}
			*buffer += partialLine
		}
	}

	return line
}

func showCal(datetime time.Time) {
	cmd := exec.Command("cal",
		"--color=never",
		fmt.Sprintf("%d", datetime.Day()),
		fmt.Sprintf("%d", datetime.Month()),
		fmt.Sprintf("%d", datetime.Year()))

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, _ := cmd.StdoutPipe()
	err := cmd.Start()
	if err != nil {
		panic(err)
	}

	cal, err := ioutil.ReadAll(stdout)
	if err != nil {
		panic(err)
	}

	reNewLine := regexp.MustCompile(`(?m)^`)
	out := reNewLine.ReplaceAllString(string(cal), " ")
	reDay := regexp.MustCompile(fmt.Sprintf(`(^| )%2d( |$)`, datetime.Day()))
	out = reDay.ReplaceAllString(out, fmt.Sprintf(`[%2d]`, datetime.Day()))
	fmt.Println(out)
}
