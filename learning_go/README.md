# Learning go

Go compiles to a single native binary without requiring any additional software
to run it which makes it easy to distribute your programs. In order to, compile,
build, format, manage and test your go programs you have the following
development tools at your disposal:

- `go build`: compiler
- `go fmt`: code formatter
- `go mod`: dependency manager
- `go test`: test runner
- `go vet`: mistakes scanner

## Overview

- [The Basics](#the-basics)
- [Hello World](./hello_world/hello.go)
- [Predeclared Types](./predeclared_types/predeclared_types.go)
- [Composite Types](./composite_types/composite_types.go)
- [Resources](#resources)

## The Basics

In go, a project is called a module and exists within its own directory. Inside
this directory you need a `go.mod` file which specifies the minimum version of
go required, its dependencies and the name of the module.

```sh
mkdir hello_world
cd hello_world

# marks directory as a module and creates go.mod
go mod init hello_world

# build and generate executable with the mod name
go build hello_world

./hello_world
やっと来たよ。
```

Go defines a standard way of formatting code making the `go fmt` tool a part of
development process: **write** -> **save** -> **format** -> **build**

```sh
# format the current directory and all its subdirectories
go fmt ./...
```

In order to have repeatable and automatable builds in Go, developers have
adopted `make` as their solution. `make` specifies the necessary steps to build
a program.

```make
.DEFAULT_GOAL := build

.PHONY:fmt vet build

fmt: 
    go fmt ./...
vet: fmt
    go vet ./...
build: vet
    go build
```

## Resources

- [The Go Programming Language](https://www.oreilly.com/library/view/the-go-programming/9780134190570/)
- [Learning Go: An Idiomatic Approach to Real-World Go Programming](https://www.oreilly.com/library/view/learning-go-2nd/9781098139285)
