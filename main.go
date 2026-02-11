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
	Education []struct {
		Degree             string   `yaml:"degree"`
		Institution        string   `yaml:"institution"`
		Location           string   `yaml:"location"`
		Dates              string   `yaml:"dates"`
		RelevantCoursework []string `yaml:"relevant_coursework"`
		Accolades          []string `yaml:"accolades"`
	} `yaml:"education"`
	Skills struct {
		Languages         []string `yaml:"languages"`
		Frameworks        []string `yaml:"frameworks"`
		Infrastructure    []string `yaml:"infrastructure"`
		ToolsAndDatabases []string `yaml:"tools_and_databases"`
	} `yaml:"skills"`
	RelevantProjects []struct {
		Title string `yaml:"title"`
		Dates string `yaml:"dates"`
		Links struct {
			GitHub  string `yaml:"github"`
			Website string `yaml:"website"`
		} `yaml:"links"`
		Description []string `yaml:"description"`
	} `yaml:"relevant_projects"`
	WorkExperience []struct {
		Title       string   `yaml:"title"`
		Company     string   `yaml:"company"`
		Location    string   `yaml:"location"`
		Dates       string   `yaml:"dates"`
		Description []string `yaml:"description"`
	} `yaml:"work_experience"`
}

func main() {
	data, err := os.ReadFile("resume.yaml")
	if err != nil {
		panic(err)
	}

	var r Resume
	if err := yaml.Unmarshal(data, &r); err != nil {
		panic(err)
	}

	f, err := os.Create("resume.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := func(format string, a ...interface{}) {
		fmt.Fprintf(f, format, a...)
	}

	w("%s\n", r.Contact.Name)
	for _, email := range r.Contact.Emails {
		w("%s\n", email)
	}
	w("%s\n", r.Contact.Phone)
	w("Website (%s)\n", r.Contact.Website)
	w("GitHub (%s)\n", r.Contact.GitHub)
	w("LinkedIn (%s)\n", r.Contact.Linkedin)
	w("\n\n")

	w("EDUCATION\n")
	for _, edu := range r.Education {
		w("%s | %s\n", edu.Degree, edu.Dates)
		w("%s, %s\n", edu.Institution, edu.Location)
		if len(edu.RelevantCoursework) > 0 {
			w("- Relevant Coursework: %s\n", strings.Join(edu.RelevantCoursework, ", "))
		}
		if len(edu.Accolades) > 0 {
			w("- Accolades: %s\n", strings.Join(edu.Accolades, ", "))
		}
		w("\n\n")
	}

	w("SKILLS\n")
	if len(r.Skills.Languages) > 0 {
		w("Languages:\n")
		for _, item := range r.Skills.Languages {
			w("  - %s\n", item)
		}
		w("\n")
	}
	if len(r.Skills.Frameworks) > 0 {
		w("Frameworks:\n")
		for _, item := range r.Skills.Frameworks {
			w("  - %s\n", item)
		}
		w("\n")
	}
	if len(r.Skills.Infrastructure) > 0 {
		w("Infrastructure:\n")
		for _, item := range r.Skills.Infrastructure {
			w("  - %s\n", item)
		}
		w("\n")
	}
	if len(r.Skills.ToolsAndDatabases) > 0 {
		w("Tools & Databases:\n")
		for _, item := range r.Skills.ToolsAndDatabases {
			w("  - %s\n", item)
		}
		w("\n")
	}

	w("WORK EXPERIENCE\n")
	for _, work := range r.WorkExperience {
		w("%s | %s\n", work.Title, work.Dates)
		w("%s, %s\n", work.Company, work.Location)
		for _, desc := range work.Description {
			w("- %s\n", desc)
		}
		w("\n")
	}
	w("\n")

	w("PROJECTS\n")
	for _, proj := range r.RelevantProjects {
		w("%s | %s\n", proj.Title, proj.Dates)

		var links []string
		if proj.Links.Website != "" {
			links = append(links, fmt.Sprintf("Website (%s)", proj.Links.Website))
		}
		if proj.Links.GitHub != "" {
			links = append(links, fmt.Sprintf("View on GitHub (%s)", proj.Links.GitHub))
		}
		if len(links) > 0 {
			w("%s\n", strings.Join(links, " | "))
		}

		for _, desc := range proj.Description {
			w("- %s\n", desc)
		}
		w("\n")
	}
	w("\n")
}
