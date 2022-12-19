# testlog
Logging of [sirkon/errors](https://github.com/sirkon/errors) in tests, including context.

## Installation

```shell
go get github.com/sirkon/errors
```

## Usage

Just use helper functions:

```go
testlog.Log(t, errors.New("log entry"))
testlog.Error(t, errors.New("some serious error"))
```

or

```go
tl := testlog.New(t)
tl.Log(errors.New("log entry"))
tl.Error(errors.New("som serious error"))
```

## PS

Tests are meant to fail in this module.


