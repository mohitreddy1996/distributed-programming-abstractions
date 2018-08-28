package perfectlink

import (
	"distributed-programming-abstractions/link"
	"distributed-programming-abstractions/link/components"
	"net"
	"net/rpc"
	"time"
)

// Payload is a wrapper over link.Message with id and address.
type Payload struct {
	ID      string
	Addr    net.Addr
	Message link.Message
}

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

// WithNode instantiates Node with the given node.
func WithNode(node link.Node) func(*Node) {
	return func(n *Node) {
		n.Node = node
	}
}

func (n *Node) init() {
	s := &server{n}
	go s.serve(n.channel)
}

func (n *Node) runbackground() {
	for f := range n.mutexChan {
		f()
	}
}

// Deliver sets deliver for Node with given function.
func (n *Node) Deliver(f func(link.Peer, link.Message)) {
	n.deliver = f
}

// Send sends m to peer q of n.
func (n *Node) Send(q link.Peer, m link.Message) error {
	result := make(chan error, 1)
	n.mutexChan <- func() {
		if q.Id() == n.Id() {
			go n.deliver(n.Node.(link.Peer), m)
			result <- nil
			return
		}
		c, err := n.connect(q)
		if err != nil {
			result <- err
			return
		}
		result <- c.Call(n.channel+".Recv", &Payload{n.Id(), n.Addr(), m}, nil)
		return
	}
	return <-result
}

func (n *Node) recv(p *Payload) {
	n.mutexChan <- func() {
		peer := component.NewPeer(p.ID, p.Addr)
		go n.deliver(peer, p.Message)
	}
}

func (n *Node) connect(q link.Peer) (*rpc.Client, error) {
	if _, ok := n.connection[q.Id()]; ok { // connection already exists.
		return n.connection[q.Id()], nil
	}
	addr, err := net.ResolveTCPAddr(q.Addr().Network(), q.Addr().String())
	if err != nil {
		return nil, err
	}
	c, err := net.DialTCP(q.Addr().Network(), nil, addr)
	if err != nil {
		return nil, err
	}
	c.SetKeepAlive(true)
	c.SetKeepAlivePeriod(n.keepalive)
	n.connection[q.Id()] = rpc.NewClient(c)
	return n.connection[q.Id()], nil
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
