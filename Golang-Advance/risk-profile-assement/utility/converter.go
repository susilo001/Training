package utility

import "strings"

func ConvertUserMandatoryFieldErrorMessage(oldErrMsg string) string {
	switch {
	case strings.Contains(oldErrMsg, "'Name' Failed on the 'required' tag"):
		return "Field Name is Required!"
	case strings.Contains(oldErrMsg, "'Email' Failed on the 'required' tag"):
		return "Field Email is Required!"
	case strings.Contains(oldErrMsg, "'Password' Failed on the 'required' tag"):
		return "Field Password is Required!"
	}

	return oldErrMsg
}
