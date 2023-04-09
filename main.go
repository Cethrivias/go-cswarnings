package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var warningRegexp = regexp.MustCompile(`^(?P<path>.*): warning (?P<code>\w*): (?P<description>.*) \[(?P<project>.*)]$`)

	warninigs := make(map[string]Warning)
	for scanner.Scan() {
		text := scanner.Text()
		if warningRegexp.MatchString(text) {
			match := warningRegexp.FindStringSubmatch(text)
			// TODO: Find capture group indexes by names
			warning := Warning{
				path:        match[1],
				code:        match[2],
				description: match[3],
				project:     match[4],
			}
			warninigs[text] = warning
			fmt.Printf("%+v\n", warning)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Panicln("Scanner error", err)
	}
}

type Warning struct {
	path        string
	code        string
	description string
	project     string
}
