# random-merge-function

Small Go project that implements a three-collection merge function without using any Go sort APIs.

## Problem

Implement:

```go
func Merge(collection1, collection2, collection3 []int) []int
```

`Merge` returns a new slice sorted in ascending order.

Input ordering assumptions:

- `collection1` is already sorted from max to min.
- `collection2` is already sorted from min to max.
- `collection3` is already sorted from min to max.

The implementation does not use `sort`, `slices.Sort`, or any other sorting helper. It performs an O(n) three-way merge by reading `collection1` from right to left and the other collections from left to right.

## Requirements

- Go 1.26.2 or newer, matching the module version in `go.mod`.
- No third-party dependencies.

## Setup

```sh
go mod tidy
```

## Run Demo

```sh
go run ./cmd/merge-demo
```

Expected output:

```text
[0 1 2 3 4 5 6 7 8 9 10]
```

## Run Unit Tests

```sh
go test ./...
```

For race detection:

```sh
go test -race ./...
```
