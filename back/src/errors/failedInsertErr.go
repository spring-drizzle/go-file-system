package errors

type FailedInsertError struct {
	Entity string
	Err    error
}

func (err *FailedInsertError) Error() string {
	return err.Entity + " was not inserted. Reson: " + err.Err.Error()
}
