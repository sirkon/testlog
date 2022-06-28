# testlog
Logging of [sirkon/errors](https://github.com/sirkon/errors) in tests, including context.

## Installation

```shell
go get github.com/sirkon/errors
```

## Usage

Just use helper functions:

```go
testlog.Log(errors.New("log entry"))
testlog.Error(errors.New("some serious error"))
```

## PS

Tests are meant to fail in this module.


