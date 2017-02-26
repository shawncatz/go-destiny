package destiny

import (
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

var (
	client  *Client
	opt     *ClientOptions
	account *Account
	err     error
)

func init() {
	opt = &ClientOptions{
		API:      os.Getenv("DESTINY_API"),
		Platform: PSN,
		Account:  os.Getenv("DESTINY_ACCOUNT"),
	}
	client = New(opt)
	account, err = client.Account()
	if err != nil {
		panic(err)
	}
}

func TestNew(t *testing.T) {
	if client.options.Platform != PSN {
		t.Fatalf("Platform = %d, want %d", client.options.Platform, PSN)
	}
}

//func TestGrimoire(t *testing.T) {
//	res, err := client.Grimoire("4611686018450625633")
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	fmt.Print(res.String())
//}

func TestClient_RequestPlayer(t *testing.T) {
	list, err := client.Player(PSN, "xenonsoul")
	if err != nil {
		t.Fatal(err)
	}

	if len(list) == 0 || len(list) > 1 {
		t.Fatal("wrong number of players found")
	}

	p := list[0]
	if p.DisplayName != "xenonsoul" {
		t.Errorf("DisplayName = %s, want %s", p.DisplayName, "xenonsoul")
	}

	if p.MembershipID != opt.Account {
		t.Errorf("MembershipID = %s, want %s", p.MembershipID, opt.Account)
	}
}

//func TestClient_Characters(t *testing.T) {
//	res, err := client.Characters()
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	fmt.Print(res.String())
//}

//func TestClient_Stats(t *testing.T) {
//	res, err := client.Stats(client.options.Platform, client.options.Account)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	fmt.Print(res.String())
//}

//func TestClient_Summary(t *testing.T) {
//	res, err := client.Summary(client.options.Platform, client.options.Account)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	prettyJSON(res.Body())
//}
