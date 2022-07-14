package srv_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/fikrirnurhidayat/kasusastran/api"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"github.com/fikrirnurhidayat/kasusastran/app/srv"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	mockService "github.com/fikrirnurhidayat/kasusastran/mocks/domain/svc"
)

func TestAuthentication_Register(t *testing.T) {
	type input struct {
		ctx context.Context
		req *api.RegisterRequest
	}

	type output struct {
		res *api.Session
		err error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockAuthenticationServer, *input, *output)
	}

	tests := []scenario{
		{
			name: "Invalid Request: Name",
			in: &input{
				ctx: context.Background(),
				req: &api.RegisterRequest{
					Name:     "",
					Email:    "fikrirnurhidayat@gmail.com",
					Password: "123456",
				},
			},
			out: &output{
				res: nil,
				err: status.Errorf(codes.InvalidArgument, "invalid RegisterRequest.Name: value length must be between 3 and 255 runes, inclusive"),
			},
			on: func(m *MockAuthenticationServer, i *input, o *output) {},
		},
		{
			name: "Invalid Request: Email",
			in: &input{
				ctx: context.Background(),
				req: &api.RegisterRequest{
					Name:     "Fikri",
					Email:    "notanemail",
					Password: "123456",
				},
			},
			out: &output{
				res: nil,
				err: status.Errorf(codes.InvalidArgument, "invalid RegisterRequest.Email: value must be a valid email address | caused by: mail: missing '@' or angle-addr"),
			},
			on: func(m *MockAuthenticationServer, i *input, o *output) {},
		},
		{
			name: "Invalid Request: Password",
			in: &input{
				ctx: context.Background(),
				req: &api.RegisterRequest{
					Name:     "Fikri",
					Email:    "fikrirnurhidayat@gmail.com",
					Password: "12345",
				},
			},
			out: &output{
				res: nil,
				err: status.Errorf(codes.InvalidArgument, "invalid RegisterRequest.Password: value length must be between 6 and 255 runes, inclusive"),
			},
			on: func(m *MockAuthenticationServer, i *input, o *output) {},
		},
		{
			name: "s.registerService.Call return error",
			in: &input{
				ctx: context.Background(),
				req: &api.RegisterRequest{
					Name:     "Fikri",
					Email:    "fikrirnurhidayat@gmail.com",
					Password: "123456",
				},
			},
			out: &output{
				res: nil,
				err: status.Errorf(codes.InvalidArgument, "something is error"),
			},
			on: func(m *MockAuthenticationServer, i *input, o *output) {
				params := &svc.RegisterParams{
					Name:     i.req.GetName(),
					Email:    i.req.GetEmail(),
					Password: i.req.GetPassword(),
				}

				m.registerService.On("Call", i.ctx, params).Return(nil, errors.New("something is error"))
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				req: &api.RegisterRequest{
					Name:     "Fikri",
					Email:    "fikrirnurhidayat@gmail.com",
					Password: "123456",
				},
			},
			out: &output{
				res: &api.Session{
					AccessToken:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
					RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJleUpoYkdjaU9pSklVekkxTmlJc0luUjVjQ0k2SWtwWFZDSjkuZXlKemRXSWlPaUl4TWpNME5UWTNPRGt3SWl3aWJtRnRaU0k2SWtwdmFHNGdSRzlsSWl3aWFXRjBJam94TlRFMk1qTTVNREl5ZlEuU2ZsS3h3UkpTTWVLS0YyUVQ0ZndwTWVKZjM2UE9rNnlKVl9hZFFzc3c1YyIsImlhdCI6MTUxNjIzOTAyMn0.K-071Np4GBVsd7utmK1dTyViLHXUAFjN2W5iVMDGbOs",
					ExpiredAt:    timestamppb.New(time.Now().Add(5 * time.Minute)),
				},
				err: nil,
			},
			on: func(m *MockAuthenticationServer, i *input, o *output) {
				params := &svc.RegisterParams{
					Name:     i.req.GetName(),
					Email:    i.req.GetEmail(),
					Password: i.req.GetPassword(),
				}

				result := &svc.RegisterResult{
					AccessToken:  o.res.GetAccessToken(),
					RefreshToken: o.res.GetRefreshToken(),
					ExpiredAt:    o.res.GetExpiredAt().AsTime(),
				}

				m.registerService.On("Call", i.ctx, params).Return(result, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockAuthenticationServer{
				registerService: new(mockService.RegisterService),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := srv.NewAuthenticationsServer(srv.WithRegisterService(m.registerService))
			res, err := subject.Register(tt.in.ctx, tt.in.req)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}

			if tt.out.res != nil {
				assert.NotNil(t, res)
				assert.Equal(t, tt.out.res, res)
			}
		})
	}
}
