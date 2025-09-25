# go-errorsx

go-errorsx attempts to overcome some of the issues that make it hard to debug issues when the only error message provided is a simple string of the error message.

This package provides an error type that adds a stack trace and attribute key/value pairs wrapping to errors, whilst still being compatible with the Go `error` interface.

## Usage:

```
func main() {
    _, err := readFile("example.txt")
    if err != nil {
        // print error message and key/value attributes
        fmt.Printf("%s", error.Error())

        // print error stack trace
        fmt.Printf("%s", error.Stack())
    }

    // if err is not nil, print error message, key/value attributes and stack trace to Stderr pipe.
    // Then terminate the program with exit code 1.
    errorsx.ExitIfErr(err)
}

func readFile(path string) (*bytes.Reader, errorsx.Error) {
    contents, err := os.ReadFile(path)
    if err != nil {
        return nil, errorsx.Wrap(err, "path", path)
    }

    return bytes.NewReader(contents), nil
}
```
