package ltag

import "strings"

type Section struct {
    tag   string
    lines []string
}

type Output struct {
    Sections []Section
}

func (o *Output) ParseLine(line string) {
    if strings.HasPrefix(line, "#") {
        tag := strings.Trim(line, "# \n")
        s := Section{tag: tag, lines: []string{}}
        o.Sections = append(o.Sections, s)
    } else {
        lastsection := &o.Sections[len(o.Sections) - 1]
        lastsection.lines = append(lastsection.lines, line)
    }
}
