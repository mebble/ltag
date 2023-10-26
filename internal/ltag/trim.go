package ltag

import "strings"

type TrimmingBuf struct {
	pattern string
}

func NewTrimmingBuf(pattern string) *TrimmingBuf {
	return &TrimmingBuf{
		pattern: pattern,
	}
}

func (s *TrimmingBuf) Transform(line string) (string, bool) {
	idx := strings.Index(line, s.pattern)
	if idx == -1 {
		return line, true
	}

	return line[:idx-1], true
}
