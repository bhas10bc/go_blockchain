package network

import (
	"fmt"
	"time"
	"y/core"
	"y/crypto"

	"github.com/sirupsen/logrus"
)

var defaultBlockTime = 5 * time.Second

type ServerOpts struct {
	Transports []Transport
	BlockTime time.Duration
	PrivateKey *crypto.PrivateKey
}

type Server struct {
	ServerOpts
	blockTime time.Duration
	memPool *TxPool
	isValidator bool
	rpcCh chan RPC
	quitCh chan struct{}
}

func NewServer(opts ServerOpts) *Server {
	if opts.BlockTime == time.Duration(0){
		opts.BlockTime = defaultBlockTime
	}
	return &Server{
		ServerOpts: opts,
		blockTime: opts.BlockTime,
		memPool: NewTxPool(),
		isValidator: opts.PrivateKey != nil,
		rpcCh:      make(chan RPC),
		quitCh: make(chan struct{}, 1),
	}
}

func (s *Server) Start() {
	s.initTransports()
	ticker := time.NewTicker(s.blockTime)
free:
	for {
		select {
		case rpc := <-s.rpcCh:
			fmt.Println(rpc)

		case <-s.quitCh:
			break free
		case <-ticker.C:
			if s.isValidator{
				s.createNewBlock()
			}
			
		}
	}

	fmt.Println("server shutdown")
}

func (s *Server) initTransports() {
	for _, tr := range s.Transports {
		go func(tr Transport){
			for rpc := range tr.Consume(){
				s.rpcCh <- rpc
			}
		}(tr)

	}
}

func (s *Server) createNewBlock() error {
	fmt.Println("Creating a new block")
	return nil
}

func (s *Server) handleTransaction(tx *core.Transaction)error{
	if err := tx.Verify(); err != nil {
		return err
	}
	hash := tx.Hash(core.TxHasher{})
	if s.memPool.Has(hash){
		logrus.WithFields(logrus.Fields{
			"hash": hash,
		}).Info("transaction already in mempool")
		return nil
	}
	logrus.WithFields(logrus.Fields{
		"hash": tx.Hash(core.TxHasher{}),
	}).Info("adding new tx to the mempool")
	return s.memPool.Add(tx)
}