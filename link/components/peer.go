package component

import "net"

type Peer struct {
	id   string
	addr net.Addr
}

// instantiate and return new peer.
func NewPeer(id string, addr net.Addr) *Peer {
	return &Peer{
		id:   id,
		addr: addr,
	}
}

// Interface functions.
func (p *Peer) Id() string {
	return p.id
}

func (p *Peer) Addr() net.Addr {
	return p.addr
}
