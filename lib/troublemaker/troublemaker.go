package troublemaker

import (
	"fmt"
	"regexp"
	"strings"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TroubleMaker interface {
	New(reason string) error
}

type TroubleMakerImpl struct {
	domain   string
	service  string
	codes    map[string]codes.Code
	messages map[string]string
}

// New implements TroubleMaker
func (tmi *TroubleMakerImpl) New(reason string) error {
	var cd codes.Code
	var ok bool

	if cd, ok = tmi.codes[reason]; !ok {
		panic("reason does not exist")
	}

	trouble := &Trouble{
		Code:    cd,
		Reason:  reason,
		Message: tmi.messages[reason],
	}

	st := status.New(trouble.Code, trouble.Message)
	st, _ = st.WithDetails(&errdetails.ErrorInfo{
		Reason: trouble.Reason,
		Domain: tmi.domain,
		Metadata: map[string]string{
			"service": tmi.service,
		},
	})

	return st.Err()
}

type TroubleMakerSetter func(*TroubleMakerImpl)

type ValidationError interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
}

type ValidationMultiError interface {
	Error() string
	AllErrors() []error
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func FromValidationError(err error) error {
	switch err := err.(type) {
	case ValidationError:
		field := ToSnakeCase(err.Field())

		trouble := &Trouble{
			Code:    codes.InvalidArgument,
			Reason:  "BAD_REQUEST",
			Message: fmt.Sprintf("%s: %s. Please pass valid %s.", field, err.Reason(), field),
		}

		st := status.New(trouble.Code, trouble.Message)
		st, _ = st.WithDetails(&errdetails.BadRequest{
			FieldViolations: []*errdetails.BadRequest_FieldViolation{
				{
					Field:       field,
					Description: err.Reason(),
				},
			},
		})

		return st.Err()
	case ValidationMultiError:
		var prefix []string
		trouble := &Trouble{
			Code:   codes.InvalidArgument,
			Reason: "BAD_REQUEST",
		}

		var fv []*errdetails.BadRequest_FieldViolation

		for _, e := range err.AllErrors() {

			if err, ok := e.(ValidationError); ok {
				prefix = append(prefix, fmt.Sprintf("%s: %s", ToSnakeCase(err.Field()), err.Reason()))

				fv = append(fv, &errdetails.BadRequest_FieldViolation{
					Field:       ToSnakeCase(err.Field()),
					Description: err.Reason(),
				})
			}
		}

		trouble.Message = fmt.Sprintf("%s. Please pass valid request parameters.", strings.Join(prefix, " & "))

		// NOTE: Provide rich gRPC error message
		st := status.New(trouble.Code, trouble.Message)
		st, _ = st.WithDetails(&errdetails.BadRequest{
			FieldViolations: fv,
		})

		return st.Err()
	default:
		return err
	}
}

func NewTroubleMaker(setters ...TroubleMakerSetter) TroubleMaker {
	tmi := new(TroubleMakerImpl)

	for _, set := range setters {
		set(tmi)
	}

	return tmi
}

func WithDomain(domain string) TroubleMakerSetter {
	return func(tmi *TroubleMakerImpl) {
		tmi.domain = domain
	}
}

func WithService(service string) TroubleMakerSetter {
	return func(tmi *TroubleMakerImpl) {
		tmi.service = service
	}
}

func WithCodes(codes map[string]codes.Code) TroubleMakerSetter {
	return func(tmi *TroubleMakerImpl) {
		tmi.codes = codes
	}
}

func WithMessages(messages map[string]string) TroubleMakerSetter {
	return func(tmi *TroubleMakerImpl) {
		tmi.messages = messages
	}
}
