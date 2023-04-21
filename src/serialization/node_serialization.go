package serialization

import (
	"encoding/json"

	"github.com/Rosa-Devs/Rosa-Router/src/manager"
)

func SerializateNode(Pubkey string, Ip string, Port string, Rating string, Hs string, HsPort string) []byte {
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
		return nil
	}
	return b
}

func DeserializeNode(jsonStr string) (manager.Node, error) {
	var node manager.Node
	err := json.Unmarshal([]byte(jsonStr), &node)
	if err != nil {
		return manager.Node{}, err
	}
	return node, nil
}
