package ltag

import (
	"strings"

	"github.com/gosimple/slug"
)

type TaggingBuf struct {
	iPattern string
	oPattern string
	Tags     []string
}

func NewTaggingBuf(inputPattern, outputPattern string) *TaggingBuf {
	return &TaggingBuf{
		iPattern: inputPattern,
		oPattern: outputPattern,
	}
}

func (s *TaggingBuf) Transform(line string) (string, bool) {
	if strings.HasPrefix(line, s.iPattern) {
		return s.transformTagLine(line)
	} else if line == "" {
		return s.transformEmptyLine(line)
	} else {
		return s.transformNormalLine(line)
	}
}

func (s *TaggingBuf) transformTagLine(line string) (string, bool) {
	tag := strings.Trim(line, s.iPattern)
	if len(s.Tags) > 1 {
		s.Tags[len(s.Tags)-1] = tag
	} else {
		s.Tags = append(s.Tags, tag)
	}
	return "", false
}

func (s *TaggingBuf) transformEmptyLine(line string) (string, bool) {
	s.Tags = []string{}
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
	for _, tag := range s.Tags {
		tagsStr = tagsStr + tagify(tag, s.oPattern)
	}

	return line + tagsStr + inlineTagsStr, true
}

func tagify(tag, pattern string) string {
	return " " + pattern + slug.Make(tag)
}
