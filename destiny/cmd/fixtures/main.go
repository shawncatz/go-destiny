package main

// pull fixtures and store

import (
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/shawncatz/go-destiny/destiny"
	"github.com/shawncatz/go-destiny/utils"
	"gopkg.in/resty.v0"
)

func main() {
	var res *resty.Response
	var err error

	opt := &destiny.ClientOptions{
		API:      os.Getenv("DESTINY_API"),
		Platform: destiny.PSN,
		Account:  os.Getenv("DESTINY_ACCOUNT"),
	}
	client := destiny.New(opt)

	res, err = client.RequestSummary(opt.Platform, opt.Account)
	check(err)
	saveFile("account_summary.json", res.Body())

	res, err = client.RequestPlayer(opt.Platform, "xenonsoul")
	check(err)
	saveFile("find_player.json", res.Body())

	res, err = client.RequestGrimoire(opt.Account)
	check(err)
	saveFile("grimoire.json", res.Body())

	res, err = client.RequestVault(opt.Platform)
	check(err)
	saveFile("vault.json", res.Body())
}

func check(err error) {
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}

func saveFile(name string, data []byte) {
	js, err := utils.PrettyJSON(data)
	check(err)

	err = ioutil.WriteFile(fmt.Sprintf("fixtures/%s", name), js, 0644)
	check(err)
}
