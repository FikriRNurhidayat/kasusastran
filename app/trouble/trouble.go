package trouble

import (
	"github.com/fikrirnurhidayat/kasusastran/app/config"
	"github.com/fikrirnurhidayat/kasusastran/lib/troublemaker"
	"google.golang.org/grpc/codes"
)

var trouble = troublemaker.NewTroubleMaker(
	troublemaker.WithDomain(config.GetDomain()),
	troublemaker.WithService(config.GetService()),
	troublemaker.WithCodes(
		map[string]codes.Code{
			REASON_EMAIL_ALREADY_EXIST:         codes.AlreadyExists,
			REASON_EMAIL_NOT_FOUND:             codes.NotFound,
			REASON_PASSWORD_INCORRECT:          codes.InvalidArgument,
			REASON_GOOGLE_ACCESS_TOKEN_INVALID: codes.InvalidArgument,
			REASON_ACCESS_TOKEN_INVALID:        codes.Unauthenticated,
			REASON_ACCESS_TOKEN_EXPIRED:        codes.Unauthenticated,
			REASON_REFRESH_TOKEN_INVALID:       codes.Unauthenticated,
			REASON_REFRESH_TOKEN_MISMATCH:      codes.Unauthenticated,
			REASON_REFRESH_TOKEN_EXPIRED:       codes.Unauthenticated,
			REASON_INTERNAL_SERVER_ERROR:       codes.Internal,
		},
	),
	troublemaker.WithMessages(
		map[string]string{
			REASON_EMAIL_ALREADY_EXIST:         "Email already exist. Please use another email.",
			REASON_EMAIL_NOT_FOUND:             "Email not found. Please pass a valid email.",
			REASON_PASSWORD_INCORRECT:          "Password not correct. Please pass a valid password for given email.",
			REASON_GOOGLE_ACCESS_TOKEN_INVALID: "Google Oauth Access Token is invalid. Make sure pass the valid Google OAuth Access Token.",
			REASON_ACCESS_TOKEN_INVALID:        "Access token not invalid. Please pass a valid access token.",
			REASON_ACCESS_TOKEN_EXPIRED:        "Access token is expired. Please refresh access token.",
			REASON_REFRESH_TOKEN_INVALID:       "Refresh token not invalid. Please pass a valid refresh token.",
			REASON_REFRESH_TOKEN_MISMATCH:      "Token pair is not valid. Please pass valid pair of access token and refresh token.",
			REASON_REFRESH_TOKEN_EXPIRED:       "Refresh token is expired. Please relogin.",
			REASON_INTERNAL_SERVER_ERROR:       "Internal server error has occured. Please report this to back-end developer.",
		},
	),
)

var (
	EMAIL_NOT_FOUND             = trouble.New(REASON_EMAIL_NOT_FOUND)
	EMAIL_ALREADY_EXIST         = trouble.New(REASON_EMAIL_ALREADY_EXIST)
	PASSWORD_INCORRECT          = trouble.New(REASON_PASSWORD_INCORRECT)
	ACCESS_TOKEN_INVALID        = trouble.New(REASON_ACCESS_TOKEN_INVALID)
	ACCESS_TOKEN_EXPIRED        = trouble.New(REASON_ACCESS_TOKEN_EXPIRED)
	REFRESH_TOKEN_INVALID       = trouble.New(REASON_REFRESH_TOKEN_INVALID)
	REFRESH_TOKEN_MISMATCH      = trouble.New(REASON_REFRESH_TOKEN_MISMATCH)
	REFRESH_TOKEN_EXPIRED       = trouble.New(REASON_REFRESH_TOKEN_EXPIRED)
	GOOGLE_ACCESS_TOKEN_INVALID = trouble.New(REASON_GOOGLE_ACCESS_TOKEN_INVALID)
	INTERNAL_SERVER_ERROR       = trouble.New(REASON_INTERNAL_SERVER_ERROR)
)
