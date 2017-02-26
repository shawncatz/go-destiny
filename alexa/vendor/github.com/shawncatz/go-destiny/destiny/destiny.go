package destiny

type Client struct {
	URL           string // base URL of API requests
	API           string // API key
	Username      string // Username
	Password      string // Password
	CharacterType int    // Character type (1 = xbox, 2 = psn, 254 = bungie)
}

func New(api, user, pass string, ctype int) *Client {
	return &Client{
		API:           api,
		Username:      user,
		Password:      pass,
		CharacterType: ctype,
	}
}
