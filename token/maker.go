package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	// CreateToken creates a token for a specific username and valid duration
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
