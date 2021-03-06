package jwt

import (
	"crypto"
	"crypto/hmac"
	"hash"
	"sync"
)

// NewSignerHS returns a new HMAC-based signer.
func NewSignerHS(alg Algorithm, key []byte) (Signer, error) {
	return newHS(alg, key)
}

// NewVerifierHS returns a new HMAC-based verifier.
func NewVerifierHS(alg Algorithm, key []byte) (Verifier, error) {
	return newHS(alg, key)
}

type hmacAlgo interface {
	// copy-pasted Signer & Verifier due to older Go versions
	Algorithm() Algorithm
	SignSize() int
	Sign(payload []byte) ([]byte, error)
	Verify(payload, signature []byte) error
	VerifyToken(token *Token) error
}

func newHS(alg Algorithm, key []byte) (hmacAlgo, error) {
	if len(key) == 0 {
		return nil, ErrNilKey
	}
	hash, ok := getHashHMAC(alg)
	if !ok {
		return nil, ErrUnsupportedAlg
	}
	return &hsAlg{
		alg:  alg,
		hash: hash,
		key:  key,
		hashPool: &sync.Pool{New: func() interface{} {
			return hmac.New(hash.New, key)
		}},
	}, nil
}

func getHashHMAC(alg Algorithm) (crypto.Hash, bool) {
	switch alg {
	case HS256:
		return crypto.SHA256, true
	case HS384:
		return crypto.SHA384, true
	case HS512:
		return crypto.SHA512, true
	default:
		return 0, false
	}
}

type hsAlg struct {
	alg      Algorithm
	hash     crypto.Hash
	key      []byte
	hashPool *sync.Pool
}

func (hs *hsAlg) Algorithm() Algorithm {
	return hs.alg
}

func (hs *hsAlg) SignSize() int {
	return hs.hash.Size()
}

func (hs *hsAlg) Sign(payload []byte) ([]byte, error) {
	return hs.sign(payload)
}

func (hs *hsAlg) VerifyToken(token *Token) error {
	if constTimeAlgEqual(token.Header().Algorithm, hs.alg) {
		return hs.Verify(token.Payload(), token.Signature())
	}
	return ErrAlgorithmMismatch
}

func (hs *hsAlg) Verify(payload, signature []byte) error {
	digest, err := hs.sign(payload)
	if err != nil {
		return err
	}
	if !hmac.Equal(signature, digest) {
		return ErrInvalidSignature
	}
	return nil
}

func (hs *hsAlg) sign(payload []byte) ([]byte, error) {
	hasher := hs.hashPool.Get().(hash.Hash)
	defer func() {
		hasher.Reset()
		hs.hashPool.Put(hasher)
	}()

	_, err := hasher.Write(payload)
	if err != nil {
		return nil, err
	}
	return hasher.Sum(nil), nil
}
