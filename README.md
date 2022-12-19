# hello

## Motivation

Sort out desirable project structure (file/dir layout.)

## Building

No environment variables, using the Debian (backports) golang install which provides

```text
hbarta@olive:~/Programming/go/src/hello$ go version
go version go1.19.3 linux/amd64
hbarta@olive:~/Programming/go/src/hello$
```

Build commands

```text
go build cmd/hello.go
go build ./cmd/hello.go
```
