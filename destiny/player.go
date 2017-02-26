package destiny

import "encoding/json"

type PlayerResponse struct {
	ErrorCode       int64    `json:"ErrorCode"`
	ErrorStatus     string   `json:"ErrorStatus"`
	Message         string   `json:"Message"`
	MessageData     struct{} `json:"MessageData"`
	ThrottleSeconds int64    `json:"ThrottleSeconds"`
	Players         []Player `json:"Response"`
}

type Player struct {
	IconPath       string `json:"iconPath"`
	MembershipType int    `json:"membershipType"`
	MembershipID   string `json:"membershipId"`
	DisplayName    string `json:"displayName"`
}

func NewPlayerFromJSON(body []byte) ([]*Player, error) {
	r := &PlayerResponse{}

	err := json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}

	list := []*Player{}
	for _, p := range r.Players {
		list = append(list, &p)
	}

	return list, nil
}
