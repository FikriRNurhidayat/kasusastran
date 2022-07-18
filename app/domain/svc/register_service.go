package svc

import (
	"context"
	"time"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/event"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/manager"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/message"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository"
	"github.com/fikrirnurhidayat/kasusastran/app/trouble"
	"github.com/google/uuid"
	"google.golang.org/grpc/grpclog"
)

type RegisterService interface {
	Call(context.Context, *RegisterParams) (*RegisterResult, error)
}

type RegisterParams struct {
	Name     string
	Email    string
	Password string
}

type RegisterResult manager.Session

type registerService struct {
	logger           grpclog.LoggerV2
	userRepository   repository.UserRepository
	userEventEmitter event.UserEventEmitter
	passwordManager  manager.PasswordManager
	tokenManager     manager.TokenManager
}

func (s *registerService) Call(ctx context.Context, params *RegisterParams) (*RegisterResult, error) {
	doesEmailExist, err := s.userRepository.EmailExist(ctx, params.Email)

	if err != nil {
		return nil, trouble.INTERNAL_SERVER_ERROR
	}

	if doesEmailExist {
		return nil, trouble.EMAIL_ALREADY_EXIST
	}

	encryptedPassword, err := s.passwordManager.Encrypt(params.Password)

	if err != nil {
		return nil, trouble.INTERNAL_SERVER_ERROR
	}

	user, err := s.userRepository.Create(ctx, &entity.User{
		Name:              params.Name,
		Email:             params.Email,
		EncryptedPassword: encryptedPassword,
	})

	if err != nil {
		return nil, trouble.INTERNAL_SERVER_ERROR
	}

	err = s.userEventEmitter.EmitRegisteredEvent(&message.User{
		ID:        uuid.New(),
		Kind:      event.USER_REGISTERED_TOPIC,
		CreatedAt: time.Now(),
		Actor:     &message.Actor{ID: user.ID.String(), Kind: "USER"},
		Payload:   user,
	})

	if err != nil {
		return nil, trouble.INTERNAL_SERVER_ERROR
	}

	session, err := s.tokenManager.NewSession(user.ID.String())

	if err != nil {
		return nil, trouble.INTERNAL_SERVER_ERROR
	}

	res := &RegisterResult{
		AccessToken:  session.AccessToken,
		RefreshToken: session.RefreshToken,
		ExpiredAt:    session.ExpiredAt,
	}

	return res, nil
}

func NewRegisterService(logger grpclog.LoggerV2, userRepository repository.UserRepository, userEventEmitter event.UserEventEmitter, passwordManager manager.PasswordManager, tokenManager manager.TokenManager) RegisterService {
	return &registerService{
		logger:           logger,
		userRepository:   userRepository,
		userEventEmitter: userEventEmitter,
		passwordManager:  passwordManager,
		tokenManager:     tokenManager,
	}
}
