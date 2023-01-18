package testutil

import (
	"bytes"
	"strings"
)

// Reindent deindents the provided strings and replace tabs by spaces,
// so yaml inlined into tests work properly when decoded.
func Reindent(in string) []byte {
	var buf bytes.Buffer
	var trim string
	var trimSet bool
	for _, line := range strings.Split(in, "\n") {
		if !trimSet {
			trimmed := strings.TrimLeft(line, "\t")
			if trimmed == "" {
				continue
			}
			if len(trimmed) != len(line) && trimmed[0] == ' ' {
				panic("Tabs and spaces mixed early on string:\n" + in)
			}
			trim = line[:len(line)-len(trimmed)]
			trimSet = true
		}
		trimmed := strings.TrimPrefix(line, trim)
		if len(trimmed) == len(line) && trim != "" && strings.Trim(line, "\t ") != "" {
			panic("Line not indented consistently:\n" + line)
		}
		trimmed = strings.ReplaceAll(trimmed, "\t", "    ")
		buf.WriteString(trimmed)
		buf.WriteByte('\n')
	}
	return buf.Bytes()
}
