package errors

type FailedSelectError struct {
	EntityType string
	Err        error
}

func (err *FailedSelectError) Error() string {
	return "Failed select for entity \"" + err.EntityType + "\". Reason: " + err.Err.Error()
}
