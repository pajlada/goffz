package goffz

import "github.com/dankeroni/jsonapi"

// RoomResponse json to go
type RoomResponse struct {
	Room struct {
		InternalID     int         `json:"_id"`
		Tid            int         `json:"_tid"`
		CSS            interface{} `json:"css"`
		DisplayName    string      `json:"display_name"`
		ID             string      `json:"id"`
		IsGroup        bool        `json:"is_group"`
		ModeratorBadge interface{} `json:"moderator_badge"`
		Set            int         `json:"set"`
	} `json:"room"`
	Sets struct {
		Num1 struct {
			Type        int         `json:"_type"`
			CSS         interface{} `json:"css"`
			Description interface{} `json:"description"`
			Emoticons   []EmoteData `json:"emoticons"`
			Icon        interface{} `json:"icon"`
			ID          int         `json:"id"`
			Title       string      `json:"title"`
		} `json:"1"`
	} `json:"sets"`
}

// GetRoom request for GET https://api.frankerfacez.com/v1/room/:roomName
func (ffzAPI *FFZAPI) GetRoom(roomName string, onSuccess func(RoomResponse),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	var room RoomResponse
	onSuccessfulRequest := func() {
		onSuccess(room)
	}
	ffzAPI.Get("/room/"+roomName, nil, &room, onSuccessfulRequest,
		onHTTPError, onInternalError)
}
