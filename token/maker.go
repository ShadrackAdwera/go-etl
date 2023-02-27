package token

import "time"

type TokenMaker interface {
	CreateToken(username string, id int64, duration time.Duration) (string, error)
	VerifyToken(token string) (*TokenPayload, error)
}
