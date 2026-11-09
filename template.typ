// Shared layout for every resume in this repo, and for letters that should
// match it. Each resume is a YAML file with `contact` and an ordered list of
// `sections`; see resume.yaml.

#let link-color = rgb("#2A61BB")

#let header(contact) = {
  let details = contact.emails.map(email => link("mailto:" + email)[#email])
  details.push([#contact.phone])
  let profiles = ()
  for (key, label) in (("website", "Website"), ("github", "GitHub"), ("linkedin", "LinkedIn")) {
    if key in contact { profiles.push(link(contact.at(key))[#label]) }
  }
  if profiles.len() > 0 { details.push(profiles.join([ | ])) }
  grid(
    columns: (1fr, auto),
    align: (left + horizon, right + horizon),
    column-gutter: 2em,
    text(size: 36pt, weight: "bold")[#contact.name],
    details.join(linebreak()),
  )
}

#let section-title(title, gap: 1.3em) = block(sticky: true, above: gap, below: 0.7em, {
  text(size: 13pt, weight: "bold")[#upper(title)]
  v(-0.7em)
  line(length: 100%, stroke: 0.5pt)
})

#let two-sides(left, right) = grid(
  columns: (1fr, auto),
  column-gutter: 1em,
  left, right,
)

#let render-entry(e) = {
  // Compact entries: one heading row and one plain sentence, no bullets.
  if e.at("compact", default: false) {
    two-sides([*#e.title*, #e.org], e.at("dates", default: ""))
    if "text" in e {
      v(-0.5em)
      [#e.text]
    }
    return
  }
  let roles = if "roles" in e { e.roles } else { ((title: e.title, dates: e.at("dates", default: "")),) }
  for (i, role) in roles.enumerate() {
    let right = role.at("dates", default: "")
    if i == 0 and "link" in e { right = link(e.link.url)[#e.link.text] }
    two-sides([*#role.title*], right)
    if i < roles.len() - 1 { v(-0.5em) }
  }
  if "org" in e {
    v(-0.5em)
    let org = e.org
    if "location" in e { org = org + ", " + e.location }
    [#org]
  }
  for key in ("stack", "note") {
    if key in e {
      v(-0.5em)
      text(style: "italic")[#e.at(key)]
    }
  }
  if "bullets" in e and e.bullets.len() > 0 {
    v(-0.4em)
    list(..e.bullets.map(b => [#b]))
  }
}

#let base(size: 10.5pt, margin: (x: 1.9cm, y: 1.7cm), body) = {
  set page(paper: "us-letter", margin: margin)
  set text(font: "Helvetica", size: size)
  set smartquote(enabled: false) // keep straight ASCII quotes in the PDF
  set par(justify: false, leading: 0.6em, spacing: 1em)
  set list(marker: text(size: 1.2em)[#sym.bullet], indent: 0.6em, body-indent: 0.5em, spacing: 0.55em)
  show link: it => text(fill: link-color)[#it]
  body
}

#let render(data, size: 10.5pt, margin: (x: 1.9cm, y: 1.7cm), entry-gap: 0.55em, section-gap: 1.3em) = base(size: size, margin: margin, {
  header(data.contact)
  for s in data.sections {
    section-title(s.title, gap: section-gap)
    if "lines" in s {
      for l in s.lines {
        if "label" in l [*#l.label:* #l.text \ ] else [#l.text \ ]
      }
    }
    if "entries" in s {
      for (i, e) in s.entries.enumerate() {
        block(breakable: false, render-entry(e))
        if i < s.entries.len() - 1 { v(entry-gap) }
      }
    }
  }
})
