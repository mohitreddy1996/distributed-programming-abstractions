package link

import "net"

type Node interface {
	Members() []Peer        // Returns all the peers including itself.
	Peers() []Peer          // Returns the list of peers of the node.
	AddPeer(p Peer)         // Adds a peer to the list.
	NodeCount() int         // Returns the number of nodes in the network.
	Listener() net.Listener //Returns the network listener.

	Peer
}
