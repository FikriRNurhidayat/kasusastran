package svc

import (
	"context"
	"time"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/event"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/manager"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/message"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository"
	"github.com/fikrirnurhidayat/kasusastran/app/trouble"
	"github.com/google/uuid"
	"google.golang.org/grpc/grpclog"
)

type LoginService interface {
	Call(context.Context, *LoginParams) (*LoginResult, error)
}

type LoginParams struct {
	Email    string
	Password string
}

type LoginResult manager.Session

type loginService struct {
	logger              grpclog.LoggerV2
	userRepository      repository.UserRepository
	sessionEventEmitter event.SessionEventEmitter
	passwordManager     manager.PasswordManager
	tokenManager        manager.TokenManager
}

// Call implements LoginService
func (s *loginService) Call(ctx context.Context, params *LoginParams) (*LoginResult, error) {
	// Find user by email
	user, err := s.userRepository.GetByEmail(ctx, params.Email)
	if err != nil {
		return nil, trouble.EMAIL_NOT_FOUND
	}

	// Compare password
	err = s.passwordManager.Compare(user.EncryptedPassword, params.Password)
	if err != nil {
		return nil, trouble.PASSWORD_INCORRECT
	}

	// Emit sessions.created event
	err = s.sessionEventEmitter.EmitCreatedEvent(&message.User{
		ID:        uuid.New(),
		Kind:      event.SESSION_CREATED_TOPIC,
		CreatedAt: time.Now(),
		Actor: &message.Actor{
			ID:   user.ID.String(),
			Kind: "USER",
		},
		Payload: user,
	})
	if err != nil {
		s.logger.Error(err)
		return nil, trouble.INTERNAL_SERVER_ERROR
	}

	session, err := s.tokenManager.NewSession(user.ID.String())
	if err != nil {
		s.logger.Error(err)
		return nil, trouble.INTERNAL_SERVER_ERROR
	}

	result := &LoginResult{
		AccessToken:  session.AccessToken,
		RefreshToken: session.RefreshToken,
		ExpiredAt:    session.ExpiredAt,
	}

	return result, nil
}

func NewLoginService(logger grpclog.LoggerV2, userRepository repository.UserRepository, sessionEventEmitter event.SessionEventEmitter, passwordManager manager.PasswordManager, tokenManager manager.TokenManager) LoginService {
	return &loginService{
		logger:              logger,
		userRepository:      userRepository,
		sessionEventEmitter: sessionEventEmitter,
		passwordManager:     passwordManager,
		tokenManager:        tokenManager,
	}
}
