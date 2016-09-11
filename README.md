# goffz

Uses [jsonapi](https://github.com/dankeroni/jsonapi) by [dankeroni](https://github.com/dankeroni) for all net code.

Inspired by [gotwitch](https://github.com/dankeroni/gotwitch)

## Example for getting a Room object
```go
package main

import (
    "fmt"

    "github.com/pajlada/goffz"
)

var api = goffz.New()

func main() {
    api.GetRoom("pajlada", onSuccess, onHTTPError, onInternalError)
}

func onSuccess(room goffz.RoomResponse) {
    fmt.Printf("%+v\n", room)
}

func onHTTPError(statusCode int, statusMessage, errorMessage string) {
    fmt.Println("statusCode:", statusCode)
    fmt.Println("statusMessage:", statusMessage)
    fmt.Println("errorMessage:", errorMessage)
}

func onInternalError(err error) {
    fmt.Println("internalError:", err)
}
```
