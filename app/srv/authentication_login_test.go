package srv_test

import (
	"context"
	"testing"
	"time"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"github.com/fikrirnurhidayat/kasusastran/app/srv"
	"github.com/fikrirnurhidayat/kasusastran/app/trouble"
	"github.com/fikrirnurhidayat/kasusastran/lib/troublemaker"
	"github.com/fikrirnurhidayat/kasusastran/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"

	mockService "github.com/fikrirnurhidayat/kasusastran/mocks/domain/svc"
)

func TestAuthentication_Login(t *testing.T) {
	type input struct {
		ctx context.Context
		req *proto.LoginRequest
	}

	type output struct {
		res *proto.Session
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
			name: "Invalid Request: Email",
			in: &input{
				ctx: context.Background(),
				req: &proto.LoginRequest{
					Email:    "notanemail",
					Password: "123456",
				},
			},
			out: &output{
				res: nil,
			},
			on: func(m *MockAuthenticationServer, i *input, o *output) {
				o.err = troublemaker.FromValidationError(i.req.Validate())
			},
		},
		{
			name: "Invalid Request: Password",
			in: &input{
				ctx: context.Background(),
				req: &proto.LoginRequest{
					Email:    "fikrirnurhidayat@gmail.com",
					Password: "12345",
				},
			},
			out: &output{
				res: nil,
			},
			on: func(m *MockAuthenticationServer, i *input, o *output) {
				o.err = troublemaker.FromValidationError(i.req.Validate())
			},
		},
		{
			name: "s.loginService.Call return error",
			in: &input{
				ctx: context.Background(),
				req: &proto.LoginRequest{
					Email:    "fikrirnurhidayat@gmail.com",
					Password: "123456",
				},
			},
			out: &output{
				res: nil,
				err: trouble.EMAIL_ALREADY_EXIST,
			},
			on: func(m *MockAuthenticationServer, i *input, o *output) {
				params := &svc.LoginParams{
					Email:    i.req.GetEmail(),
					Password: i.req.GetPassword(),
				}

				m.loginService.On("Call", i.ctx, params).Return(nil, trouble.EMAIL_ALREADY_EXIST)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				req: &proto.LoginRequest{
					Email:    "fikrirnurhidayat@gmail.com",
					Password: "123456",
				},
			},
			out: &output{
				res: &proto.Session{
					AccessToken:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
					RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJleUpoYkdjaU9pSklVekkxTmlJc0luUjVjQ0k2SWtwWFZDSjkuZXlKemRXSWlPaUl4TWpNME5UWTNPRGt3SWl3aWJtRnRaU0k2SWtwdmFHNGdSRzlsSWl3aWFXRjBJam94TlRFMk1qTTVNREl5ZlEuU2ZsS3h3UkpTTWVLS0YyUVQ0ZndwTWVKZjM2UE9rNnlKVl9hZFFzc3c1YyIsImlhdCI6MTUxNjIzOTAyMn0.K-071Np4GBVsd7utmK1dTyViLHXUAFjN2W5iVMDGbOs",
					ExpiredAt:    timestamppb.New(time.Now().Add(5 * time.Minute)),
				},
				err: nil,
			},
			on: func(m *MockAuthenticationServer, i *input, o *output) {
				params := &svc.LoginParams{
					Email:    i.req.GetEmail(),
					Password: i.req.GetPassword(),
				}

				result := &svc.LoginResult{
					AccessToken:  o.res.GetAccessToken(),
					RefreshToken: o.res.GetRefreshToken(),
					ExpiredAt:    o.res.GetExpiredAt().AsTime(),
				}

				m.loginService.On("Call", i.ctx, params).Return(result, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockAuthenticationServer{
				loginService: new(mockService.LoginService),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := srv.NewAuthenticationsServer(srv.WithLoginService(m.loginService))
			res, err := subject.Login(tt.in.ctx, tt.in.req)

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
