# Go (golang)

> Go is an open source programming language that makes it easy to build simple,
> reliable, and efficient software. - golang.org

Go is a neat little language designed to be expressive, efficient and effective
in writing robust programs at scale. It is similar to C while maintaining a
simple syntax for readability. Its concurrency model based on CSP takes
advantage of modern computers and makes go a great general purpose language.

## Overview

- [Installing](#installing)
- [Development Environment](#development-environment)
- [Learning Go](./learning_go/)
- [Hugo](./hugo/)

## Installing and Running Go

```bash
brew install go

# create folders
mkdir -p $HOME/go/{bin,src,pkg}

# check version
go version

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

- [Go by example](https://gobyexample.com/maps)
- [The Go Programming Language](https://www.oreilly.com/library/view/the-go-programming/9780134190570/)
- [Learning Go: An Idiomatic Approach to Real-World Go Programming](https://www.oreilly.com/library/view/learning-go-2nd/9781098139285)
