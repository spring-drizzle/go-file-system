package errors

type FailedInsertError struct {
	Entity string
}

func (err *FailedInsertError) Error() string {
	return err.Entity + " was not inserted"
}
