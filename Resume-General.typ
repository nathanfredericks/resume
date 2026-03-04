#let resume = yaml("resume-general.yaml")

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

#let contact(name, email1, email2, phone) = {
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
      #phone
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

#let position(title, date, company, location) = {
  entry(title, date)
  subentry(company + ", " + location)
}

#contact(
  resume.contact.name,
  resume.contact.emails.at(0),
  resume.contact.emails.at(1),
  resume.contact.phone
)

#section("Certifications")
#for item in resume.certifications [
  - #item
]

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

#section("Volunteer Activities")
#for (i, activity) in resume.volunteer_activities.enumerate() {
  position(activity.role, activity.dates, activity.organization, activity.location)

  for item in activity.description [
    - #item
  ]

  if i < resume.volunteer_activities.len() - 1 {
    v(0.5em)
  }
}

#section("Education")
#for edu in resume.education {
  entry(edu.degree, edu.dates)
  subentry(edu.institution + ", " + edu.location)

  if "accolades" in edu and edu.accolades.len() > 0 {
    list(
      [#underline[*Accolades:*] #edu.accolades.join(", ")]
    )
  }
}
