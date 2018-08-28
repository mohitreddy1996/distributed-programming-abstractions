package perfectlink_test

import "testing"
import "distributed-programming-abstractions/link/perfect_link"
import "distributed-programming-abstractions/link/test"

type BadAddr struct{}

func (_ *BadAddr) Network() string { return "tcp" }
func (_ *BadAddr) String() string  { return "bad-addr" }

func TestReliability(t *testing.T) {
	p := perfectlink.New(perfectlink.WithDefault)
	q := perfectlink.New(perfectlink.WithDefault)

	perfectlinktest.ReliableDelivery(p, q, t)
}

func TestSelfDelivery(t *testing.T) {
	p := perfectlink.New(perfectlink.WithDefault)

	perfectlinktest.SelfDelivery(p, t)
}
