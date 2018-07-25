package link

import "net"

type Peer interface {
	Id() string     // returns the id of the string.
	Addr() net.Addr // returns the network address of the node.
}
