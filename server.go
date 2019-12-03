package turn

import (
	"fmt"
	"net"

	"github.com/pion/logging"
)

const (
	inboundMTU = 1500
)

// Server is an instance of the Pion TURN Server
type Server struct {
	log logging.LeveledLogger
}

// NewServer creates the Pion TURN server
func NewServer(config ServerConfig) (*Server, error) {
	if err := config.validate(); err != nil {
		return nil, err
	}

	s := &Server{
		log: config.LoggerFactory.NewLogger("turn"),
	}

	for _, p := range config.PacketConnConfigs {
		go s.connReadLoop(p.PacketConn, p.RelayAddressGenerator)
	}

	return s, nil
}

// Close stops the TURN Server. It cleans up any associated state and closes all connections it is managing
func (s *Server) Close() error {
	return nil
}

func (s *Server) connReadLoop(p net.PacketConn, r RelayAddressGenerator) {
	buf := make([]byte, inboundMTU)
	for {
		n, addr, err := p.ReadFrom(buf)
		if err != nil {
			s.log.Debugf("exit read loop on error: %s", err.Error())
			return
		}

		fmt.Println(n)
		fmt.Println(addr)
	}

}
