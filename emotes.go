package goffz

import (
	"strconv"

	"github.com/dankeroni/jsonapi"
)

// EmoteData json to go
type EmoteData struct {
	CSS     interface{} `json:"css"`
	Height  int         `json:"height"`
	Hidden  bool        `json:"hidden"`
	ID      int         `json:"id"`
	Margins interface{} `json:"margins"`
	Name    string      `json:"name"`
	Owner   struct {
		DisplayName string `json:"display_name"`
		ID          int    `json:"id"`
		Name        string `json:"name"`
	} `json:"owner"`
	Public bool `json:"public"`
	Urls   struct {
		Num1 string `json:"1"`
	} `json:"urls"`
	Width int `json:"width"`
}

// EmoteResponse json to struct
type EmoteResponse struct {
	Emote EmoteData `json:"emote"`
}

// EmoticonsResponse json to struct
type EmoticonsResponse struct {
	Links struct {
		Next string `json:"next"`
		Self string `json:"self"`
	} `json:"_links"`
	Pages     int         `json:"_pages"`
	Total     int         `json:"_total"`
	Emoticons []EmoteData `json:"emoticons"`
}

// GetEmote request for GET http://api.frankerfacez.com/v1/emote/:emoteID
func (ffzAPI *FFZAPI) GetEmote(emoteID int, onSuccess func(EmoteResponse),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	var emote EmoteResponse
	onSuccessfulRequest := func() {
		onSuccess(emote)
	}
	ffzAPI.Get("/emote/"+strconv.Itoa(emoteID), nil, &emote, onSuccessfulRequest,
		onHTTPError, onInternalError)
}

// GetEmoticons request for GET http://api.frankerfacez.com/v1/emoticons
func (ffzAPI *FFZAPI) GetEmoticons(onSuccess func(EmoticonsResponse), onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) {
	var emoticons EmoticonsResponse
	onSuccessfulRequest := func() {
		onSuccess(emoticons)
	}
	ffzAPI.Get("/emoticons", nil, &emoticons, onSuccessfulRequest,
		onHTTPError, onInternalError)
}
