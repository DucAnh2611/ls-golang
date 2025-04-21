package constants

type ResponseCode string

const (
	SuccessOk ResponseCode = "S_OK"
)

const (
	ErrUnauthorized  ResponseCode = "E_UNAUTHORIZED"
	ErrBadRequest    ResponseCode = "E_BAD_REQUEST"
	ErrInvalidToken  ResponseCode = "E_INVALID_TOKEN"
	ErrInvalidClaims ResponseCode = "E_INVALID_CLAIMS"
	ErrInternal      ResponseCode = "E_INTERNAL"
	// ...Add more here
)
