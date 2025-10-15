# Go (golang)

> Go is an open source programming language that makes it easy to build simple,
> reliable, and efficient software. - golang.org

Go is a neat little language designed to be expressive, efficient and effective
in writing robust programs at scale. It is similar to C while maintaining a
simple syntax for readability. Its concurrency model based on CSP takes
advantage of modern computers and makes go a great general purpose language.

## Overview

- [Installing Go](#installing-and-running-go)
- [Setup Development Environment](#development-environment)
- [Learning Go](./learning_go/)
- [Hugo](./hugo/): A blog with a custom theme, archetypes and search
- [CLI Apps](#overview)
  - [Bookworms](./cli_bookworms/): Bookshelf comparison from JSON
  - [Logger](./cli_logger/): logger API with DST and Functional options pattern

## Installing and Running Go

```bash
brew install go

# create folders
mkdir -p $HOME/go/{bin,src,pkg}

# check version
go version

# initialize a module
go mod init MODULE_NAME

# execute a directory (.) or file (*.go)
go run . 

# run tests
go test
```

## Development Environment

In order to write Go programs you will need three extensions:

- Go development tools
- Delve debugger
- gopls (language server)

## Resources

- ★★★★☆ [Go by example](https://gobyexample.com/)
- ★★★☆☆ [Learning Go: An Idiomatic Approach to Real-World Go Programming](https://www.oreilly.com/library/view/learning-go-2nd/9781098139285)
- ★★★★★ [Learn Go with Pocket-Sized Projects](https://www.manning.com/books/learn-go-with-pocket-sized-projects)
- [The Go Programming Language](https://www.oreilly.com/library/view/the-go-programming/9780134190570/)
