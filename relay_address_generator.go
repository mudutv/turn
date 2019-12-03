package turn

import (
	"net"

	"github.com/pion/transport/vnet"
)

// RelayAddressGenerator is used by a RelayConfig to generate a RelayAddress when creating an allocation.
// You can use one of the provided ones or provide your own.
type RelayAddressGenerator interface {
	Validate() error
	Allocate() (net.PacketConn, net.IP, int, error)
}

// RelayAddressGeneratorStatic can be used to return static IP address each time a relay is created.
// This can be used when you have a single static IP address that you want to use
type RelayAddressGeneratorStatic struct {
	// RelayAddress is the IP returned to the user when the relay is created
	RelayAddress net.IP

	// Net can be an instance of vnet if testing in a virtual environment
	Net *vnet.Net

	// Network, Address are the arguments passed to ListenPacket
	Network, Address string
}

// Validate is caled on server startup and confirms the RelayAddressGenerator is properly configured
func (r *RelayAddressGeneratorStatic) Validate() error {
	switch {
	case r.RelayAddress == nil:
		return errRelayAddressInvalid
	case r.Network == "" || r.Address == "":
		return errListeningAddressInvalid
	default:
		return nil
	}
}

// Allocate generates a new PacketConn to receive traffic on and the IP/Port to populate the allocation response with
func (r *RelayAddressGeneratorStatic) Allocate() (net.PacketConn, net.IP, int, error) {
	return nil, nil, 0, nil
}
