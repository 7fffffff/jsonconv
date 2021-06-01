# jsonconv

In the go standard library, encoding/json contains utility functions
for encoding and decoding JSON literal strings, but they are
unexported. jsonconv provides these functions for use with other
software.

The implementations are current as of go 1.16.

## Example

```go
// decoding
quoted := []byte(`"foo \"bar\""`)
str, ok := jsonconv.Unquote(quoted)
// str should be `foo "bar"`
// ...

// encoding
unquoted := `a \ b`
jsonBytes := jsonconv.AppendQuote(nil, unquoted)
// jsonBytes should be []byte(`"a \\ b"`)
// ...
```

## Why Not strconv?

strconv provides Quote and Unquote functions, but they are intended for
handling go literal strings, which are not the same as JSON literal
strings.

## Documentation

API documentation is available at:
https://pkg.go.dev/github.com/7fffffff/jsonconv

## License

BSD-3-Clause, same as the go standard library.