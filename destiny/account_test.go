package destiny

import "testing"

func TestAccount(t *testing.T) {
	account, err := client.Account()
	if err != nil {
		t.Fatal(err)
	}

	if account.MembershipID != client.options.Account {
		t.Errorf("MembershipID = %s, want %s", account.MembershipID, client.options.Account)
	}

	//fmt.Printf("%#v\n", account)
}

func TestAccount_CharacterIDs(t *testing.T) {
	account, err := client.Account()
	if err != nil {
		t.Fatal(err)
	}

	if len(account.CharacterIDs()) < 3 {
		t.Fatal("CharacterIDs() < 3")
	}
}
