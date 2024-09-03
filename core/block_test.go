package core

import (
	"testing"
	"time"
	"y/crypto"
	"y/types"

	"github.com/stretchr/testify/assert"
)


func randomBlock(height uint32, prevBlockHash types.Hash) *Block {
	header := &Header{
		Version: 1,
		PrevBlockHash: prevBlockHash,
		Height: height,
		Timestamp: time.Now().UnixNano(),
	}
	tx := Transaction{
		Data :[]byte("test"),
	}

	return NewBlock(header,[]Transaction{tx})
}

func randomBlockWithSignature(t *testing.T, height uint32, prevBlockHash types.Hash) *Block {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(height,prevBlockHash)
	tx := randomTxWithSignature(t)
	b.AddTransaction(tx)
	assert.Nil(t,b.Sign(privKey))
	return b
}

func TestSignBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0, types.Hash{})
	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)
}

func TestVerifyBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0,types.Hash{})

	assert.Nil(t, b.Sign(privKey))
	assert.Nil(t, b.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivKey.PublicKey()

	assert.NotNil(t,b.Verify())

	b.Height = 100
	assert.NotNil(t,b.Verify())
}
