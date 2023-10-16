package ltag

import "strings"

type Section struct {
    Tag   string
    Lines []string
}

type Output struct {
    Sections []Section
}

func (o *Output) ParseLine(line string) {
    if strings.HasPrefix(line, "#") {
        tag := strings.Trim(line, "# \n")
        s := Section{Tag: tag, Lines: []string{}}
        o.Sections = append(o.Sections, s)
    } else if line == "" {
        // ignore
    } else {
        lastsection := &o.Sections[len(o.Sections) - 1]
        lastsection.Lines = append(lastsection.Lines, line)
    }
}
