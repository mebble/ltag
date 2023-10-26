package ltag

import (
	"strings"

	"github.com/gosimple/slug"
)

type TaggingBuf struct {
	iPattern string
	oPattern string
	headings []string
}

func NewTaggingBuf(inputPattern, outputPattern string) *TaggingBuf {
	return &TaggingBuf{
		iPattern: inputPattern,
		oPattern: outputPattern,
	}
}

func (s *TaggingBuf) Transform(line string) (string, bool) {
	if strings.HasPrefix(line, s.iPattern) {
		return s.transformHeadingLine(line)
	} else if line == "" {
		return s.transformEmptyLine(line)
	} else {
		return s.transformNormalLine(line)
	}
}

func (s *TaggingBuf) transformHeadingLine(line string) (string, bool) {
	heading := strings.Trim(line, s.iPattern)
	if len(s.headings) > 1 {
		s.headings[len(s.headings)-1] = heading
	} else {
		s.headings = append(s.headings, heading)
	}
	return "", false
}

func (s *TaggingBuf) transformEmptyLine(line string) (string, bool) {
	s.headings = []string{}
	return "", false
}

func (s *TaggingBuf) transformNormalLine(line string) (string, bool) {
	var inlineTagsStr string
	parts := strings.Split(line, s.iPattern)
	if len(parts) > 1 {
		line = strings.Trim(parts[0], " ")
		for _, p := range parts[1:] {
			inlineTagsStr = inlineTagsStr + tagify(p, s.oPattern)
		}
	}

	var tagsStr string
	for _, heading := range s.headings {
		parts := strings.Split(heading, s.iPattern)
		for _, p := range parts {
			tagsStr = tagsStr + tagify(p, s.oPattern)
		}
	}

	return line + tagsStr + inlineTagsStr, true
}

func tagify(s, pattern string) string {
	return " " + pattern + slug.Make(s)
}
