package destiny

type GrimoireResponse struct {
	ErrorCode       int64    `json:"ErrorCode"`
	ErrorStatus     string   `json:"ErrorStatus"`
	Message         string   `json:"Message"`
	MessageData     struct{} `json:"MessageData"`
	ThrottleSeconds int64    `json:"ThrottleSeconds"`
	Response        struct {
		Grimoire Grimoire `json:"data"`
	} `json:"Response"`
}

type Grimoire struct {
	Score   int             `json:"score"`
	Cards   []GrimoireCard  `json:"cardCollection"`
	Hidden  []GrimoireCard  `json:"carsToHide"`
	Bonuses []GrimoireBonus `json:"bonuses"`
}

type GrimoireCard struct {
	ID     int64          `json:"cardId"`
	Score  int            `json:"score"`
	Points int            `json:"points"`
	Stats  []GrimoireStat `json:"statisticsCollection"`
}

type GrimoireStat struct {
	Number       int            `json:"statNumber"`
	Value        float64        `json:"value"`
	DisplayValue string         `json:"displayValue"`
	Ranks        []GrimoireRank `json:"rankCollection"`
}

type GrimoireRank struct {
	Rank   int `json:"rank"`
	Points int `json:"points"`
}

type GrimoireBonus struct {
	CardID           int64              `json:"cardId"`
	CardName         string             `json:"cardName"`
	StatName         string             `json:"statName"`
	BonusName        string             `json:"bonusName"`
	BonusDescription string             `json:"bonusDescription"`
	BonusRank        GrimoireBonusRank  `json:"bonusRank"`
	Value            float64            `json:"value"`
	Threshold        float64            `json:"threshold"`
	SmallImage       GrimoireBonusImage `json:"smallImage"`
}

type GrimoireBonusRank struct {
	StatID int `json:"statId"`
	Rank   int `json:"rank"`
}

type GrimoireBonusImage struct {
	Rect struct {
		X      int `json:"x"`
		Y      int `json:"y"`
		Height int `json:"height"`
		Width  int `json:"width"`
	} `json:"rect"`
	SheetPath string `json:"sheetPath"`
	SheetSize struct {
		X      int `json:"x"`
		Y      int `json:"y"`
		Height int `json:"height"`
		Width  int `json:"width"`
	} `json:"sheetSize"`
}
