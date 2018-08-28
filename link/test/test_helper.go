package perfectlinktest

import (
	"distributed-programming-abstractions/link"
	"testing"
	"time"
)

// ReliableDelivery tests reliable delivery of message sent from p to q.
func ReliableDelivery(p1 link.Perfect, q1 link.Perfect, t *testing.T) {
	var done = make(chan struct{})
	peerp := p1.(link.Peer)
	peerq := q1.(link.Peer)

	q1.Deliver(func(p link.Peer, m link.Message) {
		if string(m.([]byte)) != "Hello" {
			t.Error("Delivered message is not the sent message.")
		}

		if p.Id() != peerp.Id() {
			t.Error("The message was not sent by the peer.")
		}

		done <- struct{}{}
	})

	err := p1.Send(peerq, []byte("Hello"))
	if err != nil {
		t.Error(err)
	}
	select {
	case <-done:

	case <-time.After(100 * time.Millisecond):
		t.Error("Time message was not delivered within 100 ms.")
	}
}

// SelfDelivery tests if the message sent by the node to itself is delivered.
func SelfDelivery(p1 link.Perfect, t *testing.T) {
	var done = make(chan struct{})

	peerp := p1.(link.Peer)
	p1.Deliver(func(p link.Peer, m link.Message) {
		if string(m.([]byte)) != "Hello" {
			t.Error("Delivered message not send by itself.")
		}

		if p.Id() != peerp.Id() {
			t.Error("Message received from another peer and not self.")
		}

		done <- struct{}{}
	})
	err := p1.Send(peerp, []byte("Hello"))
	if err != nil {
		t.Error(err)
	}
	select {
	case <-done:

	case <-time.After(100 * time.Millisecond):
		t.Error("The message not delivered in 100 ms.")
	}
}
