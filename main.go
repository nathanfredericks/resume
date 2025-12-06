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
		Codeberg string   `yaml:"codeberg"`
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
		KnownSoftware struct {
			Software   []string `yaml:"software"`
			Frameworks []string `yaml:"frameworks"`
		} `yaml:"known_software"`
		Technical struct {
			ProficientWith []string `yaml:"proficient_with"`
			FamiliarWith   []string `yaml:"familiar_with"`
		} `yaml:"technical"`
	} `yaml:"skills"`
	RelevantProjects []struct {
		Title string `yaml:"title"`
		Dates string `yaml:"dates"`
		Links struct {
			Codeberg string `yaml:"codeberg"`
			Website  string `yaml:"website"`
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
	w("Codeberg (%s)\n", r.Contact.Codeberg)
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
	w("KNOWN SOFTWARE\n")
	if len(r.Skills.KnownSoftware.Software) > 0 {
		w("Software:\n")
		for _, item := range r.Skills.KnownSoftware.Software {
			w("  - %s\n", item)
		}
		w("\n")
	}
	if len(r.Skills.KnownSoftware.Frameworks) > 0 {
		w("Frameworks:\n")
		for _, item := range r.Skills.KnownSoftware.Frameworks {
			w("  - %s\n", item)
		}
		w("\n")
	}

	w("TECHNICAL\n")
	if len(r.Skills.Technical.ProficientWith) > 0 {
		w("Proficient with:\n")
		for _, item := range r.Skills.Technical.ProficientWith {
			w("  - %s\n", item)
		}
		w("\n")
	}
	if len(r.Skills.Technical.FamiliarWith) > 0 {
		w("Familiar with:\n")
		for _, item := range r.Skills.Technical.FamiliarWith {
			w("  - %s\n", item)
		}
		w("\n")
	}
	w("\n")

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
		if proj.Links.Codeberg != "" {
			links = append(links, fmt.Sprintf("View on Codeberg (%s)", proj.Links.Codeberg))
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
