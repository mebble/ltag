package ltag

import (
	"strings"
)

type Section struct {
	Tag   string
	Lines []string
	Subs  []Section
}

type Output struct {
	Sections []Section
}

var isSub = false

func (o *Output) ParseLine(line string) {
	if strings.HasPrefix(line, "#") {
		o.parseTagLine(line)
	} else if line == "" {
		o.parseEmptyLine(line)
	} else {
		o.parseNormalLine(line)
	}
}

func (o *Output) parseTagLine(line string) {
	tag := strings.Trim(line, "# \n")
	if isSub && len(o.Sections) > 0 {
		lastsection := &o.Sections[len(o.Sections)-1]
		s := Section{Tag: tag, Lines: []string{}, Subs: []Section{}}
		lastsection.Subs = append(lastsection.Subs, s)
	} else {
		s := Section{Tag: tag, Lines: []string{}, Subs: []Section{}}
		o.Sections = append(o.Sections, s)
	}
	isSub = true
}

func (o *Output) parseEmptyLine(line string) {
	isSub = false
}

func (o *Output) parseNormalLine(line string) {
	lastsection := &o.Sections[len(o.Sections)-1]
	if isSub && len(lastsection.Subs) > 0 {
		lastSub := &lastsection.Subs[len(lastsection.Subs)-1]
		lastSub.Lines = append(lastSub.Lines, line)
		return
	}
	lastsection.Lines = append(lastsection.Lines, line)
}
