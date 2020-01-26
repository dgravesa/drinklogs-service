package auth

import "strconv"

// TestTokenVerifier is a simple token verifier which assumes the token is
// merely a string stating the user's ID.
type TestTokenVerifier struct{}

// NewTestTokenVerifier returns a new test token verifier.
func NewTestTokenVerifier() *TestTokenVerifier {
	return new(TestTokenVerifier)
}

// Verify converts the token to a UID, or returns error if token is not a valid UID.
func (v TestTokenVerifier) Verify(token string) (uint64, error) {
	return strconv.ParseUint(token, 10, 64)
}
