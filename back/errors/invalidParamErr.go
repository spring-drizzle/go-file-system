package errors

import "fmt"

type InvalidParamError struct {
	ParamName  string
	ParamValue any
}

func (err *InvalidParamError) Error() string {
	return "Param \"" + err.ParamName + "\" has invalid value \"" + fmt.Sprint(err.ParamValue) + "\""
}
