/* Application: Global Regular Expression Print */

package automata

import (
	"bufio"
	"io"
	"log"
	"strings"
)

const wrapper = ".*"

// Print lines in which specified argument as substring.
func Grep(r io.Reader, arg string) {
	var sb strings.Builder
	sb.WriteString(wrapper)
	sb.WriteString(arg)
	sb.WriteString(wrapper)

	regex := NewRegex(sb.String())
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if regex.Recognizes(line) {
			log.Printf("line %d: %s", i, line)
		}

		i++
	}
}
