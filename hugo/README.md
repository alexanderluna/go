# Hugo Framework

> Hugo is a framework to build fast and organized content sites.

## Overview

- [Installation](#installation)
- [Directory Structure](#directory-structure)
- [Themes](#themes)
- [Configuring Data](#configuring-data)
- [Syntax Highlighting](#syntax-highlighting)
- [Search](#search)

---

- [Resources](#resources)

## Installation

```sh
brew install hugo

hugo new site SITE_NAME

hugo server
```

## Directory Structure

After running `hugo new site ...` you are presented with a directory structure
and a config file called `hugo.toml`.

```sh
hugo/
├── archetypes      # templates for new content
│ └── default.md
├── assets          # global resources (img, css, js)
├── content         # posts, videos, etc.
├── data            # data files (json, xml)
├── layouts         # templates to render content
├── public          # published website
├── resources       # cached output from asset pipeline          
├── static          # will be copied to public
├── themes          # look and feel of the website
└── hugo.toml
```

Hugo uses layouts to generate consistent pages and give you a place to store
boilerplate code. At its minimum you need a single, list and default layout.

```sh

layouts/
├── _default
│ ├── list.html
│ └── single.html
└─── index.html
```

## Themes

Themes give your website its look and feel. Once you generate your own theme
and configure you theme in the `hugo.toml` file you can move your layouts to it.

```sh
hugo new theme THEME_NAME

echo "theme = 'THEME_NAME'" >> hugo.toml
```

Within the layouts you use `{{` to inject content and `{{-` to inject content
without whitespace. You can break down your layouts using partials and inject
them in your layout with the `{{- partial "something.html" . }}` function.
Notice the period after the html file which tells the partial function the
current context.

In `baseof.html` you will find a main block: `{{- block "main" . }}{{- end }}`.
The main block looks for a block named main in the layouts. This means you
have to define the main blocks in your layout files.

```go
// layouts/index.html

{{ define "main" }}
{{ .Content }}
{{ end }}
```

## Configuring Data

Inside the `hugo.toml` file you can define a `params` section with data which
you can access in your templates using `.Params` variable.

```toml
[params]
    author = "Alexander Luna"
```

```html
<meta name="author" content="{{ .Site.Params.author }}">
```

Your archetypes front matter is also accessible in the templates through the
`.Params` variable.

```md
---
title: "Hello World"
---
```

```handlebars
<h2>{{ .Params.title }}</h2>
```

Using the `data` directory you can save data in `JSON`, `YAML` or `TOML` format
to access it using the `.Data` variable.

```json
// data/text.json

{
    "post": "some blog post data"
}
```

```handlebars
<p>{{ .Site.Data.post }}</p>
```

You can also pull in data from an API using the hugo's `getJSON` function.
Every time you build your site, Hugo fetching from the API and renders it which
means that it will not be always up to date when you access the website.

```handlebars
  {{ $url := 'some_api_url' }}
  {{ $data := getJSON $url }}
```

> Hugo will generate an RSS 2.0 feed by default at /index.xml
> You can generate a json output for your pages using
> [Custom output formats](https://gohugo.io/templates/output-formats/)

## Syntax Highlighting

If you want to add syntax highlighting for code snippets, you have to add the
`Pygments-style` class to your `hugo.toml` config file.

```toml
# hugo.toml
pygmentsUseClasses = true
```

Now you can generate the actual CSS styles using Chroma styles.

```sh
hugo gen chromastyles --style=github > syntax.css
```

## Search

While it is possible to use a 3rd party search service which indexes your
website to add search functionality to your side, it is much faster and easier
to generate your own search index which can then be filtered on the client side
using Javascript.

## Resources

- [Build Websites with Hugo](https://pragprog.com/titles/bhhugo/build-websites-with-hugo/)
- [Hugo Documentation: consult frequently](https://gohugo.io/documentation/)
- [Pygments - syntax highlighter](https://pygments.org)
- [Chroma - general purpose syntax highlighter in pure Go](https://github.com/alecthomas/chroma)
- [Lunr - Search made simple](https://lunrjs.com/docs/index.html)
