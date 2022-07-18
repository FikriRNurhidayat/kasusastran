package svc_test

import (
	"context"
	"testing"
	"time"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/manager"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"github.com/fikrirnurhidayat/kasusastran/app/trouble"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mockEvent "github.com/fikrirnurhidayat/kasusastran/mocks/domain/event"
	mockManager "github.com/fikrirnurhidayat/kasusastran/mocks/domain/manager"
	mockRepo "github.com/fikrirnurhidayat/kasusastran/mocks/domain/repository"
	mockPackage "github.com/fikrirnurhidayat/kasusastran/mocks/package"
)

type MockLoginService struct {
	logger              *mockPackage.LoggerV2
	userRepository      *mockRepo.UserRepository
	sessionEventEmitter *mockEvent.SessionEventEmitter
	passwordManager     *mockManager.PasswordManager
	tokenManager        *mockManager.TokenManager
}

func TestLoginService_Call(t *testing.T) {
	type input struct {
		ctx    context.Context
		params *svc.LoginParams
	}

	type output struct {
		result *svc.LoginResult
		err    error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockLoginService, *input, *output)
	}

	tests := []scenario{
		{
			name: "s.userRepository.GetByEmail return error",
			in: &input{
				ctx: context.Background(),
				params: &svc.LoginParams{
					Email:    "fikrirnurhidayat@gmail.com",
					Password: "123456",
				},
			},
			out: &output{
				result: nil,
				err:    trouble.EMAIL_NOT_FOUND,
			},
			on: func(mls *MockLoginService, i *input, o *output) {
				mls.userRepository.On("GetByEmail", i.ctx, i.params.Email).Return(nil, o.err)
			},
		},
		{
			name: "s.passwordManager.Compare return error",
			in: &input{
				ctx: context.Background(),
				params: &svc.LoginParams{
					Email:    "fikrirnurhidayat@gmail.com",
					Password: "123456",
				},
			},
			out: &output{
				result: nil,
				err:    trouble.PASSWORD_INCORRECT,
			},
			on: func(mls *MockLoginService, i *input, o *output) {
				user := &entity.User{
					ID:                uuid.New(),
					Name:              "Fikri",
					Email:             i.params.Email,
					EncryptedPassword: "$2a$10$urYUQLb5L3UYn0J18obZ8OXN9IXLW77DKtG2GMBqtEl..YlDXPiMi",
					CreatedAt:         time.Now(),
					UpdatedAt:         time.Now(),
				}

				mls.userRepository.On("GetByEmail", i.ctx, i.params.Email).Return(user, nil)
				mls.passwordManager.On("Compare", user.EncryptedPassword, i.params.Password).Return(o.err)
			},
		},
		{
			name: "s.sessionEventEmitter.EmitCreatedEvent return error",
			in: &input{
				ctx: context.Background(),
				params: &svc.LoginParams{
					Email:    "fikrirnurhidayat@gmail.com",
					Password: "123456",
				},
			},
			out: &output{
				result: nil,
				err:    trouble.INTERNAL_SERVER_ERROR,
			},
			on: func(mls *MockLoginService, i *input, o *output) {
				user := &entity.User{
					ID:                uuid.New(),
					Name:              "Fikri",
					Email:             i.params.Email,
					EncryptedPassword: "$2a$10$urYUQLb5L3UYn0J18obZ8OXN9IXLW77DKtG2GMBqtEl..YlDXPiMi",
					CreatedAt:         time.Now(),
					UpdatedAt:         time.Now(),
				}

				mls.userRepository.On("GetByEmail", i.ctx, i.params.Email).Return(user, nil)
				mls.passwordManager.On("Compare", user.EncryptedPassword, i.params.Password).Return(nil)
				mls.sessionEventEmitter.On("EmitCreatedEvent", mock.AnythingOfType("*message.User")).Return(o.err)
			},
		},
		{
			name: "s.tokenManager.NewSession return error",
			in: &input{
				ctx: context.Background(),
				params: &svc.LoginParams{
					Email:    "fikrirnurhidayat@gmail.com",
					Password: "123456",
				},
			},
			out: &output{
				result: nil,
				err:    trouble.INTERNAL_SERVER_ERROR,
			},
			on: func(mls *MockLoginService, i *input, o *output) {
				user := &entity.User{
					ID:                uuid.New(),
					Name:              "Fikri",
					Email:             i.params.Email,
					EncryptedPassword: "$2a$10$urYUQLb5L3UYn0J18obZ8OXN9IXLW77DKtG2GMBqtEl..YlDXPiMi",
					CreatedAt:         time.Now(),
					UpdatedAt:         time.Now(),
				}

				mls.userRepository.On("GetByEmail", i.ctx, i.params.Email).Return(user, nil)
				mls.passwordManager.On("Compare", user.EncryptedPassword, i.params.Password).Return(nil)
				mls.sessionEventEmitter.On("EmitCreatedEvent", mock.AnythingOfType("*message.User")).Return(nil)
				mls.tokenManager.On("NewSession", user.ID.String()).Return(nil, o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				params: &svc.LoginParams{
					Email:    "fikrirnurhidayat@gmail.com",
					Password: "123456",
				},
			},
			out: &output{
				result: &svc.LoginResult{
					AccessToken:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI4ZTgyMDg5ZC1kNmRlLTQ3YTktYTQ0Ni03OTY5MDE3MWFhYTQiLCJpYXQiOjE1MTYyMzkwMjJ9.LQ2XHmPY8NHzt-nqs6x8rPOJbqGwjjNKEX6u051XqLE",
					RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJleUpoYkdjaU9pSklVekkxTmlJc0luUjVjQ0k2SWtwWFZDSjkuZXlKemRXSWlPaUk0WlRneU1EZzVaQzFrTm1SbExUUTNZVGt0WVRRME5pMDNPVFk1TURFM01XRmhZVFFpTENKcFlYUWlPakUxTVRZeU16a3dNako5LkxRMlhIbVBZOE5IenQtbnFzNng4clBPSmJxR3dqak5LRVg2dTA1MVhxTEUiLCJpYXQiOjE1MTYyMzkwMjJ9.up24l_hyewj6w0JqGJYnQ4mxg4tuL4qd2nW0uhIED54",
					ExpiredAt:    time.Now().Add(5 * time.Minute),
				},
				err: nil,
			},
			on: func(mls *MockLoginService, i *input, o *output) {
				user := &entity.User{
					ID:                uuid.New(),
					Name:              "Fikri",
					Email:             i.params.Email,
					EncryptedPassword: "$2a$10$urYUQLb5L3UYn0J18obZ8OXN9IXLW77DKtG2GMBqtEl..YlDXPiMi",
					CreatedAt:         time.Now(),
					UpdatedAt:         time.Now(),
				}

				session := &manager.Session{
					AccessToken:  o.result.AccessToken,
					RefreshToken: o.result.RefreshToken,
					ExpiredAt:    o.result.ExpiredAt,
				}

				mls.userRepository.On("GetByEmail", i.ctx, i.params.Email).Return(user, nil)
				mls.passwordManager.On("Compare", user.EncryptedPassword, i.params.Password).Return(nil)
				mls.sessionEventEmitter.On("EmitCreatedEvent", mock.AnythingOfType("*message.User")).Return(nil)
				mls.tokenManager.On("NewSession", user.ID.String()).Return(session, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockLoginService{
				logger:              new(mockPackage.LoggerV2),
				userRepository:      new(mockRepo.UserRepository),
				sessionEventEmitter: new(mockEvent.SessionEventEmitter),
				passwordManager:     new(mockManager.PasswordManager),
				tokenManager:        new(mockManager.TokenManager),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			m.logger.On("Error", mock.AnythingOfType("*status.Error"))

			subject := svc.NewLoginService(m.logger, m.userRepository, m.sessionEventEmitter, m.passwordManager, m.tokenManager)
			result, err := subject.Call(tt.in.ctx, tt.in.params)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}

			if tt.out.result != nil {
				assert.NotNil(t, result)
				assert.Equal(t, tt.out.result, result)
			}
		})
	}
}
