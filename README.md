# ao-concepts logging module

![CI](https://github.com/ao-concepts/logging/workflows/CI/badge.svg)
[![codecov](https://codecov.io/gh/ao-concepts/logging/branch/master/graph/badge.svg?token=AQVUZTRGQS)](https://codecov.io/gh/ao-concepts/logging)

This module provides a logger that fits for use within the ao-concepts ecosystem.

## Information

The ao-concepts ecosystem is still under active development and therefore the API of this module may have breaking changes until there is a first stable release.

If you are interested in contributing to this project, feel free to open a issue to discus a new feature, enhancement or improvement. If you found a bug or security vulnerability in this package, please start a issue, or open a PR against `master`.

## Installation

```
go get -u github.com/ao-concepts/logging
```

## Usage

The `DefaultLogger` is a wrapper around [github.com/rs/zerolog](https://github.com/rs/zerolog) 
It integrates into [gorm.io/gorm](https://gorm.io/gorm) by using the `CreateGormLogger` function.

You can specify a `io.Writer` that is used by the logger. By default the logger uses `os.Stdout`.

```go
log := logging.New(logging.Warn, nil)

log.ErrError(err)
log.Error("string", arg1, arg2)
```

## Used packages 

This project uses some really great packages. Please make sure to check them out!

| Package                                                            | Usage              |
| ------------------------------------------------------------------ | ------------------ |
| [github.com/rs/zerolog](https://github.com/rs/zerolog)             | The default logger |
| [github.com/stretchr/testify](https://github.com/stretchr/testify) | Testing            |
| [gorm.io/gorm](gorm.io/gorm)                                       | Database access    |
