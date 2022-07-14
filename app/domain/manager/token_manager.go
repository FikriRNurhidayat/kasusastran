package manager

import (
	"time"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/grpclog"
)

type TokenManager interface {
	NewSession(userID string) (*Session, error)
}

type Session struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiredAt    time.Time `json:"expired_at"`
}

type AccessTokenClaims struct {
	jwt.StandardClaims
}
type RefreshTokenClaims struct {
	jwt.StandardClaims
}

type TokenManagerImpl struct {
	logger                grpclog.LoggerV2
	accessTokenSecretKey  string
	accessTokenExpiresIn  time.Duration
	refreshTokenSecretKey string
	refreshTokenExpiresIn time.Duration
	signingMethod         jwt.SigningMethod
}

// NewSession implements TokenManager
func (tm *TokenManagerImpl) NewSession(userID string) (session *Session, err error) {
	session = new(Session)
	now := time.Now()
	session.ExpiredAt = now.Add(tm.accessTokenExpiresIn)

	accessTokenClaims := AccessTokenClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   userID,
			ExpiresAt: session.ExpiredAt.Unix(),
			IssuedAt:  now.Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(tm.signingMethod, accessTokenClaims)
	session.AccessToken, err = accessToken.SignedString([]byte(tm.accessTokenSecretKey))
	if err != nil {
		return nil, err
	}

	refreshTokenClaims := RefreshTokenClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   userID,
			ExpiresAt: now.Add(tm.refreshTokenExpiresIn).Unix(),
			IssuedAt:  now.Unix(),
		},
	}

	refreshToken := jwt.NewWithClaims(tm.signingMethod, refreshTokenClaims)
	session.RefreshToken, err = refreshToken.SignedString([]byte(tm.refreshTokenSecretKey))
	if err != nil {
		return nil, err
	}

	return session, nil
}

func NewTokenManager(setters ...TokenManagerSetter) TokenManager {
	tm := new(TokenManagerImpl)

	for _, set := range setters {
		set(tm)
	}

	return tm
}

type TokenManagerSetter func(*TokenManagerImpl)

func WithLogger(logger grpclog.LoggerV2) TokenManagerSetter {
	return func(tm *TokenManagerImpl) {
		tm.logger = logger
	}
}

func WithSigningMethod(signingMethod jwt.SigningMethod) TokenManagerSetter {
	return func(tm *TokenManagerImpl) {
		tm.signingMethod = signingMethod
	}
}

func WithAccessTokenSecretKey(accessTokenSecretKey string) TokenManagerSetter {
	return func(tm *TokenManagerImpl) {
		tm.accessTokenSecretKey = accessTokenSecretKey
	}
}

func WithAccessTokenExpiresIn(accessTokenExpiresIn time.Duration) TokenManagerSetter {
	return func(tm *TokenManagerImpl) {
		tm.accessTokenExpiresIn = accessTokenExpiresIn
	}
}

func WithRefreshTokenSecretKey(refreshTokenSecretKey string) TokenManagerSetter {
	return func(tm *TokenManagerImpl) {
		tm.refreshTokenSecretKey = refreshTokenSecretKey
	}
}

func WithRefreshTokenExpiresIn(refreshTokenExpiresIn time.Duration) TokenManagerSetter {
	return func(tm *TokenManagerImpl) {
		tm.refreshTokenExpiresIn = refreshTokenExpiresIn
	}
}
