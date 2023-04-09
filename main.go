package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
)

func main() {
	parseWarnings(os.Stdin)
}

func parseWarnings(r io.Reader) (map[string]Warning, error) {
	scanner := bufio.NewScanner(r)
	warnings := make(map[string]Warning)
	warningRegexp := regexp.MustCompile(`^(?P<path>.*): warning (?P<code>\w*): (?P<description>.*) \[(?P<project>.*)]$`)

	for scanner.Scan() {
		text := scanner.Text()
		if warningRegexp.MatchString(text) {
			match := warningRegexp.FindStringSubmatch(text)
			warning := Warning{
				path:        match[1],
				code:        match[2],
				description: match[3],
				project:     match[4],
			}
			warnings[text] = warning
			// fmt.Printf("%+v\n", warning)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Panicln("Scanner error", err)
	}

	return warnings, nil
}

type Warning struct {
	path        string
	code        string
	description string
	project     string
}

// Sync - 6 sec
