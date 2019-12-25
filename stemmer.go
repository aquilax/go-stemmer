// Package stemmer provides function to stem bulgarian words using the BULSTEM
// set of rules.
package stemmer

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strconv"
)

// Rules holds the stemming rules
type Rules map[string]string

// LoadRules loads stemming rules from a file
func LoadRules(fileName string, stemBoundary int) (Rules, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return LoadRulesStream(file, stemBoundary)
}

// LoadRulesStream loads stemming rules from a stream
func LoadRulesStream(reader io.Reader, stemBoundary int) (Rules, error) {
	scanner := bufio.NewScanner(reader)
	re := regexp.MustCompile(`([^\s]+)\s==>\s([^\s]+)\s([\d]+)`)
	var matches []string
	var boundary int
	var err error
	rules := make(Rules)
	for scanner.Scan() {
		matches = re.FindStringSubmatch(scanner.Text())
		if len(matches) == 4 {
			boundary, err = strconv.Atoi(matches[3])
			if err != nil {
				return nil, err
			}
			if boundary > stemBoundary {
				rules[matches[1]] = matches[2]
			}
		}
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return rules, nil
}

// Stem returns stemmed word based on the stemming rules
func Stem(word string, rules Rules) string {
	runes := []rune(word)
	var found bool
	var v string
	var suffix string
	for i := 0; i < len(runes); i++ {
		suffix = string(runes[i:])
		v, found = rules[suffix]
		if found {
			return string(runes[:i]) + v
		}
	}
	return word
}
