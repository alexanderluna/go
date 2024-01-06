# Hugo Framework

> Hugo is a framework to build fast and organized content sites.

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

## Resources

- [Build Websites with Hugo](https://pragprog.com/titles/bhhugo/build-websites-with-hugo/)