# Go (golang)

> Go is an open source programming language that makes it easy to build simple,
> reliable, and efficient software. - golang.org

Go is a neat little language designed to be expressive, efficient and effective
in writing robust programs at scale. It is similar to C while maintaining a
simple syntax for readability. Its concurrency model based on CSP takes
advantage of modern computers and makes go a great general purpose language.


## Installing

```bash
brew install golang

# in .zshrc
export GOPATH=$HOME/go
export GOROOT="$(brew --prefix golang)/libexec"
export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin"

# create folders
mkdir -p $HOME/go/{bin,src,pkg}
```

- [Hello World](tmp/001)
- [Interfaces](tmp/002)
- [Goroutines](tmp/003)
- [Comparing Type memory](tmp/004)
- [Sending JSON](tmp/005)


## Resources

- [Go in 5 minutes](https://goin5minutes.com/screencasts/)
- [Go by example](https://gobyexample.com/maps)
- [Go Gitbook](https://astaxie.gitbooks.io/build-web-application-with-golang/en/02.2.html)
- [The Go Programming Language](https://www.oreilly.com/library/view/the-go-programming/9780134190570/)
