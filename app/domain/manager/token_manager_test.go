package manager_test

import (
	"testing"
	"time"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/manager"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func TestTokenManager_NewSession(t *testing.T) {
	subject := manager.NewTokenManager(
		manager.WithSigningMethod(jwt.SigningMethodHS256),
		manager.WithAccessTokenExpiresIn(5*time.Minute),
		manager.WithAccessTokenSecretKey("rahasia"),
		manager.WithRefreshTokenExpiresIn(24*time.Hour),
		manager.WithRefreshTokenSecretKey("rahasia"),
	)

	session, err := subject.NewSession(uuid.New().String())

	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err = jwt.Parse(session.AccessToken, func(t *jwt.Token) (interface{}, error) {
		return []byte("rahasia"), nil
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err = jwt.Parse(session.RefreshToken, func(t *jwt.Token) (interface{}, error) {
		return []byte("rahasia"), nil
	})

	if err != nil {
		t.Fatalf(err.Error())
	}
}
