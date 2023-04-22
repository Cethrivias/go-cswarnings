package parser

import (
	"bufio"
	"cs-warnings/types"
	"io"
	"regexp"
	"sync"
)

var wg sync.WaitGroup
var warningRegexp = regexp.MustCompile(`^(?P<path>.*): warning (?P<code>\w*): (?P<description>.*) \[(?P<project>.*)]$`)

func ParseWarnings(r io.Reader) map[string]types.Warning {
	scanner := bufio.NewScanner(r)
	warnings := make(chan types.Warning)
	result := make(chan map[string]types.Warning)

	go mapWarnings(warnings, result)
	for scanner.Scan() {
		wg.Add(1)
		text := scanner.Text()

		go parseLine(text, warnings)
	}

	wg.Wait()
	close(warnings)

	return <-result
}

func mapWarnings(warnings chan types.Warning, warningsChan chan map[string]types.Warning) {
	result := make(map[string]types.Warning)
	for warn := range warnings {
		result[warn.Path+warn.Code] = warn
	}

	warningsChan <- result
	close(warningsChan)
}

func parseLine(line string, r chan types.Warning) {
	match := warningRegexp.FindStringSubmatch(line)
	if match != nil {
		warning := types.Warning{
			Path:        match[1],
			Code:        match[2],
			Description: match[3],
			Project:     match[4],
		}
		r <- warning
	}
	wg.Done()
}
