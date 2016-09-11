package goffz

import "github.com/dankeroni/jsonapi"

// SetResponse json to go
type SetResponse struct {
	DefaultSets []int `json:"default_sets"`
	Sets        map[string]struct {
		Type        int         `json:"_type"`
		CSS         interface{} `json:"css"`
		Description interface{} `json:"description"`
		Emoticons   []EmoteData `json:"emoticons"`
		Icon        interface{} `json:"icon"`
		ID          int         `json:"id"`
		Title       string      `json:"title"`
	} `json:"sets"`
}

// GetSet request for GET https://api.frankerfacez.com/v1/set/:setName
func (ffzAPI *FFZAPI) GetSet(setName string, onSuccess func(SetResponse),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	var set SetResponse
	onSuccessfulRequest := func() {
		onSuccess(set)
	}
	ffzAPI.Get("/set/"+setName, nil, &set, onSuccessfulRequest,
		onHTTPError, onInternalError)
}
