# goffz

Heavily based off dankeroni's gotwitch

## Example for getting a Stream object
```go
package main

import (
    "fmt"

    "github.com/dankeroni/gotwitch"
)

var api = gotwitch.New("<ClientID>")

func main() {
    api.GetStream("pajlada", onSuccess, onHTTPError, onInternalError)
}

func onSuccess(stream gotwitch.Stream) {
    fmt.Printf("%+v\n", stream)
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

## Example for unimplemented requests
```go
package main

import (
    "fmt"

    "github.com/dankeroni/gotwitch"
    "net/url"
)

// GamesTop https://api.twitch.tv/kraken/games/top (https://mholt.github.io/json-to-go/)
type GamesTop struct {
    Total int `json:"_total"`
    Top []struct {
        Game struct {
            Name        string `json:"name"`
            Popularity  int    `json:"popularity"`
            ID          int    `json:"_id"`
            GiantbombID int    `json:"giantbomb_id"`
            Box         struct {
                Large    string `json:"large"`
                Medium   string `json:"medium"`
                Small    string `json:"small"`
                Template string `json:"template"`
            } `json:"box"`
            Logo struct {
                Large    string `json:"large"`
                Medium   string `json:"medium"`
                Small    string `json:"small"`
                Template string `json:"template"`
            } `json:"logo"`
        } `json:"game"`
        Viewers  int `json:"viewers"`
        Channels int `json:"channels"`
    } `json:"top"`
}

var api = gotwitch.New("<ClientID>")

func main() {
    var gamesTop GamesTop
    requestParameters := url.Values{}
    requestParameters.Add("limit", "1")
    requestParameters.Add("offset", "4")
    onSuccess := func() {
        fmt.Printf("%+v", gamesTop)
    }

    // The oauthToken is optional and can be replaced with an empty string
    // requestParameters can also be nil
    api.Get("/games/top", requestParameters, "<oauthToken>", &gamesTop,
        onSuccess, onHTTPError, onInternalError)
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
