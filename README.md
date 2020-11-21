# gocommentator

Checks that exported names have [Doc Comments](https://github.com/golang/go/wiki/CodeReviewComments#doc-comments)
and that they have the right format.

Shamelessly copy & pasted from [golang/lint](https://github.com/golang/lint) with
minor modifications so it would be compatible with `go vet` and `golangci-lint`.

## Install

```bash
go get github.com/ealves-pt/cmd/gocommentator
```

## Usage

```bash
gocommentator
```

```bash
gocommentator ./...
```

```bash
gocommentator [path] [path] [path] [etc]
```

Add `-t` to include tests.

```bash
gocommentator -t [path]
```

Note: Paths are only inspected recursively if the Go /... recursive path suffix
is appended to the path.
