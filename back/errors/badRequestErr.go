package errors

type BadRequestError struct {
	RequestType string
	Err         error
}

func (err *BadRequestError) Error() string {
	return "Request \"" + err.RequestType + "\" failed. Reason: " + err.Error()
}
