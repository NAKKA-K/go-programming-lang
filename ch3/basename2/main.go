package basename

import (
	"strings"
)

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // If '/' not exist return -1.
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
