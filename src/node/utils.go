package node

import (
	"fmt"
	"strings"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/multiformats/go-multiaddr"
)

// this func getting a port number from host struc
func GetHostPort(host host.Host) (string, error) {
	addrs := host.Addrs()

	for _, addr := range addrs {
		// Check if the multiaddress represents TCP
		if strings.Contains(addr.String(), "/ip4") || strings.Contains(addr.String(), "/ip6") {
			// Extract the TCP port from the multiaddress
			maddr, err := multiaddr.NewMultiaddr(addr.String())
			if err != nil {
				return "", err
			}
			tcp, err := maddr.ValueForProtocol(multiaddr.P_TCP)
			if err != nil {
				return "", err
			}
			return tcp, nil
		}
	}

	return "", fmt.Errorf("no TCP multiaddress found")
}
