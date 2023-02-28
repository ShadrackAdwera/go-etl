package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"

	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (TokenMaker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("symmetric key should equal chachpoly : %v", chacha20poly1305.KeySize)
	}
	pasetoMaker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}
	return pasetoMaker, nil
}

func (maker *PasetoMaker) CreateToken(username string, id int64, email string, duration time.Duration) (*TokenPayload, string, error) {
	payload, err := NewPayload(id, username, email, duration)
	if err != nil {
		return &TokenPayload{}, "", err
	}
	tkn, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	return payload, tkn, err
}

func (pasetoMaker *PasetoMaker) VerifyToken(token string) (*TokenPayload, error) {
	payload := &TokenPayload{}
	err := pasetoMaker.paseto.Decrypt(token, pasetoMaker.symmetricKey, payload, nil)
	if err != nil {
		return nil, err
	}
	err = payload.Valid()
	if err != nil {
		return nil, err
	}
	return payload, nil
}
