package auth

import (
	"context"
)

// TestTokenVerifier is a simple token verifier which assumes the token is
// merely a string stating the user's ID.
type TestTokenVerifier struct{}

// NewTestTokenVerifier returns a new test token verifier.
func NewTestTokenVerifier() *TestTokenVerifier {
	return new(TestTokenVerifier)
}

// Verify assumes token is valid for test verifier as passes it through without error.
func (v TestTokenVerifier) Verify(ctx context.Context, token string) (string, error) {
	return token, nil
}
