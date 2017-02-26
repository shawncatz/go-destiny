package destiny

import (
	"fmt"
	"testing"
)

func TestCharacters(t *testing.T) {
	for _, c := range account.Characters {
		fmt.Println(c.String())
	}
}
