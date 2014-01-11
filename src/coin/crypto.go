package sb_coin

import (
	//"fmt"
	"errors"
	"log"
)

import "../lib/secp256k1-go"

type SecKey struct {
	Value [32]byte
}

type PubKey struct {
	Value [33]byte
}

type Sig struct {
	Value [64 + 1]byte
}

/*
	Helper Functions
*/

func (g *SecKey) Set(b []byte) {
	if len(b) != 32 {
		log.Panic()
	}
	copy(g.Value[0:32], b[0:32])
}

func (g *PubKey) Set(b []byte) {
	if len(b) != 33 {
		log.Panic()
	}
	copy(g.Value[0:33], b[0:33])
}

func (g *Sig) Set(b []byte) {
	if len(b) != 65 {
		log.Panic()
	}
	copy(g.Value[0:65], b[0:65])
}

/*
	Checks whether pubkey corresponding to address hash signed hash
	- recovers the pubkey from sig and hash
	- fail if pubkey cannot be be recovered
	- computes the address from the pubkey
	- fail if recovered address does not match pubkey hash
	- verify that signature is valid for hash for pubkey
*/
func ChkSig(address Address, hash SHA256, sig Sig) error {
	rawpubkey := secp256.RecoverPubkey(hash.Value[:], sig.Value[:])
	if rawpubkey == nil {
		return errors.New("ChkSig Error: signature invalid, pubkey recovery failed")
	}
	if address != AddressFromRawPubkey(rawpubkey) {
		return errors.New("ChkSig Error: signature invalid, address does not match output address")
	}
	if secp256.VerifySignature(hash.Value[:], sig.Value[:], rawpubkey) != 1 {
		return errors.New("ChkSig Error: signature invalid, signature invalid for hash")
	}
	return nil
}

func SignHash(hash SHA256, sec SecKey) (Sig, error) {
	sig1 := secp256.Sign(hash.Value[:], sec.Value[:])
	if sig1 == nil {
		log.Panic()
		return Sig{}, errors.New("SignHash invalid private key")
	}
	var sig2 Sig
	sig2.Set(sig1)
	return sig2, nil
}

//implement
func PubkeyFromSec(sec SecKey) PubKey {
	return PubKey{}
}
