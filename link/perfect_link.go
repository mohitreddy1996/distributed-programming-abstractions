package link

type Perfect interface {
	Send(q Peer, m Message) error    // Send message m to peer q.
	Deliver(func(p Peer, m Message)) // Delivers message m from p.
}
