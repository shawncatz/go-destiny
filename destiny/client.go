package destiny

import (
	"fmt"
	"net/http"
	"os"

	"gopkg.in/resty.v0"
)

const URL = "https://www.bungie.net/platform/destiny"

type Client struct {
	options *ClientOptions

	rest    *resty.Client
	cookies []*http.Cookie
}

type ClientOptions struct {
	API      string // API key
	Platform Platform
	Account  string
	Titan    string
	Warlock  string
	Hunter   string
}

func New(options *ClientOptions) *Client {
	r := resty.New()
	r.RetryCount = 0
	r.SetRedirectPolicy(resty.FlexibleRedirectPolicy(5))

	return &Client{
		options: options,
		rest:    r,
	}
}

func (c *Client) Account() (*Account, error) {
	res, err := c.RequestSummary(c.options.Platform, c.options.Account)
	if err != nil {
		return nil, err
	}

	return NewAccountFromJSON(res.Body(), c)
}

func (c *Client) Player(mt Platform, name string) ([]*Player, error) {
	res, err := c.RequestPlayer(mt, name)
	if err != nil {
		return nil, err
	}

	return NewPlayerFromJSON(res.Body())
}

func (c *Client) RequestGrimoire(id string) (*resty.Response, error) {
	return c.execute(resty.MethodGet, fmt.Sprintf("vanguard/grimoire/%d/%s/", c.options.Platform, id))
}

func (c *Client) RequestPlayer(mt Platform, name string) (*resty.Response, error) {
	return c.execute(resty.MethodGet, fmt.Sprintf("searchdestinyplayer/%s/%s/", mt, name))
}

func (c *Client) RequestStats(mt Platform, id string) (*resty.Response, error) {
	return c.execute(resty.MethodGet, fmt.Sprintf("Stats/Account/%d/%s/", mt, id))
}

func (c *Client) RequestSummary(mt Platform, id string) (*resty.Response, error) {
	return c.execute(resty.MethodGet, fmt.Sprintf("%d/Account/%s/Summary/", mt, id))
}

func (c *Client) RequestVault(mt Platform) (*resty.Response, error) {
	return c.execute(resty.MethodGet, fmt.Sprintf("%d/MyAccount/Vault/Summary/", mt))
}

func (c *Client) execute(method, path string) (*resty.Response, error) {
	url := URL + "/" + path

	res, err := c.rest.R().
		SetHeader("X-API-Key", os.Getenv("DESTINY_API")).
		SetHeader("Accept", "application/json").
		Execute(method, url)

	if err != nil {
		return nil, err
	}

	if res.RawResponse.StatusCode != 200 {
		return nil, fmt.Errorf("response %d code", res.RawResponse.StatusCode)
	}

	c.cookies = res.Cookies()

	return res, nil
}
