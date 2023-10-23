package ltag

import "strings"

type FilterBuf struct {
	pattern string
}

func NewFilterBuf(pattern string) *FilterBuf {
	return &FilterBuf{
		pattern: pattern,
	}
}

func (s *FilterBuf) Transform(line string) (string, bool) {
	idx := strings.Index(line, s.pattern)
	if idx == -1 {
		return line, true
	}

	return line[:idx-1], true
}
