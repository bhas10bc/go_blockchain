package core

import (
	"testing"
	"time"

	"y/crypto"
	"y/types"

	"github.com/stretchr/testify/assert"
)

func TestSignBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(t, 0, types.Hash{})

	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)
}

func TestVerifyBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(t, 0, types.Hash{})

	assert.Nil(t, b.Sign(privKey))
	assert.Nil(t, b.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivKey.PublicKey()
	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())
}

func randomBlock(t *testing.T, height uint32, prevBlockHash types.Hash) *Block {
	privKey := crypto.GeneratePrivateKey()
	tx := randomTxWithSignature(t)
	header := &Header{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		Height:        height,
		Timestamp:     time.Now().UnixNano(),
	}

	b, err := NewBlock(header, []Transaction{tx})
	assert.Nil(t, err)
	dataHash, err := CalculateDataHash(b.Transactions)
	assert.Nil(t, err)
	b.Header.DataHash = dataHash
	assert.Nil(t, b.Sign(privKey))

	return b
}