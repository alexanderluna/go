# Hugo Blog

> Hugo is a framework to build fast and organized content sites.

## Overview

- [Installation](#installation)
- [Directory Structure](#directory-structure)
- [Themes](#themes)
- [Configuring Data](#configuring-data)
- [Syntax Highlighting](#syntax-highlighting)
- [Search](#search)
- [Working with Assets](#working-with-assets)
- [Deployment](#deploying)
  - [Web Server deployment](#web-server-deployment)
  - [Cloud Storage deployment](#cloud-storage-deployment)
  - [Continuous Integration deployment](#continuous-integration-deployment)
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
using Javascript (for example with Lunr):

[view source](./themes/basic/assets/js/index.js)

## Working with Assets

Hugo comes with pipes to manage assets such as css, javascript and images for
production.

- `minify`: remove whitespace and comments
- `fingerprint`: add hash string of the resource's content
- `toCSS`: transpile Sass to CSS

> Saves your images in the `static` directory and render it using the `relURL`
> pipe.

Hugo provides several `shortcodes` for rendering images, youtube videos and
other types assets. It also possible to create your own `shortcodes` to adjust
to your needs like resizing images. For that you have to create a `shortcodes`
directory in your themes layout directory.

## Deploying

When it comes to deploying your Hugo site you have three main options:

- upload the `public` directory to a server
- push and build your code in the cloud using Continuous Integration
- dockerize your project and deploy on a docker host

> When you build your site on your local machine all Hugo does is to generate
> static files in the `public` directory. Hence, this is the easiest way to
> deploy your site.

### Web Server deployment

Hosting a hugo site on your own web server is the easiest approach. All you need
is the IP, user account and password. You build your site locally and deploy it
using `sftp`, `scp` or `rsync`.

```sh
# build your site
webpack && hugo --cleanDestinationDir

# upload your public folder to the web server
sftp -r username@hostname:public_html <<< 'put public/*'
scp -r public/* username@hostname:public_html
rsync -avz --delete public/ username@hostname.com:public_html

# or use an npm script
npm run deploy:webserver
```

### Cloud Storage deployment

If you want to deploy your site to a cloud storage provider like AWS S3, Azure
or Google Cloud Storage. Configure your storage for hosting.

```sh
# configure credentials
aws configure 

# create a bucket
aws s3api create-bucket --bucket BUCKET_NAME --acl public-read --region YOUR_REGION

# enable bucket to serve webpages
aws s3 website BUCKET_URL --index-document index.html --error-document error.html

# create a bucket policy which decides who can access your website
aws s3api put-bucket-policy --bucket BUCKET_NAME --policy file://bucketpolicy.json
```

Now you can add your AWS S3 bucket and region configuration to your `hugo.toml`
file to deploy your side using hugo.

```sh
hugo --cleanDestinationDir
hugo deploy
```

### Continuous Integration deployment

For the CI deployment you need to create a git repository and add a `.gitignore`

```sh
git init

echo 'node_modules' >> .gitignore
echo 'public' >> .gitignore

git add .
git commit -m 'Deploying site using CI'
git remote add origin GIT_REPO_URL
git push -u origin master
```

Now you can log into your continuous deployment provider
(Netlify, Cloudflare, Github pages). There you have to connect your git account
and authorize access to your repository. The build command is the npm script.
The publish directory is the `public` directory which hugo generates after the
build process finishes.

```sh
npm run deploy:cloud
```

## Resources

- [Build Websites with Hugo](https://pragprog.com/titles/bhhugo/build-websites-with-hugo/)
- [Hugo Documentation: consult frequently](https://gohugo.io/documentation/)
- [Pygments - syntax highlighter](https://pygments.org)
- [Chroma - general purpose syntax highlighter in pure Go](https://github.com/alecthomas/chroma)
- [Lunr - Search made simple](https://lunrjs.com/docs/index.html)
