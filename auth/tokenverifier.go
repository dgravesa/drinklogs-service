package auth

// TokenVerifier verifies a token and returns the corresponding user if the
// token is valid or error of the token is invalid.
type TokenVerifier interface {
	Verify(token string) (uint64, error)
}

var tokenVerifier TokenVerifier

// SetTokenVerifier sets the global token verifier.
func SetTokenVerifier(newTokenVerifier TokenVerifier) {
	tokenVerifier = newTokenVerifier
}

// VerifyToken verifies a user's identity from a token; returns error if token is invalid.
func VerifyToken(token string) (uint64, error) {
	return tokenVerifier.Verify(token)
}
