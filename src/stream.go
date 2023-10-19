package ltag

import (
	"strings"

	"github.com/gosimple/slug"
)

type StreamBuffer struct {
	Tags []string
}

func (s *StreamBuffer) Transform(line string) (string, bool) {
	if strings.HasPrefix(line, "#") {
		return s.transformTagLine(line)
	} else if line == "" {
		return s.transformEmptyLine(line)
	} else {
		return s.transformNormalLine(line)
	}
}

func (s *StreamBuffer) transformTagLine(line string) (string, bool) {
	tag := strings.Trim(line, "# \n")
	if len(s.Tags) > 1 {
		s.Tags[len(s.Tags)-1] = tag
	} else {
		s.Tags = append(s.Tags, tag)
	}
	return "", false
}

func (s *StreamBuffer) transformEmptyLine(line string) (string, bool) {
	s.Tags = []string{}
	return "", false
}

func (s *StreamBuffer) transformNormalLine(line string) (string, bool) {
	if len(s.Tags) == 0 {
		return line, true
	}

	if len(s.Tags) == 1 {
		tag := s.Tags[len(s.Tags)-1]
		return line + tagify(tag), true
	}

	topTag := s.Tags[len(s.Tags)-2]
	subTag := s.Tags[len(s.Tags)-1]
	tagsStr := tagify(topTag) + tagify(subTag)
	return line + tagsStr, true
}

func tagify(tag string) string {
	return " " + "#" + slug.Make(tag)
}
