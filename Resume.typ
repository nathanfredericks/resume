#let resume = yaml("resume.yaml")

#set page(
  margin: (x: 1.91cm, y: 1.91cm),
)

#set text(
  font: "Helvetica",
  size: 11pt,
)

#set par(justify: false, leading: 0.65em)

#set list(marker: text(size: 16pt)[•], indent: 1em, spacing: 0.5em)

#show link: it => underline(text(fill: rgb("#2A61BB"))[#it])

#let contact(
  name,
  email1,
  email2,
  phone,
  website,
  github,
  linkedin
) = {
  grid(
    columns: (1fr, auto),
    align: (left + horizon, right + horizon),
    column-gutter: 2em,
    [
      #text(size: 36pt, weight: "bold")[#name]
    ],
    [
      #set align(right)
      #link("mailto:" + email1)[#email1] \
      #link("mailto:" + email2)[#email2] \
      #phone \
      #link(website)[Website] | #link(github)[GitHub] | #link(linkedin)[LinkedIn]
    ]
  )
}


#let section(title) = {
  v(0.5em)
  text(size: 14pt, weight: "bold")[#upper(title)]
  v(-0.5em)
}

#let entry(title, date) = {
  grid(
    columns: (1fr, auto),
    [*#title*],
    [#date]
  )
}

#let subentry(org) = {
  v(-0.5em)
  text[#org]
  v(-0.25em)
}

#let project(title, date, github: none, website: none) = {
  entry(title, date)
  v(-0.5em)
  if website != none and github != none {
    link(website)[Website] + [ | ] + link(github)[View on GitHub]
  } else if github != none {
    link(github)[View on GitHub]
  } else if website != none {
    link(website)[Website]
  }
}

#let position(title, date, company, location) = {
  entry(title, date)
  subentry(company + ", " + location)
}

#contact(
  resume.contact.name,
  resume.contact.emails.at(0),
  resume.contact.emails.at(1),
  resume.contact.phone,
  resume.contact.website,
  resume.contact.github,
  resume.contact.linkedin
)

#section("Education")
#for edu in resume.education {
  entry(edu.degree, edu.dates)
  subentry(edu.institution + ", " + edu.location)
  
  list(
    [#underline[*Relevant Coursework:*] #edu.relevant_coursework.join(", ")],
    [#underline[*Accolades:*] #edu.accolades.join(", ")]
  )
}

#section("Skills")

#grid(
  columns: (1fr, 1fr),
  column-gutter: 2em,
  [
    *LANGUAGES* \
    #for item in resume.skills.languages [
      - #item
    ]

    *FRAMEWORKS* \
    #for item in resume.skills.frameworks [
      - #item
    ]
  ],
  [
    *INFRASTRUCTURE* \
    #for item in resume.skills.infrastructure [
      - #item
    ]

    *TOOLS & DATABASES* \
    #for item in resume.skills.tools_and_databases [
      - #item
    ]
  ]
)

#section("Work Experience")



#for (i, job) in resume.work_experience.enumerate() {

  position(job.title, job.dates, job.company, job.location)

  

  for item in job.description [

    - #item

  ]

  

  if i < resume.work_experience.len() - 1 {

    v(0.5em)

  }

}



#section("Projects")



#for (i, proj) in resume.relevant_projects.enumerate() {

  let website = if "links" in proj and "website" in proj.links { proj.links.website } else { none }

  let github = if "links" in proj and "github" in proj.links { proj.links.github } else { none }

  

  project(

    proj.title,

    proj.dates,

    github: github,

    website: website

  )



  for item in proj.description [

    - #item

  ]

  

  if i < resume.relevant_projects.len() - 1 {

    v(0.5em)

  }

}