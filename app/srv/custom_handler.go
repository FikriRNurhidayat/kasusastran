package srv

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

type Error struct {
	Code    codes.Code `json:"code"`
	Reason  string     `json:"reason"`
	Message string     `json:"message"`
}

type ErrorResponse struct {
	Error *Error `json:"error"`
}

func CustomErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, err error) {
	var res *ErrorResponse
	st := status.Convert(err)
	code := st.Code()
	message := st.Message()
	httpStatus := runtime.HTTPStatusFromCode(st.Code())
	details := st.Details()

	if len(details) > 0 {
		detail := details[0]
		switch detail := detail.(type) {
		case *errdetails.ErrorInfo:
			reason := detail.Reason

			res = &ErrorResponse{
				Error: &Error{
					Code:    code,
					Reason:  reason,
					Message: message,
				},
			}
		case *errdetails.BadRequest:
			res = &ErrorResponse{
				Error: &Error{
					Code:    code,
					Reason:  "BAD_REQUEST",
					Message: message,
				},
			}
		}
	}

	if res == nil {
		res = &ErrorResponse{
			Error: &Error{
				Code:    code,
				Reason:  "UNIMPLEMENTED",
				Message: message,
			},
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	body, _ := json.Marshal(res)
	w.Write(body)
}
