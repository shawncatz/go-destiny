package destiny

import "encoding/json"

type AccountSummaryResponse struct {
	ErrorCode       int64    `json:"ErrorCode"`
	ErrorStatus     string   `json:"ErrorStatus"`
	Message         string   `json:"Message"`
	MessageData     struct{} `json:"MessageData"`
	ThrottleSeconds int64    `json:"ThrottleSeconds"`
	Response        struct {
		Account Account `json:"data"`
	} `json:"Response"`
}

type Account struct {
	client *Client

	Characters    []Character `json:"characters"`
	GrimoireScore int64       `json:"grimoireScore"`
	Inventory     struct {
		Currencies []struct {
			ItemHash int64 `json:"itemHash"`
			Value    int64 `json:"value"`
		} `json:"currencies"`
		Items []interface{} `json:"items"`
	} `json:"inventory"`
	MembershipID   string `json:"membershipId"`
	MembershipType int64  `json:"membershipType"`
	Versions       int64  `json:"versions"`
}

func NewAccountFromJSON(body []byte, c *Client) (*Account, error) {
	response := &AccountSummaryResponse{}

	err := json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	account := &response.Response.Account
	account.client = c

	for _, ch := range account.Characters {
		ch.client = c
	}
	return account, nil
}

// CharacterIDs returns a list of string character ids
func (a *Account) CharacterIDs() []string {
	ret := make([]string, 0)

	if len(a.Characters) == 0 {
		return ret
	}

	for _, c := range a.Characters {
		if c.CharacterBase.CharacterID == "" {
			continue
		}

		ret = append(ret, c.CharacterBase.CharacterID)
	}

	return ret
}
