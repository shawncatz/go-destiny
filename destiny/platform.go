package destiny

import "fmt"

type Platform int

const (
	XBOX   Platform = 1
	PSN    Platform = 2
	BUNGIE Platform = 254
)

func (p Platform) String() string {
	return fmt.Sprintf("%d", int(p))
}
