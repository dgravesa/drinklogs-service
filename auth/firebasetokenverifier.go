package auth

import "fmt"

// FirebaseTokenVerifier verifies user tokens using a Firebase authentication backend.
type FirebaseTokenVerifier struct {
}

// NewFirebaseTokenVerifier creates a new token verifier using a Firebase authentication backend.
func NewFirebaseTokenVerifier() (*FirebaseTokenVerifier, error) {
	// TODO implement
	return new(FirebaseTokenVerifier), nil
}

// Verify converts the token to a UID, or returns error if token is not a valid UID.
func (v *FirebaseTokenVerifier) Verify(token string) (uint64, error) {
	// TODO implement
	return 0, fmt.Errorf("not yet implemented")
}
