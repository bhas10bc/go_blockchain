package network

import (
	"testing"
	"y/core"

	"github.com/stretchr/testify/assert"
)

func TestTxPool(t *testing.T) {
	p := NewTxPool()
	assert.Equal(t,p.Len(), 0 )
}

func TestTxPoolAddTx(t *testing.T){
	p := NewTxPool()

	tx:= core.NewTransaction([]byte ("test"))
	assert.Nil(t,p.Add(tx))
	assert.Equal(t,p.Len(), 1)

	_ = core.NewTransaction([]byte ("test"))
	assert.Equal(t,p.Len(),1)

	p.Flush()
	assert.Equal(t,p.Len(), 0 )
}