package consensus

import (
	"bytes"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/coschain/contentos-go/prototype"
	"github.com/coschain/gobft/message"
)

/********* implements gobft IPubValidator ***********/

type publicValidator struct {
	sab         *SABFT
	accountName string
	pubKey *prototype.PublicKeyType
	bftPubKey message.PubKey
}

func newPubValidator(s *SABFT, pk *prototype.PublicKeyType, n string) *publicValidator {
	return &publicValidator{
		sab: s,
		accountName: n,
		pubKey: pk,
		bftPubKey: message.PubKey(pk.ToWIF()),
	}
}

func (pv *publicValidator) VerifySig(digest, signature []byte) bool {
	sig := &prototype.SignatureType{Sig:signature}
	if err := sig.Validate(); err != nil {
		return false
	}
	buffer, err := secp256k1.RecoverPubkey(digest, signature)
	if err != nil {
		pv.sab.log.Error(err)
		return false
	}
	ecPubKey, err := crypto.UnmarshalPubkey(buffer)
	if err != nil {
		pv.sab.log.Error(err)
		return false
	}
	keyBuffer := secp256k1.CompressPubkey(ecPubKey.X, ecPubKey.Y)
	recoveredKey := new(prototype.PublicKeyType)
	recoveredKey.Data = keyBuffer
	return bytes.Equal(pv.pubKey.Data, recoveredKey.Data)
}

func (pv *publicValidator) GetPubKey() message.PubKey {
	return pv.bftPubKey
}

func (pv *publicValidator) GetVotingPower() int64 {
	return 1
}

func (pv *publicValidator) SetVotingPower(int64) {

}

/********* end gobft IPubValidator ***********/

/********* implements gobft IPrivValidator ***********/

type privateValidator struct {
	accountName    string
	sab     *SABFT
	privKey *prototype.PrivateKeyType
	pubKey  *prototype.PublicKeyType
	bftPubKey message.PubKey
}

func newPrivValidator(s *SABFT, pk *prototype.PrivateKeyType, n string) *privateValidator {
	if pk == nil {
		return nil
	}
	pub, err := pk.PubKey()
	if err != nil {
		panic(err)
	}
	return &privateValidator{
		sab: s,
		privKey: pk,
		accountName: n,
		pubKey: pub,
		bftPubKey: message.PubKey(pub.ToWIF()),
	}
}

func (pv *privateValidator) Sign(digest []byte) []byte {
	res, err := secp256k1.Sign(digest[:], pv.privKey.Data)
	if err != nil {
		pv.sab.log.Error(err)
		return nil
	}

	if pv.sab.mockMalicious {
		res[0] = 0x01
		res[3] = 0x01
		res[7] = 0x01
	}

	return res
}

func (pv *privateValidator) GetPubKey() message.PubKey {
	return pv.bftPubKey
}

/********* end gobft IPrivValidator ***********/
