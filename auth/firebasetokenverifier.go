package auth

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

// FirebaseTokenVerifier verifies user tokens using a Firebase authentication backend.
type FirebaseTokenVerifier struct {
	app    *firebase.App // TODO: not sure if this is needed to stick around
	client *auth.Client
}

// NewFirebaseTokenVerifier creates a new token verifier using a Firebase authentication backend.
func NewFirebaseTokenVerifier(keyname string) (*FirebaseTokenVerifier, error) {
	var tv FirebaseTokenVerifier
	var err error

	// initialize firebase app
	opt := option.WithCredentialsFile(keyname)
	tv.app, err = firebase.NewApp(context.Background(), nil, opt)

	// initialize authentication client
	if err == nil {
		tv.client, err = tv.app.Auth(context.Background())
	}

	return &tv, err
}

// Verify converts the token to a UID, or returns error if token is not a valid UID.
// TODO: may want to return user permissions on token verification.
func (tv *FirebaseTokenVerifier) Verify(ctx context.Context, token string) (string, error) {
	verified, err := tv.client.VerifyIDToken(ctx, token)

	return verified.UID, err
}
