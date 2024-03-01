package response

type ApiError struct {
	code    int
	message string
}

func (ae *ApiError) GetCode() int {
	return ae.code
}

func (ae *ApiError) GetMessage() string {
	return ae.message
}

var (
	ParamError = ApiError{
		code:    10000,
		message: "param error",
	}
)
