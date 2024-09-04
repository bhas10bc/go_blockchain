package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"y/types"
)

type PrivateKey struct {
	key *ecdsa.PrivateKey
}

func (k PrivateKey) Sign(data []byte ) (*Signature , error){
	r, s, err := ecdsa.Sign(rand.Reader, k.key, data)
	if err != nil {
		return nil, err
	}

	return &Signature{R:r,S:s}, nil
}

func GeneratePrivateKey() PrivateKey {
	key , err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic (err)
	}

	return PrivateKey{
		key:key,
	}
}

type PublicKey []byte

func (k PrivateKey) PublicKey() PublicKey {
	return elliptic.MarshalCompressed(k.key.PublicKey, k.key.PublicKey.X, k.key.PublicKey.Y)
}
func (k PublicKey) String() string {
	return hex.EncodeToString(k)
}

func (k PublicKey) Address() types.Address {
	h := sha256.Sum256(k)

	return types.AddressFromBytes(h[len(h)-20:])
}

type Signature struct {
	S *big.Int
	R *big.Int
}

func (sig Signature) Verify(pubKey PublicKey, data []byte) bool {
	x, y := elliptic.UnmarshalCompressed(elliptic.P256(), pubKey)
	key := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	return ecdsa.Verify(key, data, sig.R, sig.S)
}