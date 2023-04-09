package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sync"
)

var wg sync.WaitGroup

func main() {
	warnings := parseWarnings(os.Stdin)
	fmt.Printf("Warnings count: %v\n", len(warnings))
}

func parseWarnings(r io.Reader) map[string]Warning {
	scanner := bufio.NewScanner(r)
	warnings := make(chan Warning)
	result := make(chan map[string]Warning)

	go mapWarnings(warnings, result)
	for scanner.Scan() {
		wg.Add(1)
		text := scanner.Text()

		go parseLine(text, warnings)
	}

	wg.Wait()
	close(warnings)

	if err := scanner.Err(); err != nil {
		log.Panicln("Scanner error", err)
	}

	return <-result
}

func mapWarnings(warnings chan Warning, warningsChan chan map[string]Warning) {
	result := make(map[string]Warning)
	for warn := range warnings {
		result[warn.path+warn.code] = warn
	}

	warningsChan <- result
	close(warningsChan)
}

func parseLine(line string, r chan Warning) {
	warningRegexp := regexp.MustCompile(`^(?P<path>.*): warning (?P<code>\w*): (?P<description>.*) \[(?P<project>.*)]$`)
	match := warningRegexp.FindStringSubmatch(line)
	if match != nil {
		warning := Warning{
			path:        match[1],
			code:        match[2],
			description: match[3],
			project:     match[4],
		}
		r <- warning
	}
	wg.Done()
}

type Warning struct {
	path        string
	code        string
	description string
	project     string
}

// Sync - 6 sec
