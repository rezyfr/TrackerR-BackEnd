package token

import "time"

// Maker is the interface for managing tokens
type Maker interface {
	// Create token for a specific email and duration
	CreateToken(email string, duration time.Duration) (string, *Payload, error)
	// Validate token check if the token is invalid or not
	VerifyToken(token string) (*Payload, error)
}
