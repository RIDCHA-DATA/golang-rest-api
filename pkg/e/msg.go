package e

var MsgFlags = map[int]string{
	SUCCESS:                        "Ok",
	ERROR:                          "Fail",
	INVALID_PARAMS:                 "Request parameter error",
	ERROR_NOT_EXIST_ACTIONS:        "The ACTION does not exist",
	ERROR_ADD_ACTION_FAIL:          "Failed to add ACTION",
	ERROR_DELETE_ACTION_FAIL:       "Failed to delete ACTION",
	ERROR_CHECK_EXIST_ACTION_FAIL:  "Failed to check if the ACTION exists",
	ERROR_EDIT_ACTION_FAIL:         "Failed to modify the ACTION",
	ERROR_COUNT_ACTIONS_FAIL:       "Statistics actions failed",
	ERROR_GET_ACTIONS_FAIL:         "Failed to get multiple ACTIONs",
	ERROR_GET_ACTION_FAIL:          "Failed to get a single ACTION",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token Authentication failed",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token Timed out",
	ERROR_AUTH_TOKEN:               "Token Build failed",
	ERROR_AUTH:                     "Token Error",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
