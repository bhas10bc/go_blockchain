package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privkey := GeneratePrivateKey()
	pubKey :=  privkey.PublicKey()
	// address := pubKey.Address()

	msg := []byte("hello world")
	sig, err := privkey.Sign(msg)
	assert.Nil(t,err)
	b := sig.Verify(pubKey,msg)
	assert.True(t,b )


	fmt.Println(sig)
}

func TestKeyPair_Sign_Verify(t *testing.T){
	privkey := GeneratePrivateKey()
	pubKey :=  privkey.PublicKey()
	msg := []byte("hello world")

	sig , err := privkey.Sign(msg)
	assert.Nil(t, err)

	assert.True(t, sig.Verify(pubKey,msg))
}


func TestKeyPair_Sign_Verify_Fail(t *testing.T){
	privkey := GeneratePrivateKey()
	Pubkey := privkey.PublicKey()
	msg := []byte("hello world")

	otherprivkey := GeneratePrivateKey()
	otherPubkey := otherprivkey.PublicKey()

	sig , err := privkey.Sign(msg)
	assert.Nil(t, err)

	assert.False(t, sig.Verify(otherPubkey,msg))
	assert.False(t, sig.Verify(Pubkey,[]byte("dvcedw")))
}