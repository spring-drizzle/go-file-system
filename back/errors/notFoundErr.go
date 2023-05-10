package errors

type NotFoundError struct {
	Entity string
	Id     string
}

func (err *NotFoundError) Error() string {
	return err.Entity + " for id " + err.Id + " was not found\n"
}
