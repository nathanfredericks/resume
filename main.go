// Command main writes a plain-text copy of a resume YAML file.
//
//	go run main.go                               # resume.yaml -> resume.txt
//	go run main.go resume-physio.yaml physio.txt # any resume -> any file
package main

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Resume struct {
	Contact struct {
		Name     string   `yaml:"name"`
		Emails   []string `yaml:"emails"`
		Phone    string   `yaml:"phone"`
		Website  string   `yaml:"website"`
		GitHub   string   `yaml:"github"`
		Linkedin string   `yaml:"linkedin"`
	} `yaml:"contact"`
	Sections []struct {
		Title   string `yaml:"title"`
		Entries []struct {
			Title string `yaml:"title"`
			Dates string `yaml:"dates"`
			Roles []struct {
				Title string `yaml:"title"`
				Dates string `yaml:"dates"`
			} `yaml:"roles"`
			Stack string `yaml:"stack"`
			Link  struct {
				Text string `yaml:"text"`
				URL  string `yaml:"url"`
			} `yaml:"link"`
			Org      string   `yaml:"org"`
			Location string   `yaml:"location"`
			Note     string   `yaml:"note"`
			Text     string   `yaml:"text"` // body of a compact entry
			Bullets  []string `yaml:"bullets"`
		} `yaml:"entries"`
		Lines []struct {
			Label string `yaml:"label"`
			Text  string `yaml:"text"`
		} `yaml:"lines"`
	} `yaml:"sections"`
}

func main() {
	in, out := "resume.yaml", "resume.txt"
	if len(os.Args) > 1 {
		in = os.Args[1]
	}
	if len(os.Args) > 2 {
		out = os.Args[2]
	}

	data, err := os.ReadFile(in)
	if err != nil {
		panic(err)
	}
	var r Resume
	if err := yaml.Unmarshal(data, &r); err != nil {
		panic(err)
	}

	var b strings.Builder
	w := func(format string, a ...any) { fmt.Fprintf(&b, format, a...) }

	// Same layout as the PDF: one item per line, no separator characters.
	c := r.Contact
	w("%s\n", c.Name)
	for _, line := range append(append([]string{}, c.Emails...), c.Phone, c.Website, c.GitHub, c.Linkedin) {
		if line != "" {
			w("%s\n", line)
		}
	}

	withDates := func(title, dates string) string {
		if dates == "" {
			return title
		}
		return title + ", " + dates
	}

	for _, s := range r.Sections {
		w("\n%s\n", strings.ToUpper(s.Title))
		for _, l := range s.Lines {
			if l.Label != "" {
				w("%s: %s\n", l.Label, l.Text)
			} else {
				w("%s\n", l.Text)
			}
		}
		for _, e := range s.Entries {
			w("\n")
			if len(e.Roles) > 0 {
				for _, role := range e.Roles {
					w("%s\n", withDates(role.Title, role.Dates))
				}
			} else {
				w("%s\n", withDates(e.Title, e.Dates))
			}
			org := e.Org
			if e.Location != "" {
				org = strings.TrimPrefix(org+", "+e.Location, ", ")
			}
			for _, line := range []string{org, e.Link.URL, e.Stack, e.Note, e.Text} {
				if line != "" {
					w("%s\n", line)
				}
			}
			for _, bullet := range e.Bullets {
				w("- %s\n", bullet)
			}
		}
	}

	if err := os.WriteFile(out, []byte(b.String()), 0o644); err != nil {
		panic(err)
	}
}
