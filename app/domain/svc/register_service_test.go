package svc_test

import (
	"context"
	"testing"
	"time"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/manager"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"github.com/fikrirnurhidayat/kasusastran/app/trouble"
	mockEvent "github.com/fikrirnurhidayat/kasusastran/mocks/domain/event"
	mockManager "github.com/fikrirnurhidayat/kasusastran/mocks/domain/manager"
	mockRepo "github.com/fikrirnurhidayat/kasusastran/mocks/domain/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRegisterService struct {
	userRepository   *mockRepo.UserRepository
	userEventEmitter *mockEvent.UserEventEmitter
	passwordManager  *mockManager.PasswordManager
	tokenManager     *mockManager.TokenManager
}

func TestRegisterService_Call(t *testing.T) {
	type input struct {
		ctx    context.Context
		params *svc.RegisterParams
	}

	type output struct {
		result *svc.RegisterResult
		err    error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockRegisterService, *input, *output)
	}

	tests := []scenario{
		{
			name: "s.userRepository.EmailExist return error",
			in: &input{
				ctx: context.Background(),
				params: &svc.RegisterParams{
					Name:     "Fikri",
					Email:    "fikrirnurhidayat@gmail.com",
					Password: "123456",
				},
			},
			out: &output{
				result: nil,
				err:    trouble.INTERNAL_SERVER_ERROR,
			},
			on: func(mrs *MockRegisterService, i *input, o *output) {
				mrs.userRepository.On("EmailExist", i.ctx, i.params.Email).Return(false, o.err)
			},
		},
		{
			name: "s.userRepository.EmailExist return true",
			in: &input{
				ctx: context.Background(),
				params: &svc.RegisterParams{
					Name:     "Fikri",
					Email:    "fikrirnurhidayat@gmail.com",
					Password: "123456",
				},
			},
			out: &output{
				result: nil,
				err:    trouble.EMAIL_ALREADY_EXIST,
			},
			on: func(mrs *MockRegisterService, i *input, o *output) {
				mrs.userRepository.On("EmailExist", i.ctx, i.params.Email).Return(true, nil)
			},
		},
		{
			name: "s.passwordManager.Encrypt return error",
			in: &input{
				ctx: context.Background(),
				params: &svc.RegisterParams{
					Name:     "Fikri",
					Email:    "fikrirnurhidayat@gmail.com",
					Password: "123456",
				},
			},
			out: &output{
				result: nil,
				err:    trouble.INTERNAL_SERVER_ERROR,
			},
			on: func(mrs *MockRegisterService, i *input, o *output) {
				mrs.userRepository.On("EmailExist", i.ctx, i.params.Email).Return(false, nil)
				mrs.passwordManager.On("Encrypt", i.params.Password).Return("", o.err)
			},
		},
		{
			name: "s.userRepository.Create return error",
			in: &input{
				ctx: context.Background(),
				params: &svc.RegisterParams{
					Name:     "Fikri",
					Email:    "fikrirnurhidayat@gmail.com",
					Password: "123456",
				},
			},
			out: &output{
				result: nil,
				err:    trouble.INTERNAL_SERVER_ERROR,
			},
			on: func(mrs *MockRegisterService, i *input, o *output) {
				encryptedPassword := "$2a$10$O2pp2NvX/Y3OgBM66NUrjOtASWhg3rMhft4X0Ii4U8gX3AOJqcItK"

				iuser := &entity.User{
					Name:              i.params.Name,
					Email:             i.params.Email,
					EncryptedPassword: encryptedPassword,
				}

				ouser := &entity.User{
					ID:                uuid.New(),
					Name:              i.params.Name,
					Email:             i.params.Email,
					EncryptedPassword: encryptedPassword,
					CreatedAt:         time.Now(),
					UpdatedAt:         time.Now(),
				}

				mrs.userRepository.On("EmailExist", i.ctx, i.params.Email).Return(false, nil)
				mrs.passwordManager.On("Encrypt", i.params.Password).Return(encryptedPassword, nil)
				mrs.userRepository.On("Create", i.ctx, iuser).Return(ouser, o.err)
			},
		},
		{
			name: "s.userEventEmitter.EmitRegisteredEvent return error",
			in: &input{
				ctx: context.Background(),
				params: &svc.RegisterParams{
					Name:     "Fikri",
					Email:    "fikrirnurhidayat@gmail.com",
					Password: "123456",
				},
			},
			out: &output{
				result: nil,
				err:    trouble.INTERNAL_SERVER_ERROR,
			},
			on: func(mrs *MockRegisterService, i *input, o *output) {
				encryptedPassword := "$2a$10$O2pp2NvX/Y3OgBM66NUrjOtASWhg3rMhft4X0Ii4U8gX3AOJqcItK"

				iuser := &entity.User{
					Name:              i.params.Name,
					Email:             i.params.Email,
					EncryptedPassword: encryptedPassword,
				}

				ouser := &entity.User{
					ID:                uuid.New(),
					Name:              i.params.Name,
					Email:             i.params.Email,
					EncryptedPassword: encryptedPassword,
					CreatedAt:         time.Now(),
					UpdatedAt:         time.Now(),
				}

				mrs.userRepository.On("EmailExist", i.ctx, i.params.Email).Return(false, nil)
				mrs.passwordManager.On("Encrypt", i.params.Password).Return(encryptedPassword, nil)
				mrs.userRepository.On("Create", i.ctx, iuser).Return(ouser, nil)
				mrs.userEventEmitter.On("EmitRegisteredEvent", mock.AnythingOfType("*message.User")).Return(o.err)
			},
		},
		{
			name: "s.tokenManager.NewSession return error",
			in: &input{
				ctx: context.Background(),
				params: &svc.RegisterParams{
					Name:     "Fikri",
					Email:    "fikrirnurhidayat@gmail.com",
					Password: "123456",
				},
			},
			out: &output{
				result: nil,
				err:    trouble.INTERNAL_SERVER_ERROR,
			},
			on: func(mrs *MockRegisterService, i *input, o *output) {
				encryptedPassword := "$2a$10$O2pp2NvX/Y3OgBM66NUrjOtASWhg3rMhft4X0Ii4U8gX3AOJqcItK"

				iuser := &entity.User{
					Name:              i.params.Name,
					Email:             i.params.Email,
					EncryptedPassword: encryptedPassword,
				}

				ouser := &entity.User{
					ID:                uuid.New(),
					Name:              i.params.Name,
					Email:             i.params.Email,
					EncryptedPassword: encryptedPassword,
					CreatedAt:         time.Now(),
					UpdatedAt:         time.Now(),
				}

				mrs.userRepository.On("EmailExist", i.ctx, i.params.Email).Return(false, nil)
				mrs.passwordManager.On("Encrypt", i.params.Password).Return(encryptedPassword, nil)
				mrs.userRepository.On("Create", i.ctx, iuser).Return(ouser, nil)
				mrs.userEventEmitter.On("EmitRegisteredEvent", mock.AnythingOfType("*message.User")).Return(nil)
				mrs.tokenManager.On("NewSession", ouser.ID.String()).Return(nil, o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				params: &svc.RegisterParams{
					Name:     "Fikri",
					Email:    "fikrirnurhidayat@gmail.com",
					Password: "123456",
				},
			},
			out: &output{
				result: &svc.RegisterResult{
					AccessToken:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI4ZTgyMDg5ZC1kNmRlLTQ3YTktYTQ0Ni03OTY5MDE3MWFhYTQiLCJpYXQiOjE1MTYyMzkwMjJ9.LQ2XHmPY8NHzt-nqs6x8rPOJbqGwjjNKEX6u051XqLE",
					RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJleUpoYkdjaU9pSklVekkxTmlJc0luUjVjQ0k2SWtwWFZDSjkuZXlKemRXSWlPaUk0WlRneU1EZzVaQzFrTm1SbExUUTNZVGt0WVRRME5pMDNPVFk1TURFM01XRmhZVFFpTENKcFlYUWlPakUxTVRZeU16a3dNako5LkxRMlhIbVBZOE5IenQtbnFzNng4clBPSmJxR3dqak5LRVg2dTA1MVhxTEUiLCJpYXQiOjE1MTYyMzkwMjJ9.up24l_hyewj6w0JqGJYnQ4mxg4tuL4qd2nW0uhIED54",
					ExpiredAt:    time.Now().Add(5 * time.Minute),
				},
				err: nil,
			},
			on: func(mrs *MockRegisterService, i *input, o *output) {
				encryptedPassword := "$2a$10$O2pp2NvX/Y3OgBM66NUrjOtASWhg3rMhft4X0Ii4U8gX3AOJqcItK"

				iuser := &entity.User{
					Name:              i.params.Name,
					Email:             i.params.Email,
					EncryptedPassword: encryptedPassword,
				}

				ouser := &entity.User{
					ID:                uuid.New(),
					Name:              i.params.Name,
					Email:             i.params.Email,
					EncryptedPassword: encryptedPassword,
					CreatedAt:         time.Now(),
					UpdatedAt:         time.Now(),
				}

				session := &manager.Session{
					AccessToken:  o.result.AccessToken,
					RefreshToken: o.result.RefreshToken,
					ExpiredAt:    o.result.ExpiredAt,
				}

				mrs.userRepository.On("EmailExist", i.ctx, i.params.Email).Return(false, nil)
				mrs.passwordManager.On("Encrypt", i.params.Password).Return(encryptedPassword, nil)
				mrs.userRepository.On("Create", i.ctx, iuser).Return(ouser, nil)
				mrs.userEventEmitter.On("EmitRegisteredEvent", mock.AnythingOfType("*message.User")).Return(nil)
				mrs.tokenManager.On("NewSession", ouser.ID.String()).Return(session, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockRegisterService{
				userEventEmitter: new(mockEvent.UserEventEmitter),
				userRepository:   new(mockRepo.UserRepository),
				passwordManager:  new(mockManager.PasswordManager),
				tokenManager:     new(mockManager.TokenManager),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := svc.NewRegisterService(m.userRepository, m.userEventEmitter, m.passwordManager, m.tokenManager)
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
