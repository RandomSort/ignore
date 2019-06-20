# Ignore

[![pipeline status](https://gitlab.com/RandomSort/ignore/badges/master/pipeline.svg)](https://gitlab.com/RandomSort/ignore/commits/master)

Starting point for a Git extension that can help manage `.gitignore` files.

## Current feature level

- Able to detect if inside a Git repository
- When passed a path, adds it to `.gitignore` in root of workspace

## Developing

Edit `ignore.go` and run using `go run ignore.go`.

## Testing

`go test`

## Installing

`go install github.com/randomsort/ignore`

If you want to use this as a git extension rename the executable to `git-ignore` or `git-ignore.exe` on Windows, and put it on your path.

## Pipeline

Pipeline can be found on:
[https://gitlab.com/RandomSort/ignore/pipelines](Gitlab)
