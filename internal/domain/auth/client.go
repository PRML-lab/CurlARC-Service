package auth

import (
	"context"

	"firebase.google.com/go/v4/auth"
)

type AuthClient interface {
	VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error)
	DeleteUser(ctx context.Context, uid string) error
	UpdateUser(ctx context.Context, uid string, user *auth.UserToUpdate) (*auth.UserRecord, error)
	// Add any other methods you're using from the Firebase Auth Client
}
