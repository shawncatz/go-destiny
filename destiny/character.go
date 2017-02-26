package destiny

import "fmt"

var genders = []string{"Male", "Female"}
var classes = []string{"Titan", "Hunter", "Warlock"}

type Character struct {
	client *Client

	BackgroundPath     string  `json:"backgroundPath"`
	BaseCharacterLevel int64   `json:"baseCharacterLevel"`
	CharacterLevel     int64   `json:"characterLevel"`
	EmblemHash         int64   `json:"emblemHash"`
	EmblemPath         string  `json:"emblemPath"`
	IsPrestigeLevel    bool    `json:"isPrestigeLevel"`
	PercentToNextLevel float64 `json:"percentToNextLevel"`
	LevelProgression   struct {
		CurrentProgress     int64 `json:"currentProgress"`
		DailyProgress       int64 `json:"dailyProgress"`
		Level               int64 `json:"level"`
		NextLevelAt         int64 `json:"nextLevelAt"`
		ProgressToNextLevel int64 `json:"progressToNextLevel"`
		ProgressionHash     int64 `json:"progressionHash"`
		Step                int64 `json:"step"`
		WeeklyProgress      int64 `json:"weeklyProgress"`
	} `json:"levelProgression"`
	CharacterBase struct {
		BuildStatGroupHash  int64  `json:"buildStatGroupHash"`
		CharacterID         string `json:"characterId"`
		ClassHash           int64  `json:"classHash"`
		ClassType           int64  `json:"classType"`
		CurrentActivityHash int64  `json:"currentActivityHash"`
		Customization       struct {
			DecalColor   int64 `json:"decalColor"`
			DecalIndex   int64 `json:"decalIndex"`
			EyeColor     int64 `json:"eyeColor"`
			Face         int64 `json:"face"`
			FeatureColor int64 `json:"featureColor"`
			FeatureIndex int64 `json:"featureIndex"`
			HairColor    int64 `json:"hairColor"`
			HairIndex    int64 `json:"hairIndex"`
			LipColor     int64 `json:"lipColor"`
			Personality  int64 `json:"personality"`
			SkinColor    int64 `json:"skinColor"`
			WearHelmet   bool  `json:"wearHelmet"`
		} `json:"customization"`
		DateLastPlayed           string `json:"dateLastPlayed"`
		GenderHash               int64  `json:"genderHash"`
		GenderType               int64  `json:"genderType"`
		GrimoireScore            int64  `json:"grimoireScore"`
		LastCompletedStoryHash   int64  `json:"lastCompletedStoryHash"`
		MembershipID             string `json:"membershipId"`
		MembershipType           int64  `json:"membershipType"`
		MinutesPlayedThisSession string `json:"minutesPlayedThisSession"`
		MinutesPlayedTotal       string `json:"minutesPlayedTotal"`
		PeerView                 struct {
			Equipment []struct {
				Dyes []struct {
					ChannelHash int64 `json:"channelHash"`
					DyeHash     int64 `json:"dyeHash"`
				} `json:"dyes"`
				ItemHash int64 `json:"itemHash"`
			} `json:"equipment"`
		} `json:"peerView"`
		PowerLevel int64 `json:"powerLevel"`
		RaceHash   int64 `json:"raceHash"`
		Stats      struct {
			StatAgility struct {
				MaximumValue int64 `json:"maximumValue"`
				StatHash     int64 `json:"statHash"`
				Value        int64 `json:"value"`
			} `json:"STAT_AGILITY"`
			StatArmor struct {
				MaximumValue int64 `json:"maximumValue"`
				StatHash     int64 `json:"statHash"`
				Value        int64 `json:"value"`
			} `json:"STAT_ARMOR"`
			StatDefense struct {
				MaximumValue int64 `json:"maximumValue"`
				StatHash     int64 `json:"statHash"`
				Value        int64 `json:"value"`
			} `json:"STAT_DEFENSE"`
			StatDiscipline struct {
				MaximumValue int64 `json:"maximumValue"`
				StatHash     int64 `json:"statHash"`
				Value        int64 `json:"value"`
			} `json:"STAT_DISCIPLINE"`
			StatIntellect struct {
				MaximumValue int64 `json:"maximumValue"`
				StatHash     int64 `json:"statHash"`
				Value        int64 `json:"value"`
			} `json:"STAT_INTELLECT"`
			StatLight struct {
				MaximumValue int64 `json:"maximumValue"`
				StatHash     int64 `json:"statHash"`
				Value        int64 `json:"value"`
			} `json:"STAT_LIGHT"`
			StatOptics struct {
				MaximumValue int64 `json:"maximumValue"`
				StatHash     int64 `json:"statHash"`
				Value        int64 `json:"value"`
			} `json:"STAT_OPTICS"`
			StatRecovery struct {
				MaximumValue int64 `json:"maximumValue"`
				StatHash     int64 `json:"statHash"`
				Value        int64 `json:"value"`
			} `json:"STAT_RECOVERY"`
			StatStrength struct {
				MaximumValue int64 `json:"maximumValue"`
				StatHash     int64 `json:"statHash"`
				Value        int64 `json:"value"`
			} `json:"STAT_STRENGTH"`
		} `json:"stats"`
	} `json:"characterBase"`
}

func (c *Character) String() string {
	return fmt.Sprintf("character: %d (%d) %s %s %s", c.Level(), c.Light(), c.Gender(), c.Class(), c.LastPlayed())
}

func (c *Character) Gender() string {
	return genders[c.CharacterBase.GenderType]
}

func (c *Character) Class() string {
	return classes[c.CharacterBase.ClassType]
}

func (c *Character) Light() int64 {
	return c.CharacterBase.PowerLevel
}

func (c *Character) Level() int64 {
	return c.CharacterLevel
}

func (c *Character) LastPlayed() string {
	return c.CharacterBase.DateLastPlayed
}
