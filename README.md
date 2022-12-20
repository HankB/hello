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

## misc

```text
go fmt ./cmd
```

## Layout

```text
hbarta@olive:~/Programming/go/src/hello$ tree
.
├── cmd
│   ├── hello.go
│   └── hello_test.go
├── go.mod
├── internal
│   └── lib.go
└── README.md

2 directories, 5 files
hbarta@olive:~/Programming/go/src/hello$ 
```

## Status

Build working.  
Testing working.
First cut, reorganize directories.

## Resources

* <https://developer20.com/how-to-structure-go-code/>
*  <https://jchiang1225.medium.com/golang-project-structure-beginner-guide-2022-b18285ddc1be>
* <https://github.com/golang-standards/project-layout>
