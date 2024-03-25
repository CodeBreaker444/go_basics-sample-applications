# Go basics

```
$ go mod init example/hello
go: creating new go.mod: module example/hello
```
-----
```
$ go mod tidy
go: finding module for package rsc.io/quote
go: found rsc.io/quote in rsc.io/quote v1.5.


$ go run .
```
------
```
❯ protoc -I=. --go_out=. book.proto

```