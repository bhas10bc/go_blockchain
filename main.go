package main

import (
	"bytes"
	"math/rand"
	"strconv"
	"time"
	"y/core"
	"y/crypto"
	"y/network"

	"github.com/sirupsen/logrus"
)

func main() {
	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func(){
		for{
		// trRemote.SendMessage(trLocal.Addr(), []byte("hello world"))
		if err := sendTransaction(trRemote, trLocal.Addr()); err != nil {
			logrus.Error(err)
		}
		time.Sleep(1 * time.Second)
		}
	}()

	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	s:= network.NewServer(opts)
	s.Start()
}

func sendTransaction(tr network.Transport, to network.NetAddr) error {
	privKey := crypto.GeneratePrivateKey()
	data := []byte(strconv.FormatInt(int64(rand.Intn(1000)), 10))
	tx := core.NewTransaction(data)
	tx.Sign(privKey)

	buf := &bytes.Buffer{}
	if err := tx.Encode(core.NewGobTxEncoder(buf)); err != nil {
		return err
	}
	
	msg :=  network.NewMessage(network.MessageTypeTx,buf.Bytes())
	return tr.SendMessage(to, msg.Bytes())
}