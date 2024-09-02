package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct {
	addr  NetAddr
	consumeCh chan RPC
	lock  sync.RWMutex
	peers map[NetAddr]*LocalTransport
}

func NewLocalTransport(addr NetAddr) Transport {
	return &LocalTransport {
		addr: addr,
		consumeCh: make(chan RPC, 1024),
		peers: make(map[NetAddr]*LocalTransport),
	}
}

func (t *LocalTransport) Consume() <-chan RPC {
	return t.consumeCh
}

func (t *LocalTransport) Connect(tr Transport) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.peers[tr.Addr()] = tr.(*LocalTransport)

	return nil
}

func (t *LocalTransport) SendMessage(to NetAddr, Payload []byte) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	peer, ok := t.peers[to]
	if !ok {
		return fmt.Errorf("%s: could not send message to %s",t.addr, to)
	}

	peer.consumeCh <- RPC {
		From: t.addr,
		Payload: Payload,
	}
	return nil
}

func (t *LocalTransport) Addr() NetAddr {
	return t.addr
}