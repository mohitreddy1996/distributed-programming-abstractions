package perfectlink

import (
	"distributed-programming-abstractions/link"
	"distributed-programming-abstractions/link/components"
	"net/rpc"
	"time"
)

// Node implements perfect p2p link nodes.
type Node struct {
	deliver    func(link.Peer, link.Message)
	keepalive  time.Duration
	connection map[string]*rpc.Client
	channel    string
	mutexChan  chan func()
	link.Node
}

// WithDefault instantiates Node with default settings.
var WithDefault = func(n *Node) {
	n.Node = component.New(component.WithDefault)
}

// WithChannel instantiates Node with the given channel.
func WithChannel(channel string) func(*Node) {
	return func(n *Node) {
		n.channel = channel
	}
}

// WithKeepalive instantiates Node with the given keepalive time.
func WithKeepalive(keepalive time.Duration) func(*Node) {
	return func(n *Node) {
		n.keepalive = keepalive
	}
}

func (n *Node) init() {

}

func (n *Node) runbackground() {

}

// Deliver sets deliver for Node with given function.
func (n *Node) Deliver(f func(link.Peer, link.Message)) {
	n.deliver = f
}

// Send sends m to peer q of n.
func (n *Node) Send(q link.Peer, m link.Message) error {
	return nil
}

func (n *Node) connect(q link.Peer) (*rpc.Client, error) {
	return nil, nil
}

// New instantiates a new Node with capable of connecting to peers using tcp (rpc)
// based connection.
func New(configs ...func(*Node)) *Node {
	n := &Node{
		keepalive:  1 * time.Second,
		channel:    "p2p",
		deliver:    func(link.Peer, link.Message) {},
		connection: make(map[string]*rpc.Client),
		mutexChan:  make(chan func()),
	}

	for _, config := range configs {
		config(n)
	}

	n.init()
	go n.runbackground()
	return n
}
