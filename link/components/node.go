package component

import (
	"crypto/rand"
	"distributed-programming-abstractions/link"
	"fmt"
	"io"
	"log"
	"net"
)

// Node represents a node in the network and responsible for communication with
// peers.
type Node struct {
	peers    []link.Peer
	isset    map[string]bool
	listener net.Listener

	*Peer
}

// WithDefault instantiates node with default options.
var WithDefault = func(n *Node) {
	l, e := net.Listen("tcp", "129.1.1.1:0")
	if e != nil {
		log.Fatalf("net.Listen tcp :0: %v", e)
	}
	n.listener = l
	n.addr = l.Addr()
}

// WithListener instantiates node with listener.
func WithListener(l net.Listener) func(*Node) {
	return func(n *Node) {
		n.listener = l
		n.addr = l.Addr()
	}
}

// WithID instantiates node with given id.
func WithID(id string) func(*Node) {
	return func(n *Node) { n.id = id }
}

// WithAddr instantiates node with given address.
func WithAddr(addr net.Addr) func(*Node) {
	return func(n *Node) { n.addr = addr }
}

// WithPeer instantiates node with peer.
func WithPeer(p link.Peer) func(*Node) {
	return func(n *Node) { n.AddPeer(p) }
}

// New instantiates a new node.
func New(configs ...func(*Node)) *Node {
	n := &Node{
		isset: make(map[string]bool),
		peers: make([]link.Peer, 0),
		Peer:  &Peer{},
	}

	for _, config := range configs {
		config(n)
	}
	// set random uid.
	if n.Id() == "" {
		uid := make([]byte, 16)
		io.ReadFull(rand.Reader, uid)
		n.id = fmt.Sprintf("%X", uid)
	}
	return n
}

// Interface functions.
func (n *Node) Peers() []link.Peer {
	return n.peers
}

func (n *Node) Members() []link.Peer {
	return append(n.peers, n)
}

func (n *Node) Listener() net.Listener {
	return n.listener
}

func (n *Node) AddPeer(p link.Peer) {
	if n.isset[p.Id()] || p.Id() == n.Id() {
		return
	}
	n.isset[p.Id()] = true
	n.peers = append(n.peers, p)
}

func (n *Node) NodeCount() int {
	return len(n.peers) + 1
}
