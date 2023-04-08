package serialization

import (
	"encoding/json"

	"github.com/Mihalic2040/Rosa-Router/src/manager"
)

func SerializateNode(Pubkey string, Ip string, Port string, Rating string, Hs bool, HsPort string) ([]byte, error) {
	p := manager.Node{
		Pubkey: Pubkey,
		Ip:     Ip,
		Port:   Port,
		Rating: Rating,
		Hs:     Hs,
		HsPort: HsPort,
	}

	b, err := json.Marshal(p)
	if err != nil {
		return err
	}
	return b
}
