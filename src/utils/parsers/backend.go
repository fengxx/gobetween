/**
 * backend.go - backend parser utils
 *
 * @author Ievgen Ponomarenko <kikomdev@gmail.com>
 * @author Yaroslav Pogrebnyak <yyyaroslav@gmail.com>
 */

package parsers

import (
	"../../core"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

const (
	DEFAULT_BACKEND_PATTERN = `^(?P<host>\S+):(?P<port>\d+)(\sweight=(?P<weight>\d+))?(\spriority=(?P<priority>\d+))?$`
)

/**
 * Do parding of backend line with default pattern
 */
func ParseBackendDefault(line string) (*core.Backend, error) {
	return ParseBackend(line, DEFAULT_BACKEND_PATTERN)
}

/**
 * Do parsing of backend line
 */
func ParseBackend(line string, pattern string) (*core.Backend, error) {

	//trim string
	line = strings.TrimSpace(line);

	// parse string by regexp
	var reg = regexp.MustCompile(pattern)
	match := reg.FindStringSubmatch(line)

	if len(match) == 0 {
		return nil, errors.New("Cant parse " + line)
	}

	result := make(map[string]string)

	// get named capturing groups
	for i, name := range reg.SubexpNames() {
		if name != "" {
			result[name] = match[i]
		}
	}

	weight, err := strconv.Atoi(result["weight"])
	if err != nil {
		weight = 1
	}

	priority, err := strconv.Atoi(result["priority"])
	if err != nil {
		priority = 1
	}

	backend := core.Backend{
		Target: core.Target{
			Host: result["host"],
			Port: result["port"],
		},
		Weight:   weight,
		Priority: priority,
		Live:     true,
	}

	return &backend, nil
}
