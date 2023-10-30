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
	currentLevel := getLevel(s.iPattern, line)
	if len(s.headings) == 0 {
		s.headings = append(s.headings, heading)
	} else if currentLevel <= len(s.headings) {
		// the index+1 denotes the level i.e. 0 is level 1, etc. This also means len is last level
		preservedHeadings := s.headings[:currentLevel-1]
		s.headings = append(preservedHeadings, heading)
	} else {
		emptyHeadings := make([]string, currentLevel - len(s.headings) - 1)
		s.headings = append(append(s.headings, emptyHeadings...), heading)
	}
	return "", false
}

func getLevel(pattern, s string) int {
	level := 0
	for _, char := range s {
		c := string(char)
		if c != pattern {
			return level
		}
		level++
	}
	return level
}

func (s *TaggingBuf) transformEmptyLine(line string) (string, bool) {
	return "", false
}

func (s *TaggingBuf) transformNormalLine(line string) (string, bool) {
	var inlineTagsStr string
	parts := strings.Split(line, s.iPattern)
	if len(parts) > 1 {
		line = strings.Trim(parts[0], " ")
		for _, p := range parts[1:] {
			inlineTagsStr = inlineTagsStr + tagify(s.oPattern, p)
		}
	}

	var tagsStr string
	for _, heading := range s.headings {
		parts := strings.Split(heading, s.iPattern)
		for _, p := range parts {
			tagsStr = tagsStr + tagify(s.oPattern, p)
		}
	}

	return line + tagsStr + inlineTagsStr, true
}

func tagify(pattern, s string) string {
	if s == "" {
		return ""
	}
	return " " + pattern + slug.Make(s)
}
