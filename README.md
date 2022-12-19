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

## Testing

```text
go test test/hello_test.go
```

## Layout

```text
hbarta@olive:~/Programming/go/src/hello$ tree
.
├── cmd
│   └── hello.go
├── go.mod
├── hello
├── pkg
│   └── lib
│       └── lib.go
├── README.md
└── test
    └── hello_test.go

4 directories, 6 files
hbarta@olive:~/Programming/go/src/hello$ 
```

## Status

Build working.  
Testing working.

## Errata

Project format as specified at <https://jchiang1225.medium.com/golang-project-structure-beginner-guide-2022-b18285ddc1be> and <https://github.com/golang-standards/project-layout>
