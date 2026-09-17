# dsa-go

From-scratch Go implementations. Notes live in [DSA](https://github.com/Aryagorjipour/DSA).

No wrappers around the standard library until the version in this repo works and has tests.

## Layout

```
ds/         data structures
algo/       named algorithms
paradigm/   worked examples of a design method (optional)
pattern/    two pointers, sliding window, prefix sums, …
```

One package per topic. Name matches the Obsidian note.

```
ds/array/
    array.go
    array_test.go
```

Prefer `[]T` plus integer indexes over pointer graphs. Same layout ports to Rust later.

## Rule per topic

1. Implement the structure or algorithm.
2. Table tests for the operations.
3. One small driver or extra test that would be ugly without it.
4. Link the note in DSA.

## Run

```
go test ./...
```

## Start

`ds/array` is the first package. Grow-by-2 lives in `ds/dynarray` next.
