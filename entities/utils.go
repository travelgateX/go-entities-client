package entities

import (
	"strconv"
	"strings"
)

// sliceToQuotedStringFormat converts string slice to formatted args string
// "[uno dos tres]" -> "["uno", "dos", "tres"]"
func sliceToQuotedStringFormat(s []string) string {
	var gr []string
	for _, g := range s {
		gr = append(gr, strconv.Quote(g))
	}
	return "[" + strings.Join(gr, ", ") + "]"
}
